package csvoperations

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"stock_Project/common"
	"strconv"
	"time"
)

func ReadCSV(filePath string) ([]common.Info, error) {
	sol := []common.Info{}
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("Error opening file %s\n Error : %v\n", filePath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	data, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading file %s\n Error : %v\n", filePath, err)
	}
	for _, line := range data {
		if len(line) != 3 {
			log.Fatal("Invalid data from csv file ")
		}
		pret, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			fmt.Printf("Price invalid%v\n", err)
			return nil, err
		}
		timp, err := time.Parse(common.TimeLayout, line[1])
		if err != nil {
			fmt.Printf("Timestamp invalid %v\n", err)
			return nil, err
		}
		var value = common.Info{Timestamp: timp, Name: line[0], Price: pret}
		sol = append(sol, value)
	}
	return sol, nil
}

func WriteCSV(filePath string, data []common.Info) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error Creating File %v", err)
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, info := range data {
		record := []string{
			info.Name,
			info.Timestamp.Format(common.TimeLayout),
			strconv.FormatFloat(info.Price, 'f', 2, 64),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record to CSV: %v", err)
		}
	}

	return nil
}
