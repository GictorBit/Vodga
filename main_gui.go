package main

import (
	"github.com/TheWeirdDev/Vodga/ui"
	"github.com/gotk3/gotk3/gtk"
)


func main() {
	gtk.Init(nil)

	ui.StartGui()

	gtk.Main()
}

