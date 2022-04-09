package vermouth

import (
	"flag"
	"fmt"
	"runtime"
)

var GitCommit string
var Version string
var BuildDate string

func init() {
	usage := flag.Usage
	flag.Usage = func() {
		fmt.Println("App Version:", Version)
		fmt.Println("Code Version:", GitCommit)
		fmt.Println("Build Date:", BuildDate)
		fmt.Println("Go Version:", runtime.Version())
		fmt.Println("GOOS:", runtime.GOOS)
		usage()
	}
}
