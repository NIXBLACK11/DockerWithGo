package container

import (
	"os"
	"os/exec"
	"syscall"
	"DockerWithGo/utils" // Adjust this path according to your module name
)

func Run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	utils.Must(cmd.Run())
}
