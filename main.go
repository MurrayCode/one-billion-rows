package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Station struct {
	min   float64
	max   float64
	count int
	sum   float64
}

type Stations map[string]Station

func filter(input string) (string, float64) {
	station := strings.Split(input, ";")[0]
	temp, _ := strconv.ParseFloat(strings.Split(input, ";")[1], 64)
	return station, temp
}

func readData(path string) Stations {
	stations := make(Stations)
	file, _ := os.Open(path)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		station, temp := filter(fileScanner.Text())
		if entry, ok := stations[station]; ok {
			entry.min = min(stations[station].min, temp)
			entry.max = max(stations[station].max, temp)
			entry.count++
			entry.sum += temp
			stations[station] = entry
		} else {
			stations[station] = Station{temp, temp, 1, temp}
		}
	}
	file.Close()
	return stations
}

func main() {
	start := time.Now()

	stations := readData("measurements.txt")
	keys := make([]string, 0, len(stations))
	for k := range stations {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var s string

	for _, k := range keys {
		s = s + fmt.Sprintf("%s/%f/%f/%f\n", k, stations[k].min, stations[k].max, stations[k].sum/float64(stations[k].count))
	}

	fmt.Println(s)
	elapsed := time.Since(start)
	log.Printf("processing took: ", elapsed)
}
