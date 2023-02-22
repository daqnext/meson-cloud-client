package api

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
)

type httpWrap struct {
    Err    uint8       `json:"error""`
    ErrMsg string      `json:"errMsg"`
    Msg    interface{} `json:"msg,omitempty"`
}

func reqGET(c http.Client, url string, resp interface{}) error {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }

    req.Header.Add("Content-Type", "application/json")

    r, err := c.Do(req)
    if err != nil {
        return err
    }
    defer r.Body.Close()
    if r.StatusCode != http.StatusOK {
        return errors.New(fmt.Sprintf("HTTP Status code error: %v", r.Status))
    }

    tmp := httpWrap{Msg: resp}
    if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
        return err
    }
    if tmp.Err == 1 {
        return errors.New(tmp.ErrMsg)
    }
    return nil
}
