package daemon

import (
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"strings"
	"time"

	"daqnext/meson-cloud-client/logger"
	"daqnext/meson-cloud-client/portable"
)

type IpfsCfg struct {
    IpfsCmd      string
    IpfsDataRoot string
}

type IpfsDaemon struct {
    cfg           *IpfsCfg
    restartSignal chan bool
}

func NewIpfsDaemon(cfg *IpfsCfg) *IpfsDaemon {
    return &IpfsDaemon{
        cfg:           cfg,
        restartSignal: make(chan bool, 1),
    }
}

func (i *IpfsDaemon) Start(parentCtx context.Context) error {
    ctx, _ := context.WithCancel(parentCtx)

    // TODO: Check the binary for IPFS or download
    runNode, _ := portable.CmdGen(i.cfg.IpfsCmd, "daemon")
    if err := runNode.Start(); err != nil {
        return err
    }

    logger.L.Infow("IPFS starts", "PID", runNode.Process.Pid)

    go func() {
        er := make(chan error, 1)
        go func() { er <- runNode.Wait() }()

        for {
            select {
            case <-i.restartSignal:
                logger.L.Infow("Restart IPFS daemon")
                portable.CmdKill(runNode)
                for j := 0; j < 3; j++ {
                    time.Sleep(5 * time.Second)
                    if err := i.Start(parentCtx); err != nil {
                        logger.L.Errorw("Failed to start IPFS daemon", "err", err)
                    } else {
                        break
                    }
                }
            case <-ctx.Done():
                logger.L.Infow("Receive Cancel Signal", "pid", runNode.Process.Pid, ctx.Err())
                portable.CmdKill(runNode)
                return
            case err := <-er:
                logger.L.Warn("Runtime error ", err.Error())
                return
            }
        }
    }()
    return nil
}

type pJson struct {
    Addrs []string `json:"Addrs"`
    ID    string   `json:"ID"`
}

type peersJson struct {
    Peers []pJson `json:"Peers"`
}

func (i *IpfsDaemon) ReadConfig() (*peersJson, error) {
    out, err := exec.Command(i.cfg.IpfsCmd, "config", "Peering").Output()
    if err != nil {
        return nil, err
    }
    // TODO
    var pl peersJson
    if err := json.Unmarshal(out, &pl); err != nil {
        return nil, err
    }
    return &pl, nil
}

func (i *IpfsDaemon) UpdateConfig(newPeers map[string]bool) error {
    existed, err := i.ReadConfig()
    if err != nil {
        return err
    }

    hit := make(map[string]bool)
    for _, p := range existed.Peers {
        hit[p.Addrs[0]+p.ID] = true
    }

    flag := false
    for p, _ := range newPeers {
        s1 := strings.Split(p, "/p2p/")
        if len(s1) != 2 {
            return errors.New("malformed" + p)
        }

        if hit[s1[0]+s1[1]] == true {
            continue
        } else {
            flag = true
            hit[s1[0]+s1[1]] = true
        }

        s2 := strings.Split(s1[0], "/tcp/")

        existed.Peers = append(existed.Peers, pJson{
            Addrs: []string{s2[0] + "/tcp/" + s2[1], s2[0] + "/udp/" + s2[1]},
            ID:    s1[1],
        })
    }

    if !flag {
        logger.L.Info("No need to update ipfs config")
        return nil
    }

    tp, err := json.Marshal(existed)
    if err != nil {
        return err
    }

    if err := exec.Command(i.cfg.IpfsCmd, "config", "--json", "Peering", string(tp)).Run(); err != nil {
        return err
    }
    logger.L.Info("Success to update peer info and send restart signal")
    i.restartSignal <- true
    return nil
}
