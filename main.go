package main

import (
    "context"
    "log"

    "daqnext/meson-cloud-client/api"
    "daqnext/meson-cloud-client/daemon"
    "daqnext/meson-cloud-client/logger"
    "github.com/spf13/viper"
)

func main() {
    // Read Config
    viper.SetConfigName("config") // name of config file (without extension)
    viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
    viper.AddConfigPath(".")      // path to look for the config file in
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

    logger.L.Debugln("cmd", ipfsCfg.IpfsCmd, "dataRoot", ipfsCfg.IpfsDataRoot, "queryUrl", queryUrl, "token", token)

    ctx, cancelFunc := context.WithCancel(context.Background())
    defer cancelFunc()

    // Register & Run Daemon(s)
    ipfsDaemon := daemon.NewIpfsDaemon(&ipfsCfg)
    if err := ipfsDaemon.Start(ctx); err != nil {
        logger.L.Panicw("Failed to start the IPFS node", "err", err.Error())
    }

    // Api Jobs
    apiMgr := api.NewApiMgr(queryUrl, token, ipfsDaemon)
    apiMgr.Run()
}
