package main

import "fmt"

// START OMIT
import "encoding/json"

var data = []byte(`{"id": "lightning", "folderId": "skuidother",
	"js": [{"resource": "LightningBuilderJS",
		    "namespace": "skuid"}]}`)

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
	component := Component{}
	json.Unmarshal(data, &component)
	fmt.Printf("%#v\n", component)
	fmt.Printf("Id = \"%s\"\n", component.Id)
}

// END OMIT
