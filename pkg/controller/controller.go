package controller

import (
	"os"

	"github.com/dark-shade/go-gen/pkg/changelog"
	"github.com/dark-shade/go-gen/pkg/ci"
	"github.com/dark-shade/go-gen/pkg/dir"
	"github.com/dark-shade/go-gen/pkg/dockerfile"
	"github.com/dark-shade/go-gen/pkg/gitignore"
	"github.com/dark-shade/go-gen/pkg/makefile"
	"github.com/dark-shade/go-gen/pkg/module"
	"go.uber.org/zap"
)

// CreateStructure creates the files and directories and the project structure
func CreateStructure(logger *zap.Logger) {
	path, err := os.Getwd()
	if err != nil {
		logger.Error("unable to get the current working directory", zap.Error(err))
	}

	logger.Info("creating CHANGELOG.md")
	if err := changelog.CreateChangeLog(path); err != nil {
		logger.Error("changelog file not created", zap.Error(err))
	}

	logger.Info("creating Jenkinsfile")
	if err := ci.CreateJenkinsfile(path); err != nil {
		logger.Error("jenkinsfile not created", zap.Error(err))
	}

	logger.Info("creating directory structure")
	if errs := dir.CreateDirectoryStructure(path); len(errs) != 0 {
		for _, err := range errs {
			logger.Error("directory structure not created", zap.Error(err))
		}
	}

	logger.Info("creating Dockerfile")
	if err := dockerfile.CreateDockerfile(path); err != nil {
		logger.Error("dockerfile not created", zap.Error(err))
	}

	logger.Info("creating gitignore")
	if err := gitignore.CreateGitIgnore(path); err != nil {
		logger.Error("gitignore not created", zap.Error(err))
	}

	logger.Info("creating Makefile")
	if err := makefile.CreateMakefile(path); err != nil {
		logger.Error("makefile not created", zap.Error(err))
	}

	logger.Info("initializing go modules")
	if err := module.InitiateGoMod(); err != nil {
		logger.Error("go mod not initialized", zap.Error(err))
	}
}
