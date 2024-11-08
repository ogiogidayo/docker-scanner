package usecase

import (
	"bufio"
	"context"
	"github.com/ogiogidayo/docker-scanner/internal/domain"
	"os"
	"path/filepath"
	"strings"
)

type DockerfileUsecase interface {
	ParseDockerfile(ctx context.Context, filePath string) (*domain.DockerfileInfo, error)
	FindDockerfiles(ctx context.Context) ([]string, error)
}

type dockerfileUsecase struct{}

func NewDockerfileUsecase() *dockerfileUsecase {
	return &dockerfileUsecase{}
}

func (u *dockerfileUsecase) ParseDockerfile(_ context.Context, filePath string) (*domain.DockerfileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	var from, cmd, run string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "FROM ") {
			from = strings.TrimPrefix(line, "FROM ")
		} else if strings.HasPrefix(line, "CMD ") {
			cmd = strings.TrimPrefix(line, "CMD ")
		} else if strings.HasPrefix(line, "RUN") {
			run = strings.TrimPrefix(line, "RUN")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	a := domain.DockerfileInfo{From: from, Cmd: cmd, Run: run}

	npms := domain.NewNpmPackages()
	npms.ExtractNpmPackages(a)
	npms.PrintNpmPackages()

	return &domain.DockerfileInfo{From: from, Cmd: cmd, Run: run}, nil
}

func (u *dockerfileUsecase) FindDockerfiles(_ context.Context) ([]string, error) {
	var dockerfiles []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "Dockerfile" {
			dockerfiles = append(dockerfiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dockerfiles, nil
}
