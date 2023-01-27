package cast

import (
	"fmt"
	"testing"

	"github.com/go-zoox/testify"
)

func TestToMap(t *testing.T) {
	type User struct {
		ID   int64  `cast:"i_d"`
		Name string `cast:"n_ame"`
	}

	user := User{
		ID:   1,
		Name: "Zero",
	}

	userMap, _ := ToMap(user)
	testify.Equal(t, user.ID, userMap["ID"].(int64))
	testify.Equal(t, user.Name, userMap["Name"].(string))

	userMap, _ = ToMap(&user)
	testify.Equal(t, user.ID, userMap["ID"].(int64))
	testify.Equal(t, user.Name, userMap["Name"].(string))

	fmt.Println("userMap:", userMap)

	userMap, _ = ToMap(&user, "cast")
	testify.Equal(t, user.ID, userMap["i_d"].(int64))
	testify.Equal(t, user.Name, userMap["n_ame"].(string))

	fmt.Println("userMap:", userMap)
}

func TestToMapNest(t *testing.T) {
	type User struct {
		ID      int64  `cast:"i_d"`
		Name    string `cast:"n_ame"`
		Address struct {
			Road string
			Port int
		}
	}

	user := User{
		ID:   1,
		Name: "Zero",
		Address: struct {
			Road string
			Port int
		}{
			Road: "xxx",
			Port: 666,
		},
	}

	userMap, _ := ToMap(user)
	testify.Equal(t, user.ID, userMap["ID"].(int64))
	testify.Equal(t, user.Name, userMap["Name"].(string))

	fmt.Println("userMap:", userMap)
}
