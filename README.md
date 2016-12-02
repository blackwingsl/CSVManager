# CSVManager
CSVManager

```
csvManager := new(CSVManager.CSVManager)
csvManager.Load("csv/init_mercenary.csv")
value, found := csvManager.List.GetValue("init_mercenary", "quality", "5", "id")
if found {
  fmt.Println("list.GetValue:", value)
}
```
