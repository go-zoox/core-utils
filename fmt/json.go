package fmt

import (
	"encoding/json"
	gofmt "fmt"
	gostrings "strings"
)

// PrintJSON prints the JSON representation of the object.
func PrintJSON(vs ...interface{}) error {
	str := new(gostrings.Builder)
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
