package main

import (
	"engine/engine"
	"engine/webhook"
)

func main() {
	engine.InitRules()
	webhook.HookStart()
}
