//go:build !windows

package portable

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"daqnext/meson-cloud-client/logger"
)

const enableSid = false

func CmdGen(exePath string, arg string) (*exec.Cmd, error) {
	logger.L.Debugln(exePath)
	_, err := os.Stat(exePath)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(exePath, arg)
	if (enableSid) {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}
	return cmd, nil
}

// kill the app pid with his children pid
func CmdKill(cmd *exec.Cmd) error {
	if cmd != nil {
		if (enableSid) {
			// negtive standard for group kill
			return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		} else {
			return cmd.Process.Kill()
		}
	}
	return nil
}

func SysSingalFunc(exitFunc func(os.Signal)) {
    c := make(chan os.Signal)

    signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    go func() {
        for s := range c {
            switch s {
            case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
                logger.L.Debugln("singal", s)
                exitFunc(s)
            default:
                logger.L.Debugln("singal other", s)
            }
        }
    }()
}