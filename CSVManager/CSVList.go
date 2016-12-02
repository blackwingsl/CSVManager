package CSVManager

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

// CSVList ...
type CSVList struct {
	CSVs []CSVContent
}

// isExist check if the csv already in CSVs
// return the position(int) in array and isExist(bool)
func (l *CSVList) isExist(name string) (int, bool) {
	for k, v := range l.CSVs {
		if v.Name == name {
			return k, true
		}
	}
	return 0, false
}

// Update target CSV
func (l *CSVList) Update(pathAndFilename string) error {
	filename := strings.Split(path.Base(pathAndFilename), ".")[0]

	file, e := ioutil.ReadFile(pathAndFilename)
	if e != nil {
		return e
	}

	// 检查是否已存在同名csv
	n, isExist := l.isExist(filename)
	if isExist { // 如果存在，则更新CSVList里的内容
		fmt.Println("csv Exist:", n)
	} else { // 如果不存在，则创建新的CSVContent，并加入CSVList
		csvContent := new(CSVContent)
		content := csvContent.New(filename, file)
		l.CSVs = append(l.CSVs, content)
	}
	fmt.Println(l.CSVs)
	return nil
}
