package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const filePath = "./data/by_size.json"

type BySize struct {
	Size    int
	Mallocs int
	Frees   int
}

func main() {
	// Read BySize data
	bySizeJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var bySize []BySize
	err = json.Unmarshal(bySizeJson, &bySize)
	if err != nil {
		panic(err)
	}

	// Calculate total bytes malloc'd and free'd
	var totalMallocs int
	var totalFrees int
	var totalObjects int
	for _, bs := range bySize {
		totalMallocs += bs.Mallocs * bs.Size
		totalFrees += bs.Frees * bs.Size
		totalObjects += bs.Size
	}

	fmt.Println("Total bytes malloc'd:", totalMallocs)
	fmt.Println("Total bytes free'd:", totalFrees)
	fmt.Println("Total objects allocated:", totalObjects)
}
