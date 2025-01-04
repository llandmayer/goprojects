package secrets

import "cli-skeleton/plugins/secrets/models"

type SecretsAPI interface {
	ListSecrets() ([]models.Component, error)
	AddSecret(key, value string) error
}

type secretsAPIImpl struct {
	handler *SecretsHandler
}

func NewSecretsAPI(handler *SecretsHandler) SecretsAPI {
	return &secretsAPIImpl{
		handler: handler,
	}
}

func (api *secretsAPIImpl) ListSecrets() ([]models.Component, error) {
	return api.handler.ListSecrets()
}

func (api *secretsAPIImpl) AddSecret(key, value string) error {
	return api.handler.AddSecret(key, value)
}
