package fmt

import (
	"encoding/json"
	gofmt "fmt"
)

func PrintJSON(v interface{}) {
	j, _ := json.MarshalIndent(v, "", "  ")
	gofmt.Println(string(j))
}
