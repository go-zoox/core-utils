package object

import "encoding/json"

// StructToMap converts struct to map
func StructToMap[K comparable, V any](structValue any) (mapValue map[K]V, err error) {
	// Convert struct to bytes
	bytes, err := json.Marshal(structValue)
	if err != nil {
		return nil, err
	}

	// Convert bytes to map
	err = json.Unmarshal(bytes, &mapValue)
	if err != nil {
		return nil, err
	}

	return mapValue, nil
}

// MustStructToMap converts struct to map and panics if there is an error
func MustStructToMap[K comparable, V any](structValue any) map[K]V {
	mapValue, _ := StructToMap[K, V](structValue)
	return mapValue
}
