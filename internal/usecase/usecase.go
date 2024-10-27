package usecase

import (
	"bufio"
	"github.com/ogiogidayo/docker-scanner/internal/domain"
	"os"
	"strings"
)

type DockerfileUsecase interface {
	ParseDockerfile(filePath string) (*domain.DockerfileInfo, error)
}

type dockerfileUsecase struct{}

func NewDockerfileUsecase() DockerfileUsecase {
	return &dockerfileUsecase{}
}

func (u *dockerfileUsecase) ParseDockerfile(filePath string) (*domain.DockerfileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var from, cmd string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "FROM ") {
			from = strings.TrimPrefix(line, "FROM ")
		} else if strings.HasPrefix(line, "CMD ") {
			cmd = strings.TrimPrefix(line, "CMD ")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &domain.DockerfileInfo{From: from, Cmd: cmd}, nil
}
