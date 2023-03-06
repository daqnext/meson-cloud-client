package main

import (
    "context"
    "log"
    "os"
    "path/filepath"

    "daqnext/meson-cloud-client/api"
    "daqnext/meson-cloud-client/daemon"
    "daqnext/meson-cloud-client/logger"

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

func parserIpfsCmd(path string) (string, error) {

    relpath, exist, err := RelPathAndCheck(BINARY_DIR, path)
    if err != nil {
        logger.L.Panicw("Failed to read ifps cmd", "err", err.Error())
    }
    if !exist {
        log.Panicln("ifps cmd not exists", path)
    }

    return relpath, err
}

func RelPathAndCheck(DIR_PREFIX, path string) (string, bool, error) {
    //check abs path or relative path
    is_abs := filepath.IsAbs(path)
    if is_abs {
        exist, err := pathExist(path)
        if err != nil {
            return "", false, err
        }
        return path, exist, nil
    }

    logger.L.Debugln("new path", path)

    // workding dir
    path1, exist, err := toRelPath(path)
    if err != nil {
        return "", false, err
    }

    if exist {
        path = path1
        return path, exist, nil
    }

    // binary dir
    path = filepath.Join(DIR_PREFIX, path)
    path, exist, err = toRelPath(path)
    if err != nil {
        return "", false, err
    }

    return path, exist, nil
}

func toRelPath(path string) (string, bool, error) {
    // to rel directory
    path, err := filepath.Abs(path)
    if err != nil {
        return "", false, err
    }
    exist, err := pathExist(path)
    if err != nil {
        return "", false, err
    }
    return path, exist, nil
}

func pathExist(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
