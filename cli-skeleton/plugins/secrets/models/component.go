package models

import (
	"cli-skeleton/plugins/secrets/utils"
	"fmt"
	"path/filepath"
	"reflect"
	"slices"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

type Component struct {
	Account       string `validate:"required"`
	Region        string `validate:"required,awsregion"`
	Environment   string `validate:"required,environment"`
	ComponentType string `validate:"required"`
	ComponentName string `validate:"required"`
}

var AvailableAWSRegions = map[string]struct{}{
	"us-east-1":      {},
	"eu-west-1":      {},
	"eu-west-2":      {},
	"ap-southeast-2": {},
}

func validateAWSRegion(fl validator.FieldLevel) bool {
	region := fl.Field().String()
	_, ok := AvailableAWSRegions[region]
	return ok
}

var AvailableEnvironment = map[string]struct{}{
	"dev":        {},
	"staging":    {},
	"production": {},
}

func validateEnvironment(fl validator.FieldLevel) bool {
	environment := fl.Field().String()
	_, ok := AvailableEnvironment[environment]
	return ok
}

func BuildComponent(filePath string) (*Component, error) {
	ignoreList := []string{".terragrunt-cache"}

	parts := strings.Split(filePath, "/")

	// Check if terragrunt exist in the path
	if !utils.ExistInSlice(parts, "terragrunt") {
		return nil, fmt.Errorf("file path has no terragrunt: %v", filePath)
	}

	// Excludes paths that contain item in the ignoraList
	for _, e := range ignoreList {
		if utils.ExistInSlice(parts, e) {
			return nil, fmt.Errorf("file path contains element from ignore list: %s, skipping", e)
		}
	}

	foundIndex := slices.Index(parts, "terragrunt")

	if len(parts) <= foundIndex+7 {
		return nil, fmt.Errorf("the path doesn't have enough element to build a component, %v", parts)
	}

	validate := validator.New()
	validate.RegisterValidation("awsregion", validateAWSRegion)
	validate.RegisterValidation("environment", validateEnvironment)

	component := &Component{
		Account:       parts[foundIndex+2],
		Region:        parts[foundIndex+3],
		Environment:   parts[foundIndex+4],
		ComponentType: parts[foundIndex+5],
		ComponentName: parts[foundIndex+6],
	}

	err := validate.Struct(component)
	if err != nil {
		return nil, fmt.Errorf("error validating component")
	}

	return component, nil
}

// Finds file path based on component struct.
// currentPath: Used to know from where to start looking for components. (i.e: ".")
// file: It's the file used to as an anchor to build a component. (i.e: terragrunt.hcl)
// Return absolut path of a component (i.e: /terragrunt/aws/account/region/environment/component_type/component_name)
func (c *Component) FindComponentPath(currentPath string, file string) (string, error) {
	var path string
	pathsFound, err := utils.FindFilePath(currentPath, file)

	if err != nil {
		return "", fmt.Errorf("error, %w", err)
	}

	// Finds the absolute path to the file that matches the component requested
	for _, i := range pathsFound {
		info, err := BuildComponent(i)
		if info != nil && err == nil {
			if reflect.DeepEqual(info, c) {
				path = filepath.Dir(i)
				break
			}
		}
	}

	return path, nil
}
