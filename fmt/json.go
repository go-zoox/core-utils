package fmt

import (
	"encoding/json"
	"fmt"
	gofmt "fmt"
	gostrings "strings"
)

// PrintJSON prints the JSON representation of the object.
func PrintJSON(vs ...interface{}) error {
	str := PrettyJSON(vs...)

	if _, err := gofmt.Println(str); err != nil {
		return err
	}

	return nil
}

// PrettyJSON returns the JSON representation of the object.
func PrettyJSON(vs ...interface{}) string {
	str := new(gostrings.Builder)
	for _, v := range vs {
		j, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Sprintf("Error: %s", err.Error())
		}

		str.WriteString(string(j))
		str.WriteString(" ")
	}

	return str.String()
}
