package main

import (
	"encoding/json"
	"fmt"
)

type BagItem struct {
	Name    string    `json:"server"`
	Content []Content `json:"exclusions"`
}
type Content struct {
	GUID        string `json:"guid"`
	Path        string `json:"path"`
	Extension   string `json:"extension"`
	ProcessPath string `json:"processPath"`
}

func main() {
	serverA := &Content{
		GUID:        "357743e9-8fee-4479-997d-66df04f741d5",
		Path:        "bar",
		Extension:   "baz",
		ProcessPath: "babu",
	}
	serverB := &Content{
		GUID:        "22c2e7ee-a792-418e-b588-9271d24b4aff",
		Path:        "bar",
		Extension:   "baz",
		ProcessPath: "babu",
	}
	item := &BagItem{
		Name:    "item-a",
		Content: []Content{*serverA, *serverB},
	}

	payload, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(payload))
}
