package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"DockerWithGo/utils" // Adjust this path according to your module name
)

func Child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Change root directory to a new root file system
	utils.Must(syscall.Chroot("./rootfs"))
	utils.Must(os.Chdir("/"))
	// Mount proc file system in the new namespace
	utils.Must(syscall.Mount("proc", "proc", "proc", 0, ""))
	defer syscall.Unmount("proc", 0)

	utils.Must(cmd.Run())
}
