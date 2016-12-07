package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("开始比较：")
	pathAndFilename1 := "../csv/Language_chs1.csv"
	file1, e := ioutil.ReadFile(pathAndFilename1)
	if e != nil {
		fmt.Println(e)
	}
	lines1 := bytes.Split(file1, []byte{'\n'})
	for k, v := range lines1 {
		if k > 3 {
			break
		}
		if len(v) <= 0 { // 空行过滤
			continue
		}

		fmt.Println(k,v)
		if v[len(v)-1] == '\r' {
			fmt.Println("gotcha",k)
		}
	}

	fmt.Println(file1[0:15])

	pathAndFilename2 := "../csv/Language_chs2.csv"
	file2, e := ioutil.ReadFile(pathAndFilename2)
	if e != nil {
		fmt.Println(e)
	}
	// lines2 := bytes.Split(file2, []byte{'\n', '\r'})/
	fmt.Println(file2[0:15])
	lines2 := bytes.Split(file2, []byte{'\n'})
	for k, v := range lines2 {
		if k > 3 {
			break
		}
		if len(v) <= 0 { // 空行过滤
			continue
		}

		fmt.Println(k,v)
		if v[len(v)-1] == '\r' {

			v = v[:len(v)-1]
			fmt.Println(k,v)
			fmt.Println("gotcha",k)
		}
	}

	/*
		key1 := []string{}
		for k, v := range lines1 {
			if len(v) <= 0 { // 空行过滤
				continue
			}

			// if v[len(v)-1] == '\r' || v[len(v)-1] == '\n' { // 去除尾部\r
			// 	v = v[:len(v)-2]
			// }
			//
			// if v[0] == '\r' || v[0] == '\n' { //去除头部\r
			// 	v = v[1:]
			// }

			if v[0] == '#' { // 注释行
				fmt.Println("line:",k,v)
				if v[1] == '!' && len(key1) <= 0 { // 标题行
					key1 = strings.Split(string(v[2:]), ",")
				}
				continue
			}
		}

		for _, v := range lines2 {
			if string(v) == "" {
				fmt.Println("test:", v)
			}
		}
	*/
	// fmt.Println(key1)
	// fmt.Println((start))
	// fmt.Println(string(file1[11:50]))
	// if file1[11] == '#' {
	// 	fmt.Println("equal:", file1[11])
	// } else {
	// 	two := byte('#')
	// 	fmt.Println("not equal:", file1[11], "!=", two)
	// }
	//
	// fmt.Println("before:", string(file1[:13]))
	// fmt.Println("before:", (file1[:13]))
	// fmt.Println("---")
	// fmt.Println("after:", string(file2[:13]))
	// fmt.Println("after:", (file2[:13]))

	// 检查是否已存在同名csv
	// filename := strings.Split(path.Base(pathAndFilename), ".")[0]
	// n, isExist := list.isExist(filename)
	// content := new(CSVContent)
	// if isExist { // 如果存在，则更新CSVList里的内容
	// 	list.CSVs[n] = content.New(filename, file)
	// } else { // 如果不存在，则创建新的CSVContent，并加入CSVList
	// 	contentNew := content.New(filename, file)
	// 	list.CSVs = append(list.CSVs, contentNew)
	// }
	// csvManager := new(CSVManager.CSVManager)
	// csvManager.Load("../csv/Language_chs1.csv")
	// csvManager.Load("../csv/Language_chs2.csv")
	// fmt.Println(csvManager)
}
