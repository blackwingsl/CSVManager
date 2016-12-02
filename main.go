package main

import (
	"CSVManager/CSVManager"
	"fmt"
	"path"
)

func main() {
	csvManager := new(CSVManager.CSVManager)
	csvManager.Load("csv/init_mercenary.csv")
	fmt.Println(csvManager)
	fmt.Println(path.Base("CSVManager/CSVContent.go"))
	// fmt.Println(path.Ext("/a/b/c/bar.css"))
	// fmt.Println(path.Split("static/myfile.css"))
}
