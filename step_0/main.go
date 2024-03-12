package main

import (
	"step-0/gate"
	"step-0/helpers"
)

func InitConf() {
	helpers.ReadConfig()
}

func main() {
	InitConf()

	// pretty print
	//docs.DrawStart()

	gate.InitChoice()
}
