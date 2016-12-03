package CSVManager

import "strings"

// CSVLine ...
type CSVLine struct {
	Key    *CSVKeys
	Values []string
}

// New ...
func (line *CSVLine) New(b []byte) CSVLine {
	return CSVLine{Values: strings.Split(string(b), ",")}
}

// Get ...
func (line *CSVLine) Get(key string) (string, bool) {
	n, result := line.Key.GetN(key)
	if result {
		return line.Values[n], true
	}

	return "", false
}
