package CSVManager

// Keys ...
type Keys struct {
  Value []string
}

// GetN ...
func (key *Keys) GetN(fieldName string) (int, bool){
  for k,v := range key.Value {
    if v == fieldName {
      return k,true
    }
  }
  return 0, false
}
