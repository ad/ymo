# ymo

[![GoDoc](https://godoc.org/github.com/ad/ymo?status.svg)](https://godoc.org/github.com/ad/ymo)
[![Test](https://github.com/ad/ymo/actions/workflows/gotest.yml/badge.svg)](https://github.com/ad/ymo/actions/workflows/gotest.yml)

ymo is a Go package for sending offline conversions to Yandex.Metrika. It provides a convenient way to upload offline conversion data to Yandex.Metrika's API.

For detailed documentation and API reference, please visit the [official Yandex.Metrika documentation](https://yandex.ru/dev/metrika/doc/api2/management/offline_conversion/upload.html).

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
    client := ymo.NewYMOClient("1234567", "your ym token", "CLIENT_ID", true)
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
