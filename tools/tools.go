package tools

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
)

type NodeClusterFile map[string]NodeFile

type NodeFile struct {
	Provider         string `json:"provider"`
	Region           string `json:"region"`
	Type             string `json:"type"`
	Disk             int    `json:"disk"`
	Target_num_nodes int    `json:"target_num_nodes"`
	Stackfile        string `json:"stackfile"`
}

func Parsefile(y []byte) ([]byte, error) {
	file, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return file, nil
}

func ConvertToStruct(data []byte) (NodeClusterFile, error) {
	var response NodeClusterFile
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
