# stock_project
 A simple gin project that processes stock data and predicts behaviour.
How to run:
go mod 
go run main.go
If you want to run tasks separately you can.
For task1:
task1 can process multiple csv file and sort the stock data by the date
go run task1.go a.csv b.csv
For task2:
task2 takes the data from the output from task1(which is stored in OutputFile which is declared in common.go) and prediction the following 3 days
go run task2.go
