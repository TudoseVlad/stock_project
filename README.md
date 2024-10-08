# stock_project
A simple gin project that processes stock data and predicts behaviour.

**_How to run:_**

**go mod**

**go run main.go**

If you want to run tasks separately you can.

**_For Extractor:_**

Extractor can process multiple csv file and sort the stock data by the date

**go run extracter.go a.csv b.csv**

**_For Predictor:_**

Predictor takes the data from the output from task1(which is stored in OutputFile which is declared in common.go) and prediction the following 3 days

**go run predictor.go**
