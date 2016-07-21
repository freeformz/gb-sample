package main

import (
	"fmt"
	"mypackage"

	"github.com/heroku/slog"
)

func main() {
	fmt.Println(mypackage.Hello())
	fmt.Println(slog.Context{"hello": "there"})
}
