# go-arp
[![godoc](https://godoc.org/github.com/ItsJimi/go-arp?status.svg)](https://godoc.org/github.com/ItsJimi/go-arp)

Parse ARP file from `/proc/net/arp`

## Install
```shell
go get -u github.com/itsjimi/go-arp
```

## Example
#### GetEntries
```go
package main

import (
  "fmt"
  "github.com/ItsJimi/go-arp"
)

func main() {
  list, err := arp.GetEntries()
  if err != nil {
    fmt.Println(err)
    return
   }

  fmt.Println(list[0].IPAddress)
}
```

## Contribute
Feel free to fork and make pull requests

## License
[MIT](https://github.com/ItsJimi/go-arp/blob/master/LICENSE)
