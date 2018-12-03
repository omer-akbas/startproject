package main

import (
	"github.com/zserge/webview"
)

func main() {
	webview.Open("Minimal webview example",
		"https://scan.digitastic.de/login", 800, 600, true)
}
