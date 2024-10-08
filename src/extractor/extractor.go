package main

import (
	"log"
	"math/rand"
	"os"
	"sort"
	"stock_Project/common"
	"stock_Project/csvoperations"
	"time"
)

func randomize(data []common.Info) []common.Info {
	if len(data) < 10 {
		return data

	}
	sort.Sort(common.Sorter(data))
	rand.Seed(time.Now().UnixNano())
	sol := []common.Info{}
	startingpos := rand.Intn(len(data) - 10)
	for i := startingpos; i < startingpos+10; i++ {
		sol = append(sol, data[i])
	}
	return sol
}

func main() {
	allData := []common.Info{}
	if len(os.Args) < 2 {
		log.Fatal("no files provided")
	}

	for _, file := range os.Args[1:] {
		rez, err := csvoperations.ReadCSV(file)
		if err != nil {
			log.Fatal("error encountered", err)
		}
		allData = append(allData, rez...)
	}
	sol := randomize(allData)
	err := csvoperations.WriteCSV(common.OutputFile, sol)
	if err != nil {
		log.Fatal(err)
	}
	common.Printdata(sol)
}
