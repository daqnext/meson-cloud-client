package api

import (
    "testing"
)

func TestApiMgr_QueryNode(t *testing.T) {
    t.Skip("Skipped")
    a := NewApiMgr("http://127.0.0.1:3535", "2", nil)

    a.queryPeers()
}
