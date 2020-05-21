# wakeonlan

Very small library to:
1. create a Wake-on-LAN magic packet
2. send it to the broacast address `255.255.255.255:9`

## Usage

```golang
import "github.com/da-rod/wakeonlan"

randomMAC := "a7:cc:00:0c:eb:40"
if mp, err := wakeonlan.NewMagicPacket(randomMAC); err == nil {
    mp.Send()
}
```
See [example](cmd/wol/main.go) that builds the CLI binary.
