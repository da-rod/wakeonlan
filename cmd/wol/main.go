package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/da-rod/wakeonlan"
)

func main() {
	dst := flag.String("mac", "", "MAC Address to wake up")
	flag.Parse()
	mp, err := wakeonlan.NewMagicPacket(*dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = mp.Send(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	os.Exit(0)
}
