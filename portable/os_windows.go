package portable

import (
    "os"
    "os/exec"
    "os/signal"
    "syscall"
    "strconv"

    "daqnext/meson-cloud-client/logger"
)

func CmdGen(exePath string, arg string, env string) (*exec.Cmd, error) {
    exePath = exePath + ".exe"
    logger.L.Debugln(exePath)
    if _, err := os.Stat(exePath); err != nil {
        return nil, err
    }
    cmd := exec.Command(exePath, arg)
    if env != "" {
        newEnv := append(os.Environ(), env)
        cmd.Env = newEnv
    }
    return cmd, nil
}

// kill the app pid with his children pid
func CmdKill(cmd *exec.Cmd) error {
    if cmd != nil {
        cmd := exec.Command("taskkill.exe", "/PID", strconv.Itoa(cmd.Process.Pid), "/T", "/F")
        return cmd.Run()
    }
    return nil
}

func SysSignalFunc(exitFunc func(os.Signal)) {
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