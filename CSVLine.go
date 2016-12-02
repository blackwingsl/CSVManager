package CSVManager

import "strings"

// CSVLine ...
type CSVLine struct {
	Key *Keys
	Values []string
}

// New ...
func (line *CSVLine) New(b []byte) CSVLine {
	return CSVLine{Values: strings.Split(string(b), ",")}
}

// GetValue ...
func (line *CSVLine) GetValue(key string) (string ,bool) {
	n, result := line.Key.GetN(key)
	if result {
		return line.Values[n], true
	}
	
	return "", false

}
