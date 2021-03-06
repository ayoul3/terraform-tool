package main

import (
	"flag"
	"os"

	"github.com/ayoul3/terraform-tool/lib"
	log "github.com/sirupsen/logrus"
)

var folderName, tag string

func init() {
	flag.StringVar(&folderName, "path", ".", "subfolder to check for diff")
	flag.StringVar(&tag, "tag", "origin/master", "Tag or branch name to diff against")
	flag.Parse()
}
func main() {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		log.Fatalf("Cannot find folder %s", folderName)
	}
	log.Infof("Checking diff on folder %s and tag: %s", folderName, tag)
	lib.PrintComponents(folderName, tag)
}
