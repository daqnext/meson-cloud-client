package daemon

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpfsDaemon_ReadConfig(t *testing.T) {
	t.Skip("Skipped")
	cfg := IpfsCfg{
		IpfsCmd:      "../ipfs",
		IpfsDataRoot: "~/.ipfs",
	}
	iDaemon := NewIpfsDaemon(&cfg)

	resp, err := iDaemon.ReadConfig()
	assert.NoError(t, err)
	tp, err := json.Marshal(resp)
	fmt.Println(string(tp))
}
