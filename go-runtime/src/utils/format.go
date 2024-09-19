package utils

import (
	"fmt"
	"strings"
)

func SliceToString(slice []int) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, num := range slice {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", num))
	}
	sb.WriteString("]")
	return sb.String()
}

func BoolToStorageQuery(value bool) string {
	if value {
		return "T"
	}
	return "F"
}
