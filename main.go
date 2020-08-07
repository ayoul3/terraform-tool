package main

import (
	"flag"
	"os"

	"github.com/ayoul3/terraform-tool/lib"
	"github.com/prometheus/common/log"
)

var folderName, tag string

const threadNum = 5

func init() {
	flag.StringVar(&folderName, "path", ".", "subfolder to check for diff")
	flag.StringVar(&tag, "tag", "origin/master", "Tag or branch name to diff against")
	flag.Parse()
}
func main() {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		log.Fatalf("Cannot find folder %s", folderName)
	}
	lib.Chan = make(chan string, threadNum)
	for i := 0; i < threadNum; i++ {
		go lib.PrepareWorkers(folderName)
	}

	log.Infof("Checking diff on folder %s and tag: %s", folderName, tag)
	lib.StartWork(folderName, tag)
}
