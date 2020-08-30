package helper

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	outfile, _ = os.Create(wd + "server.log")
	WriteLog   = log.New(outfile, time.Now().String()+" > ", 0)
)

func Println(s ...interface{}) {
	WriteLog.Println(s...)
	fmt.Println(s...)
}
