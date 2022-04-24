package object

// func WalkByDeepFirst[K comparable, V any](object map[K]V, fn func(V, K) bool) map[K]V {
// 	result := make(map[K]V)

// 	for key, value := range object {
// 		switch value.(type) {
// 		case map[K]V:
// 			result[key] = WalkByDeepFirst(value.(map[K]V), fn)
// 		default:
// 			if fn(value, key) {
// 				result[key] = value
// 			}
// 		}
// 	}

// 	return result
// }
