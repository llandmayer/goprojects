package terragrunt

type TerragruntAPI interface {
	Plan() ([]string, error)
}

type terragruntAPIImpl struct {
	handler *TerragruntHandler
}

func NewTerragruntAPI(handler *TerragruntHandler) TerragruntAPI {
	return &terragruntAPIImpl{
		handler: handler,
	}
}

func (api *terragruntAPIImpl) Plan() ([]string, error) {
	return api.handler.Plan()
}
