package main

import "encoding/json"
import "fmt"

type Resource struct {
	Resource  string `json:"resource"`
	NameSpace string `json:"namespace"`
}

type Component struct {
	Id       string     `json:"id"`
	FolderId string     `json:"folderId"`
	Js       []Resource `json:"js"`
}

func main() {
	component := Component{
		Id:       "lightning",
		FolderId: "skuidother",
		Js: []Resource{
			{Resource: "LightningBuilderJS", NameSpace: "skuid"}}}
	output, _ := json.MarshalIndent(&component, "", "    ")
	fmt.Println(string(output))
}
