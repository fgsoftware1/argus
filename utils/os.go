package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func CheckOS() string {
	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Operating System: Windows")
	case "linux":
		fmt.Println("Operating System: Linux")
	default:
		fmt.Printf("Operating System: %s\n", os)
	}

	return os
}

func Clear() {
	os := CheckOS()

	switch os {
	case "linux":
		exec.Command("clear")
	}
}
