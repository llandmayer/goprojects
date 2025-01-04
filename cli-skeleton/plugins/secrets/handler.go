package secrets

import (
	"cli-skeleton/plugins/secrets/models"
	"cli-skeleton/plugins/secrets/utils"
	"fmt"
)

type SecretsHandler struct{}

func NewSecretsHandler() *SecretsHandler {
	return &SecretsHandler{}
}

func (h *SecretsHandler) ListSecrets() ([]models.Component, error) {
	// table := utils.NewTable([]string{"Account", "Region", "Environment", "Component_Type", "Component_Name"})
	pathsFound, err := utils.FindFilePath(".", "encrypted_variables.txt")
	if err != nil {
		return nil, fmt.Errorf("error, %w", err)
	}

	components := []models.Component{}

	for _, i := range pathsFound {
		info, err := models.BuildComponent(i)
		if info != nil && err == nil {
			components = append(components, *info)
		}
	}

	return components, nil
}

func (h *SecretsHandler) AddSecret(key, value string) error {
	fmt.Printf("Added secret: %s = %s\n", key, value)
	return nil
}
