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

// Update target CSV
func (list *CSVList) Update(pathAndFilename string) error {
	filename := strings.Split(path.Base(pathAndFilename), ".")[0]

	file, e := ioutil.ReadFile(pathAndFilename)
	if e != nil {
		return e
	}

	// 检查是否已存在同名csv
	n, isExist := list.isExist(filename)
	csvContent := new(CSVContent)
	if isExist { // 如果存在，则更新CSVList里的内容
		list.CSVs[n] = csvContent.New(filename, file)
	} else { // 如果不存在，则创建新的CSVContent，并加入CSVList
		content := csvContent.New(filename, file)
		list.CSVs = append(list.CSVs, content)
	}

	return nil
}

// GetValue ...
func (list *CSVList) GetValue(csvName, keyFiledName, keyFiledValue, needField string) (string, bool) {
	csv, result := list.GetCSV(csvName)
	if !result {
		return "", false
	}

	content, contentResult := csv.GetContent(keyFiledName, keyFiledValue)
	if !contentResult {
		return "", false
	}

	n, nResult := csv.GetN(needField)
	if !nResult {
		return "", false
	}

	return content.GetValue(n)
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
