package handler

import (
	"fmt"
	"github.com/ogiogidayo/docker-scanner/internal/usecase"
)

type DockerfileHandler struct {
	Usecase usecase.DockerfileUsecase
}

func NewDockerfileHandler(u usecase.DockerfileUsecase) *DockerfileHandler {
	return &DockerfileHandler{Usecase: u}
}

func (h *DockerfileHandler) Handle(filePath string) {
	info, err := h.Usecase.ParseDockerfile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("FROM: %s\nCMD: %s\n", info.From, info.Cmd)
}
