package fmt

import (
	"encoding/json"
	gofmt "fmt"
	gostrings "strings"
)

// PrintJSON prints the JSON representation of the object.
func PrintJSON(vs ...interface{}) error {
	str, err := PrettyJSON(vs...)
	if err != nil {
		return err
	}

	if _, err = gofmt.Println(str); err != nil {
		return err
	}

	return nil
}

// PrettyJSON returns the JSON representation of the object.
func PrettyJSON(vs ...interface{}) (string, error) {
	str := new(gostrings.Builder)
	for _, v := range vs {
		j, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return "", err
		}

		str.WriteString(string(j))
		str.WriteString(" ")
	}

	return str.String(), nil
}
