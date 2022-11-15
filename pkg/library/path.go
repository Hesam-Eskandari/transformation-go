package library

import (
	"os"
	"path/filepath"
)

func getProjectRootPath() (string, error) {
	return os.Getwd()
}

type configPath struct {
	inputPath          string
	outputPath         string
	transformationPath string
}

type ConfigPath interface {
	GetInputPath() string
	GetOutputPath() string
	GetTransformationPath() string
}

func GetConfigPath(configName string) ConfigPath {
	projectPath, _ := getProjectRootPath()
	configParentDirectory := "config"
	return &configPath{
		inputPath:          filepath.Join(projectPath, configParentDirectory, configName, "input.json"),
		outputPath:         filepath.Join(projectPath, configParentDirectory, configName, "output.json"),
		transformationPath: filepath.Join(projectPath, configParentDirectory, configName, "transformation.json"),
	}
}

func (cp *configPath) GetInputPath() string {
	return cp.inputPath
}

func (cp *configPath) GetOutputPath() string {
	return cp.outputPath
}

func (cp *configPath) GetTransformationPath() string {
	return cp.transformationPath
}
