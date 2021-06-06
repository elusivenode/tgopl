package main

import (
	"fmt"
	"os"

	"github.com/elusivenode/tgopl/echo"
)

func main() {
	s := echo.EchoIdxArgs(os.Args[:])
	fmt.Println(s)
}