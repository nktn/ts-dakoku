package main

import "github.com/nktn/ts-dakoku/app"

func main() {
	if _, err := app.Run(); err != nil {
		panic(err)
	}
}
