package common

import (
	"fmt"
	"os"
	"time"
)

type Info struct {
	Timestamp time.Time
	Name      string
	Price     float64
}

var TimeLayout = "02-01-2006"
var OutputFile = "data/output.csv"
var Location, _ = os.Getwd()

type Sorter []Info

func (a Sorter) Len() int { return len(a) }
func (a Sorter) Less(i, j int) bool {
	if a[i].Timestamp.Equal(a[j].Timestamp) {
		if a[i].Price == a[j].Price {
			return a[i].Name < a[j].Name
		}
		return a[i].Price < a[j].Price
	}
	return a[i].Timestamp.Before(a[j].Timestamp)
}
func (a Sorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Printdata(date []Info) {
	for _, val := range date {
		fmt.Printf("Timestamp: %s, Name: %s, Price: %.2f\n", val.Timestamp.Format(time.RFC822), val.Name, val.Price)
	}
}
