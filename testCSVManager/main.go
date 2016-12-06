package main

import (
	"fmt"

	"github.com/faint/CSVManager"
)

func main() {
	fmt.Println("CSVManager Start...")

	csvManager := new(CSVManager.CSVManager)
	csvManager.Load("../csv/params.csv")

  // fmt.Println(csvManager.List.CSVs)
	v, f := csvManager.List.Get("params", "name", "BattleConfiguration", "desc")
	fmt.Println("v:",v,f)

  vs, fs := csvManager.List.Gets("params", "name", "BattleConfiguration", "desc")
  fmt.Println("vs:",vs,fs)


	fmt.Println("CSVManager End...")
}
