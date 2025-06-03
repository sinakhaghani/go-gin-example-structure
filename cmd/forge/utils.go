package forge

import (
	"strings"
)

// ToCamelCase converts snake_case to CamelCase
func ToCamelCase(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
