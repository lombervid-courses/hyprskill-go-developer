package main

import (
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	//github.com/fatih/color
	//go get github.com/fatih/color@v1.12.0
	programmer := emoji.Sprint(":man_technologist:")
	emoji.Printf(":woman_technologist:")
	color.Green(" Learning about modules! " + programmer)

	//github.com/kyokomi/emoji/v2
	//programmer := emoji.Sprint(":man_technologist:")
	//emoji.Printf(":woman_technologist:")
	//emoji.Println(" Learning about modules! " + programmer)
}
