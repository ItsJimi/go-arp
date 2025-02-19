# go-arp
[![go reference](https://pkg.go.dev/badge/github.com/ItsJimi/go-arp)](https://pkg.go.dev/github.com/ItsJimi/go-arp)

Parse ARP file from `/proc/net/arp`

## Install
```shell
go get -u github.com/ItsJimi/go-arp
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

#### GetEntryFromMAC
```go
package main

import (
  "fmt"
  "github.com/ItsJimi/go-arp"
)

func main() {
  entry, err := arp.GetEntryFromMAC("00:00:00:00:00:00")
  if err != nil {
    fmt.Println(err)
    return
   }

  fmt.Println(entry.IPAddress)
}
```

## Contribute
Feel free to fork and make pull requests

## License
[MIT](https://github.com/ItsJimi/go-arp/blob/master/LICENSE)
