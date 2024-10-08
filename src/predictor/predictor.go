package main

import (
	"log"
	"stock_Project/common"
	"stock_Project/csvoperations"
	"time"
)

func findmax(data []common.Info) int {
	sol := 0
	max := data[0].Price
	for i := 1; i < len(data); i++ {
		if max < data[i].Price {
			max = data[i].Price
			sol = i
		}
	}
	return sol
}

func find2ndmax(data []common.Info) (string, float64) {
	index := findmax(data)
	max := data[0].Price - 1
	var name string
	for _, val := range data {
		if max < val.Price && val.Price != data[index].Price {
			name = val.Name
			max = val.Price
		}
	}
	return name, max
}

func simplePredict(data []common.Info) []common.Info {
	timestamp := data[len(data)-1].Timestamp
	nprice := data[len(data)-1].Price
	name, n1price := find2ndmax(data)
	n2price := n1price + (nprice-n1price)/2
	n3price := n2price + (n1price-n2price)/4
	sol := []common.Info{
		{timestamp.Add(24 * time.Hour), name, n1price},
		{timestamp.Add(48 * time.Hour), name, n2price},
		{timestamp.Add(72 * time.Hour), name, n3price},
	}
	return sol
}

func main() {
	sol := []common.Info{}
	sol, err := csvoperations.ReadCSV(common.OutputFile)
	if err != nil {
		log.Fatal(err)
	}
	predicted := simplePredict(sol)
	sol = append(sol, predicted...)
	common.Printdata(sol)
}
