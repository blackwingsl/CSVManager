package main

import (
	"CSVManager/CSVManager"
	"fmt"
)

func main() {
	csvManager := new(CSVManager.CSVManager)
	csvManager.Load("csv/init_mercenary.csv")
	s1, found := csvManager.List.GetValue("init_mercenary", "quality", "5", "id")
	if found {
		fmt.Println("list.GetValue:", s1)
	}
}
