package main

import (
	"flag"
	"log"

	"github.com/da-rod/wakeonlan"
)

func main() {
	dst := flag.String("mac", "", "MAC Address to wake up")
	flag.Parse()
	mp, err := wakeonlan.NewMagicPacket(*dst)
	if err != nil {
		log.Fatal(err)
	}
	if err = mp.Send(); err != nil {
		log.Fatal(err)
	}
}
