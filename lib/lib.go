package lib

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/ghodss/yaml"
)

type File struct {
	Provider         string `json:"provider"` // Affects YAML field names too.
	Region           string `json:"region"`
	Type             string `json:"type"`
	Disk             int    `json:"disk"`
	Target_num_nodes int    `json:"target_num_nodes"`
	Stackfile        string `json:"stackfile"`
}

func ReadFile(path string) {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(buf, f)
	defer f.Close()
	parsefile(buf.Bytes())
}

func parsefile(y []byte) {
	file, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(file))
}

func sendAPIRequest() {

}
