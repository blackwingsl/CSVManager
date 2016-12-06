package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/faint/CSVManager"
)

func main() {
	fmt.Println("检测CSV读取合法性")
	csvManager := new(CSVManager.CSVManager)

	dirList, e := ioutil.ReadDir("./csv/")
	if e != nil {
		fmt.Println("read dir error")
		return
	}

	count := 0
	for _, v := range dirList {
		e := csvManager.Load("./csv/" + v.Name())
		if e != nil {
			fmt.Println(e)
		}
		csvName := strings.Split(v.Name(), ".")[0]
		_, result := csvManager.List.GetCSV(csvName)
		if !result {
			fmt.Println(csvName, ":", result)
			count++
		}
	}
	fmt.Println("检测完毕，共有", count, "文件读取失败！")
}
