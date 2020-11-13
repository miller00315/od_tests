package addresses

import (
	"strings"
)

// TypeAddress verifica se o endereço tem um tipo valido e o retorna
func TypeAddress(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"} // Utilizar Chaves para o slice

	minimumCase := strings.ToLower(address)

	firstWord := strings.Split(minimumCase, " ")[0]

	addressHasValidType := false

	for _, value := range validTypes {
		if value == firstWord {
			addressHasValidType = true
			break
		}
	}

	if addressHasValidType {
		return strings.Title(firstWord)
	}

	return "Invalid type"
}
