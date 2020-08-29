package helper

import (
	"log"
	"os"
)

var (
	outfile, _ = os.Create(wd + "server.log")
	WriteLog   = log.New(outfile, "", 0)
)

// func init() {
// 	// set location of log file
// 	var logpath = wd + "log/server.log"

// 	flag.Parse()
// 	var file, err1 = os.Create(logpath)

// 	if err1 != nil {
// 		panic(err1)
// 	}
// 	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
// 	Log.Println("LogFile : " + logpath)
// }
