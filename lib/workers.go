package lib

import (
	"fmt"
	"time"

	"github.com/prometheus/common/log"
)

var Chan chan string
var Output chan []string

func PrepareWorkers(folderName string) {
	for {
		path := <-Chan
		Output <- LookupComponents(folderName, path)
	}
}

func Report() {
	listComponent := make(map[string]bool, 0)
	var paths []string
	var ok bool
	for {
		select {
		case paths = <-Output:
			ok = true
		case <-time.After(50 * time.Millisecond):
			ok = false
		}
		if !ok {
			break
		}
		for _, p := range paths {
			listComponent[p] = true
		}
	}
	log.Infof("Found %d components", len(listComponent))
	for k, _ := range listComponent {
		fmt.Println(k)
	}
}
