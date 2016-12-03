package CSVManager

import (
	"io/ioutil"
	"path"
	"strings"
)

// CSVList ...
type CSVList struct {
	CSVs []CSVContent
}

// Get return only one line
func (list *CSVList) Get(csvName, keyFiledName, keyFiledValue, needField string) (string, bool) {
	content, result := list.GetCSV(csvName) //  读表
	if !result {
		return "", false
	}

	line, result := content.Get(keyFiledName, keyFiledValue) // 读行
	if !result {
		return "", false
	}

	value, result := line.Get(needField)
	if !result {
		return "", false
	}

	return value, true
}

// Gets return multi line
func (list *CSVList) Gets(csvName, keyFiledName, keyFiledValue, needField string) ([]string, bool) {
	content, result := list.GetCSV(csvName) //  读表
	if !result {
		return []string{}, false
	}

	lines, result := content.Gets(keyFiledName, keyFiledValue)
	if !result {
		return []string{}, false
	}

	values := []string{}
	for _, v := range lines {
		value, result := v.Get(needField)
		if result {
			values = append(values, value)
		}
	}

	return values, true
}

// GetCSV ...
func (list *CSVList) GetCSV(csvName string) (CSVContent, bool) {
	for _, v := range list.CSVs {
		if v.Name == csvName {
			return v, true
		}
	}
	return CSVContent{}, false
}

// Update target CSV
func (list *CSVList) Update(pathAndFilename string) error {
	file, e := ioutil.ReadFile(pathAndFilename)
	if e != nil {
		return e
	}

	// 检查是否已存在同名csv
	filename := strings.Split(path.Base(pathAndFilename), ".")[0]
	n, isExist := list.isExist(filename)
	content := new(CSVContent)
	if isExist { // 如果存在，则更新CSVList里的内容
		list.CSVs[n] = content.New(filename, file)
	} else { // 如果不存在，则创建新的CSVContent，并加入CSVList
		contentNew := content.New(filename, file)
		list.CSVs = append(list.CSVs, contentNew)
	}

	return nil
}

// isExist check if the csv already in CSVs
// return the position(int) in array and isExist(bool)
func (list *CSVList) isExist(name string) (int, bool) {
	for k, v := range list.CSVs {
		if v.Name == name {
			return k, true
		}
	}
	return 0, false
}
