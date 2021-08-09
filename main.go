package main

import (
	"pcqq/core"
)

func main() {
	var pc = core.PCQQ{}
	pc.Init()
	pc.GetQrCode()
	pc.ListenMsg()


}
