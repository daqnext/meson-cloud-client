package api

import (
    "fmt"
    "net/http"
    "time"

    "daqnext/meson-cloud-client/daemon"
    "daqnext/meson-cloud-client/logger"
)

type apiMgr struct {
    url        string
    token      string
    httpCli    http.Client
    peersQueue chan string
    ipfsDaemon *daemon.IpfsDaemon
}

func NewApiMgr(url, token string, ipfsDaemon *daemon.IpfsDaemon) *apiMgr {
    return &apiMgr{
        url:        url,
        token:      token,
        httpCli:    http.Client{Timeout: time.Duration(5 * time.Second)},
        peersQueue: make(chan string, 10*1024),
        ipfsDaemon: ipfsDaemon,
    }
}

func (a *apiMgr) Run() {
    go a.tickerRun(time.NewTicker(30*time.Second), a.queryPeers)
    a.tickerRun(time.NewTicker(30*time.Second), a.updatePeersConfig)
}

func (a *apiMgr) queryPeers() {
    params := "token=" + a.token
    url := fmt.Sprintf("%s/cloud?%s", a.url, params)

    var peerList []string
    if err := reqGET(a.httpCli, url, &peerList); err != nil {
        logger.L.Error(err)
        return
    }

    for _, peer := range peerList {
        logger.L.Debug("Add peer into queue ", "peer=", peer)
        a.peersQueue <- peer
    }
    logger.L.Info("Finished querying peers")
}

func (a *apiMgr) updatePeersConfig() {
    peers := make(map[string]bool)
    flag := false
    outBreak := false
    for {
        logger.L.Debugw("Loop", "flag", flag, " outbreak", outBreak)
        if outBreak {
            break
        }
        select {
        case peer := <-a.peersQueue:
            flag = true
            peers[peer] = true
        default:
            outBreak = true
            break
        }
    }
    if flag {
        logger.L.Debug("Ready to call updating peer")
        if err := a.ipfsDaemon.UpdateConfig(peers); err != nil {
            logger.L.Errorw("Failed to update IPFS config", "err", err)
        }
    }
}

func (a *apiMgr) tickerRun(t *time.Ticker, f func()) {
    for {
        select {
        case <-t.C:
            f()
        }
    }
}
