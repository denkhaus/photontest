package main

import (
	"github.com/denkhaus/photontest/components"
	"github.com/denkhaus/vecty"
)

func main() {
	vecty.SetTitle("photontest")
	vecty.RenderBody(components.NewIndex("http://localhost:8000"))
}
