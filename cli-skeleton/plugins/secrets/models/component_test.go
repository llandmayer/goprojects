package models_test

import (
	"cli-skeleton/plugins/secrets/models"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildComponent_Success(t *testing.T) {
	mockPath := "terragrunt/aws/sykes-aws-non-prod/eu-west-1/staging/apps/checkout/encrypted_variables.txt"
	expectedComponentInfo := &models.Component{
		Account:       "sykes-aws-non-prod",
		Region:        "eu-west-1",
		Environment:   "staging",
		ComponentType: "apps",
		ComponentName: "checkout",
	}

	componentInfo, err := models.BuildComponent(mockPath)
	assert.Nil(t, err, "Expected no error during component retrieval")
	assert.Equal(t, expectedComponentInfo, componentInfo, "Component Info does not match the expected.")
}

func TestBuildComponent_CantBuildComponent(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "Path with terragrunt cache", input: "terragrunt/aws/sykes-aws-non-prod/.terragrunt-cache/eu-west-1/staging/apps/checkout/terragrunt.hcl"},
		{name: "Path with not enough elements", input: "terragrunt/terragrunt.hcl"},
		{name: "Path with no terragrunt", input: "teratera/aws/sykes-aws-non-prod/eu-west-1/staging/apps/checkout/terragrunt.hcl"},
	}
	for _, tc := range tests {
		componentInfo, err := models.BuildComponent(tc.input)
		assert.Error(t, err, "Test: %s - Unexpected result, an error should occur here: %v", tc.name, err)
		assert.Empty(t, componentInfo, "Test: %s - This should be empty", tc.name)
	}

}

func TestBuildComponent_ValidateComponent(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "Malform component - unsuported environment", input: "terragrunt/aws/sykes-aws-non-prod/eu-west-1/stag/apps/checkout/encrypted_variables.txt"},
		{name: "Malform component - unsuported region", input: "terragrunt/aws/sykes-aws-non-prod/eu-west-3/stag/apps/checkout/encrypted_variables.txt"},
	}

	for _, tc := range tests {
		component, err := models.BuildComponent(tc.input)
		assert.Nil(t, component, "component should be nil")
		assert.Error(t, err, "error validating component", err)
	}
}

func setup(t *testing.T, component models.Component) (string, string) {
	t.Helper()

	baseTempDir := t.TempDir()

	expectedPath := filepath.Join(baseTempDir, fmt.Sprintf("terragrunt/aws/%s/%s/%s/%s/%s",
		component.Account,
		component.Region,
		component.Environment,
		component.ComponentType,
		component.ComponentName))

	if err := os.MkdirAll(expectedPath, 0755); err != nil {
		t.Fatalf("Failed to create nested directories: %v", err)
	}

	if _, err := os.Create(fmt.Sprintf("%s/terragrunt.hcl", expectedPath)); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	return baseTempDir, expectedPath
}

func teardown(t *testing.T, path string) {
	t.Helper()

	if err := os.RemoveAll(filepath.Dir(path)); err != nil {
		t.Fatalf("Failed to delete directory: %v", err)
	}
}

func TestFindComponentPath(t *testing.T) {
	component := models.Component{
		Account:       "sykes-aws-non-prod",
		Region:        "eu-west-1",
		Environment:   "staging",
		ComponentType: "apps",
		ComponentName: "checkout",
	}

	baseTempDir, expectedPath := setup(t, component)
	defer teardown(t, expectedPath)

	expectedPath = filepath.Join(baseTempDir, "terragrunt/aws", component.Account, component.Region, component.Environment, component.ComponentType, component.ComponentName)
	path, err := component.FindComponentPath(baseTempDir, "terragrunt.hcl")
	if err != nil {
		t.Errorf("Failed to find the path of component %v - err: %v", component, err)
	}

	// Asserts
	assert.Equal(t, expectedPath, path, "Testing FindComponentPath, should return %s", expectedPath)
}
