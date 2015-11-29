package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/MaximeHeckel/tutum-machine/tools"
	"github.com/tutumcloud/go-tutum/tutum"
)

func ReadFile(path string) {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(buf, f)
	defer f.Close()
	file, err := tools.Parsefile(buf.Bytes())
	if err != nil {
		fmt.Println(err)
	}

	nodecluster, err := tools.ConvertToStruct(file)
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range nodecluster {
		fmt.Printf("==> Processing cluster %s", key)
		sendNodeClusterAPIRequest(key, value)
	}
}

func sendNodeClusterAPIRequest(name string, nodecluster tools.NodeFile) {
	newCluster := tutum.NodeCreateRequest{Name: name,
		Region:           "/api/v1/region/" + nodecluster.Provider + "/" + nodecluster.Region + "/",
		NodeType:         "/api/v1/nodetype/" + nodecluster.Provider + "/" + nodecluster.Type + "/",
		Target_num_nodes: nodecluster.Target_num_nodes,
		Tags:             []tutum.NodeTag{{Name: name}},
		Disk:             nodecluster.Disk}

	cluster, err := tutum.CreateNodeCluster(newCluster)

	if err != nil {
		log.Println(err)
	}

	err = cluster.Deploy()

	if err == nil {
	Loop:
		for {
			nodeclusters, _ := tutum.ListNodeClusters()
			for _, cluster := range nodeclusters.Objects {
				if cluster.Name == name && cluster.State == "Deployed" {
					break Loop
				} else {
					time.Sleep(30 * time.Second)
				}
			}
		}
		fmt.Printf("Cluster %s launched!\n", name)
		fmt.Printf("Processing Stackfile %s", nodecluster.Stackfile)
	}
}

func fetchStackfile() {

}

func processStackfile() {

}

func launchStackfile() {

}
