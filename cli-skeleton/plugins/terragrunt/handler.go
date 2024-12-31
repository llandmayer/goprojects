package terragrunt

type TerragruntHandler struct{}

func NewTerragruntHandler() *TerragruntHandler {
	return &TerragruntHandler{}
}

func (h *TerragruntHandler) Plan() ([]string, error) {
	paths := []string{"super", "plan", "happening"}
	return paths, nil
}
