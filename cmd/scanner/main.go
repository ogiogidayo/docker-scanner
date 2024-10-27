package main

import (
	"fmt"
	"github.com/ogiogidayo/docker-scanner/internal/handler"
	"github.com/ogiogidayo/docker-scanner/internal/usecase"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the Dockerfile.")
		return
	}

	dockerfilePath := os.Args[1]

	u := usecase.NewDockerfileUsecase()
	h := handler.NewDockerfileHandler(u)

	h.Handle(dockerfilePath)

	return
}
