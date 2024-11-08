package handler

import (
	"context"
	"fmt"
	"github.com/ogiogidayo/docker-scanner/internal/usecase"
)

type DockerfileHandler struct {
	Usecase usecase.DockerfileUsecase
}

func NewDockerfileHandler(u usecase.DockerfileUsecase) *DockerfileHandler {
	return &DockerfileHandler{Usecase: u}
}

func (h *DockerfileHandler) Handle(ctx context.Context) error {
	dockerfiles, err := h.Usecase.FindDockerfiles(ctx)
	if err != nil {
		err := fmt.Errorf("Error: error in find Dockerfile %s\n", err)
		return err
	}

	for _, dockerfile := range dockerfiles {
		info, err := h.Usecase.ParseDockerfile(ctx, dockerfile)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Printf("FROM: %s\nCMD: %s\nRUN: %s\n", info.From, info.Cmd, info.Run)
	}
	return nil
}
