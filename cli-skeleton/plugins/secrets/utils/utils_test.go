package utils_test

import (
	"cli-skeleton/plugins/secrets/utils"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFilePath_Success(t *testing.T) {
	dir, err := os.MkdirTemp("", "findFilePath_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	filePath := filepath.Join(dir, "encrypted_variables.txt")
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	paths, err := utils.FindFilePath(dir, "encrypted_variables.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(paths) != 1 || paths[0] != filePath {
		t.Errorf("Expected to find path %s, got %v", filePath, paths)
	}
}

func TestFindFilePath_NotFound(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		file string
	}{
		{name: "Empty directory", dir: "", file: "encrypted_variables.txt"},
		{name: "Non-existent directory", dir: "/path/that/doesnt/exist", file: "encrypted_variables.txt"},
		{name: "Wrong file extension", dir: ".", file: "encrypted_variables.wrong"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			paths, err := utils.FindFilePath(tc.dir, tc.file)
			if err == nil {
				t.Errorf("Expected error, got nil")
			} else if !strings.Contains(err.Error(), "file not found") {
				t.Errorf("Expected error containing 'file not found', got: %v", err)
			}
			if len(paths) > 0 {
				t.Errorf("Expected no paths found, got %v", paths)
			}
		})
	}
}

func TestExistInSlice(t *testing.T) {
	mockPath := "Repos/tmp/sykes-cli//aws/test/eu-west-1/staging/apps/test/terragrunt.hcl"
	parts := strings.Split(mockPath, "/")
	assert.True(t, utils.ExistInSlice(parts, "aws"), "Check if there is terragrunt in the path.")
}

func TestTerragruntRootPath(t *testing.T) {
	cases := []struct {
		name  string
		setup func() (string, string)
	}{
		{
			name: "valid terragrunt root path",
			setup: func() (string, string) {
				tempDir, err := os.MkdirTemp("", "tmp")
				if err != nil {
					t.Fatal(err)
				}

				terragruntDir := filepath.Join(tempDir, "terragrunt")
				if err := os.MkdirAll(terragruntDir, os.ModePerm); err != nil {
					t.Fatal(err)
				}

				return tempDir, terragruntDir
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			root, expected := c.setup()
			defer os.RemoveAll(root)

			cwd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			defer os.Chdir(cwd)

			if err := os.Chdir(root); err != nil {
				t.Fatal(err)
			}

			path, err := utils.TerragruntRootPath()
			if err != nil {
				t.Fatal(err)
			}

			expectedPath, err := filepath.EvalSymlinks(expected)
			if err != nil {
				t.Fatal(err)
			}

			actualPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				t.Fatal(err)
			}

			expectedPath = filepath.Clean(expectedPath)
			actualPath = filepath.Clean(actualPath)

			assert.Equal(t, expectedPath, actualPath)
		})
	}
}
