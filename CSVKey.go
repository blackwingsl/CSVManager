package CSVManager

// CSVKeys ...
type CSVKeys struct {
	Value []string
}

// GetN ...
func (key *CSVKeys) GetN(fieldName string) (int, bool) {
	for k, v := range key.Value {
		if v == fieldName {
			return k, true
		}
	}
	return 0, false
}
