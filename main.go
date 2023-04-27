package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"daqnext/meson-cloud-client/api"
	"daqnext/meson-cloud-client/daemon"
	"daqnext/meson-cloud-client/logger"
	"daqnext/meson-cloud-client/portable"
	"daqnext/meson-cloud-client/utils"
)

var BINARY_DIR string

func parserIpfsCmd(path string) (string, error) {
	relpath, exist, err := utils.RelPathAndCheck(BINARY_DIR, path)
	if err != nil {
		logger.L.Panicw("Failed to read ipfs cmd", "err", err.Error())
	}
	if !exist {
		logger.L.Panicln("ipfs cmd not exists", path)
	}

	return relpath, err
}

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	BINARY_DIR = filepath.Dir(ex)

	defaultAction := func(clictx *cli.Context) error {

		appConfig := loadConfig(BINARY_DIR)
		appCfg := appConfig.cfg
		queryUrl := appCfg.QueryUrl
		token := appCfg.Token
		logLevel := appCfg.LogLevel
		ipfsCfg := appCfg.Ipfs

		// Register logger
		logger.RegisterLogger(logLevel)
		defer logger.L.Sync()

		ipfsCfg.IpfsCmd, err = parserIpfsCmd(ipfsCfg.IpfsCmd)
		ipfsCfg.IpfsDataRoot, err = daemon.IpfsDir(ipfsCfg.IpfsDataRoot)

		logger.L.Debugln("cmd", ipfsCfg.IpfsCmd, "dataRoot", ipfsCfg.IpfsDataRoot, "queryUrl", queryUrl, "token", token)

		mainCtx, mainCancel := context.WithCancel(context.Background())
		defer mainCancel()

		// Daemon(s) Register
		ipfsDaemon := daemon.NewIpfsDaemon(&ipfsCfg)

		// Daemon(s) Init
		if exists, err := utils.PathExist(ipfsCfg.IpfsDataRoot); !exists || err != nil {
			logger.L.Debugln("ipfs Repo Init")

			if err := ipfsDaemon.Init(); err != nil {
				logger.L.Panicw("Failed to start the IPFS node", "err", err.Error())
			}
		} else {
			logger.L.Debugln("ipfs Repo Found")
		}

		// Daemon(s) Run
		if err := ipfsDaemon.Start(mainCtx); err != nil {
			logger.L.Panicw("Failed to start the IPFS node", "err", err.Error())
		}

		GracefulExit := func(s os.Signal) {
			logger.L.Debugln("Program Exit...", s)
			mainCancel()
		}
		portable.SysSignalFunc(GracefulExit)

		// Api Jobs
		apiMgr := api.NewApiMgr(queryUrl, token, ipfsDaemon)
		apiMgr.Run()

		return nil
	}

	// config app to run
	errRun := ConfigCmd(defaultAction).Run(os.Args)
	if errRun != nil {
		panic(errRun)
	}
}
