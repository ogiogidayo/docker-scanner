package main

import (
	"context"
	"fmt"
	"github.com/ogiogidayo/docker-scanner/internal/handler"
	"github.com/ogiogidayo/docker-scanner/internal/usecase"
)

func main() {

	ctx := context.Background()
	u := usecase.NewDockerfileUsecase()
	h := handler.NewDockerfileHandler(u)

	err := h.Handle(ctx)
	if err != nil {
		fmt.Printf("error in Handler: %s\n", err)
	}

	return
}
