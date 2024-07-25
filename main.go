package main

import (
	"DockerWithGo/container"
	"DockerWithGo/utils"
	"errors"
	// "fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: DockerWithGo run|child [args...]")
	}

	// fmt.Println("Checking Image...")

	// err := utils.CheckImage()

	// if err!=nil {
	// 	utils.Must(err)
	// }

	// fmt.Println("Image checked...")

	switch os.Args[1] {
	case "run":
		container.Run()
	case "child":
		container.Child()
	default:
		utils.Must(errors.New("wrong command"))
	}
}
