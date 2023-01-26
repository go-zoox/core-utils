package regexp

import (
	"fmt"
	"testing"

	"github.com/go-zoox/testify"
)

func TestNew(t *testing.T) {
	re, _ := New("\\d+")

	// match
	testify.Equal(t, true, re.Match("123"))
	testify.Equal(t, true, re.Match("abc123"))
	testify.Equal(t, false, re.Match("abc"))

	// find
	testify.Equal(t, "123", re.FindOne("a123b456c7"))

	all := re.Find("a123b456c7")
	testify.Equal(t, 3, len(all))
	testify.Equal(t, "123", all[0])
	testify.Equal(t, "456", all[1])
	testify.Equal(t, "7", all[2])

	all = re.FindN("a123b456c7", 2)
	testify.Equal(t, 2, len(all))
	testify.Equal(t, "123", string(all[0]))
	testify.Equal(t, "456", string(all[1]))
	all = re.FindN("a123b456c7", 1)
	testify.Equal(t, 1, len(all))

	// split
	all = re.Split("a123b456c7")
	testify.Equal(t, 4, len(all))
	testify.Equal(t, "a", all[0])
	testify.Equal(t, "b", all[1])
	testify.Equal(t, "c", all[2])
	testify.Equal(t, "", all[3])

	all = re.SplitN("a123b456c7", 2)
	testify.Equal(t, 2, len(all))
	testify.Equal(t, "a", all[0])
	testify.Equal(t, "b456c7", all[1])

	// replace
	testify.Equal(t, "a#b#c#", re.Replace("a123b456c7", "#"))
	testify.Equal(t, "a*123*b*456*c*7*", re.ReplaceFunc("a123b456c7", func(s string) string {
		return fmt.Sprintf("*%s*", s)
	}))

	// submatch
	re, _ = New("^/api/([^$]*)")
	testify.Equal(t, "/123", re.Replace("/api/123", "/$1"))
	testify.Equal(t, "/user/123", re.Replace("/api/user/123", "/$1"))
	// empty
	testify.Equal(t, "/", re.Replace("/api/", "/$1"))
	// no match
	testify.Equal(t, "/api", re.Replace("/api", "/$1"))
	testify.Equal(t, "/service", re.Replace("/service", "/$1"))

	re, _ = New("^/api/users/(\\d+)/posts/(\\d+)")
	testify.Equal(t, "/users/1/posts/3", re.Replace("/api/users/1/posts/3", "/users/$1/posts/$2"))
}
