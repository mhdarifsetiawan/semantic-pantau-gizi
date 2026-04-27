package sliceutil

func EnsureSlice[T any](slice []T) []T {
	if slice == nil {
		return []T{}
	}
	return slice
}

func ToStringInterfaceSlice(slice []string) []interface{} {
	interfaces := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaces[i] = v
	}
	return interfaces
}
