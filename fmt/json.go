package fmt

import (
	"encoding/json"
	gofmt "fmt"
	"strings"
)

func PrintJSON(vs ...interface{}) error {
	str := new(strings.Builder)
	for _, v := range vs {
		j, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}

		str.WriteString(string(j))
		str.WriteString(" ")
	}

	_, err := gofmt.Println(str)
	return err
}
