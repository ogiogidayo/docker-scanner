package domain

import (
	"fmt"
	"strings"
)

type DockerfileInfo struct {
	From string
	Cmd  string
	Run  string
}

type NpmPackage string

type NpmPackages struct {
	np []NpmPackage
}

func NewNpmPackages() *NpmPackages {
	return &NpmPackages{}
}

func isOption(s string) bool {
	return strings.HasPrefix(s, "-")
}

func (n *NpmPackages) ExtractNpmPackages(dockerfile DockerfileInfo) {
	commandParts := strings.Fields(dockerfile.Run)

	npmIndex := -1
	installIndex := -1

	for i, part := range commandParts {
		if part == "npm" {
			npmIndex = i
		}
		if part == "install" && npmIndex != -1 && i > npmIndex {
			installIndex = i
			break
		}
	}

	if npmIndex != -1 && installIndex != -1 {
		for _, pkg := range commandParts[installIndex+1:] {
			if pkg == "&&" {
				break
			}
			if !isOption(pkg) && pkg != "" {
				n.np = append(n.np, NpmPackage(pkg))
			}
		}
	}
}

func (n *NpmPackages) PrintNpmPackages() {

	for _, pkg := range n.np {
		fmt.Println(pkg)
	}
}
