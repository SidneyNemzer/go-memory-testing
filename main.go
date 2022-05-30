package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
)

const dataDirectory = "./data/"

func main() {
	writeMemStats()

	stuff := allocMem()

	time.Sleep(time.Second * 10)

	fmt.Printf("%d\n", len(stuff))

	writeMemStats()

	time.Sleep(time.Second * 10)
}

func allocMem() []int {
	stuff := make([]int, 0, 1000000*1000)
	writeMemStats()
	fmt.Printf("%d\n", len(stuff))
	return stuff
}

func writeMemStats() {
	runtime.GC()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	memStatsString, err := json.MarshalIndent(memStats, "", "  ")
	if err != nil {
		panic(err)
	}

	fileName, err := getNextMemStatsFileName()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("data/"+fileName, []byte(memStatsString), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Wrote memstats to", "data/"+fileName)
}

func getNextMemStatsFileName() (string, error) {
	entries, err := os.ReadDir(dataDirectory)
	if err != nil {
		return "", err
	}

	memStatsFileNameFormat := "memstats-%03d.json"

	suffix := 0
	for {
		found := false
		for _, entry := range entries {
			if entry.Name() == fmt.Sprintf(memStatsFileNameFormat, suffix) {
				found = true
				break
			}
		}

		if !found {
			break
		}

		suffix++
	}

	return fmt.Sprintf(memStatsFileNameFormat, suffix), nil
}
