# ymo

[![GoDoc](https://godoc.org/github.com/ad/ymo?status.svg)](https://godoc.org/github.com/ad/ymo)
[![Go Report Card](https://goreportcard.com/badge/github.com/ad/ymo)](https://goreportcard.com/report/github.com/ad/ymo)
[![Build Status](https://travis-ci.org/ad/ymo.svg?branch=master)](https://travis-ci.org/ad/ymo)
[![Coverage Status](https://coveralls.io/repos/github/ad/ymo/badge.svg?branch=master)](https://coveralls.io/github/ad/ymo?branch=master)

ymo is a Go package for sending offline conversions to Yandex.Metrica. It provides a convenient way to upload offline conversion data to Yandex.Metrica's API.

For detailed documentation and API reference, please visit the [official Yandex.Metrica documentation](https://yandex.ru/dev/metrika/doc/api2/management/offline_conversion/upload.html).

## Installation

To use ymo in your Go project, you need to import it:


### Usage

```go
package main

import (
    "fmt"
    "github.com/ad/ymo"
)

func main() {
    client := ymo.NewYMOClient(1234567, "your ym token", "CLIENT_ID", true)
    err := client.SendEvent(
        ymo.Event{
            ClientId: "your client id",
            Target: "GOAD_ID",
            DateTime: strconv.FormatInt(time.Now().Unix(), 10),
            Price: fmt.Sprintf("%.2f", 999.99),
            Currency: "RUB",
        },
    )
    if err != nil {
        fmt.Println(err)
    }
}
```
