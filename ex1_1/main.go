package main

import (
	"fmt"
	"os"

	"github.com/elusivenode/tgopl/echo"
)

func main() {
	s := echo.Echo(os.Args[:])
	fmt.Println(s)
}
