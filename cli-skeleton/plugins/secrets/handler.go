package secrets

import "fmt"

type SecretsHandler struct{}

func NewSecretsHandler() *SecretsHandler {
	return &SecretsHandler{}
}

func (h *SecretsHandler) ListSecrets() ([]string, error) {
	paths := []string{"secret1", "secret2", "secret3"}
	return paths, nil
}

func (h *SecretsHandler) AddSecret(key, value string) error {
	fmt.Printf("Added secret: %s = %s\n", key, value)
	return nil
}
