package object

import "github.com/go-zoox/core-utils/cast"

// Unique retturns a unique array of objects.
func Unique[T comparable](origin []T, keyFn ...func(origin T, index int) string) []T {
	if len(keyFn) == 0 {
		return UniqueBy(origin, func(origin T, index int) string {
			return cast.ToString(origin)
		})
	}

	return UniqueBy(origin, keyFn[0])
}

// UniqueBy returns a unique array of objects by keyFn.
func UniqueBy[T comparable](origin []T, keyFn func(origin T, index int) string) []T {
	news := make([]T, 0, len(origin))
	newsMap := make(map[string]bool)
	for i, one := range origin {
		if key := keyFn(one, i); key != "" {
			if !newsMap[key] {
				newsMap[key] = true
				news = append(news, one)
			}
		}
	}

	return news
}
