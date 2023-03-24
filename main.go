package main

import (
    "context"
    "log"
    "os"
    "path/filepath"

    "daqnext/meson-cloud-client/api"
    "daqnext/meson-cloud-client/daemon"
    "daqnext/meson-cloud-client/logger"
    "daqnext/meson-cloud-client/portable"
    "daqnext/meson-cloud-client/utils"

    "github.com/spf13/viper"
)

var BINARY_DIR string

func main() {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    BINARY_DIR = filepath.Dir(ex)

    // Read Config
    viper.SetConfigName("config")   // name of config file (without extension)
    viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
    viper.AddConfigPath(".")        // path to look for the config file in
    viper.AddConfigPath(BINARY_DIR) // path to look for the config file in

    //viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
    if err := viper.ReadInConfig(); err != nil {
        log.Panicln("Failed to read the config", err.Error())
    }

    token := viper.GetString("token")
    queryUrl := viper.GetString("queryUrl")
    logLevel := viper.GetString("logLevel")

    var ipfsCfg daemon.IpfsCfg
    if err := viper.UnmarshalKey("ipfs", &ipfsCfg); err != nil {
        logger.L.Panicw("Failed to read ipfs confg", "err", err.Error())
    }

    // Register logger
    logger.RegisterLogger(logLevel)
    defer logger.L.Sync()

    ipfsCfg.IpfsCmd, err = parserIpfsCmd(ipfsCfg.IpfsCmd)
    logger.L.Debugln("cmd", ipfsCfg.IpfsCmd, "dataRoot", ipfsCfg.IpfsDataRoot, "queryUrl", queryUrl, "token", token)

    mainCtx, mainCancel := context.WithCancel(context.Background())
    defer mainCancel()

    // Register & Run Daemon(s)
    ipfsDaemon := daemon.NewIpfsDaemon(&ipfsCfg)
    if err := ipfsDaemon.Start(mainCtx); err != nil {
        logger.L.Panicw("Failed to start the IPFS node", "err", err.Error())
    }

    GracefulExit := func (s os.Signal) {
        logger.L.Debugln("Program Exit...", s)
        mainCancel()
    }
    portable.SysSingalFunc(GracefulExit)

    // Api Jobs
    apiMgr := api.NewApiMgr(queryUrl, token, ipfsDaemon)
    apiMgr.Run(mainCtx)
}

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
