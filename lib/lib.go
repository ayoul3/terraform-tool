package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/prometheus/common/log"
)

const tfExtension = ".txt"

func PrintComponents(path, tag string) error {
	listComponent := make(map[string]bool, 0)

	fileList, err := getGitMasterDiff(tag)

	if err != nil {
		return err
	}
	for _, file := range fileList {
		dir := filepath.Dir(file)
		if dir == path {
			continue
		}
		if strings.Contains(dir, "module") {
			for _, c := range GetAffectedComponents(path, dir) {
				listComponent[c] = true
			}
			continue
		}
		if !ContainsTFFiles(dir) {
			continue
		}
		listComponent[dir] = true
	}
	log.Infof("Found %d components", len(listComponent))
	for k, _ := range listComponent {
		fmt.Println(k)
	}

	return nil
}

func GetAffectedComponents(path, modulePath string) []string {
	var affectedComponents []string
	moduleName := getModuleName(modulePath)
	modulePattern := fmt.Sprintf(`source\s+=\s+"[:_.\w/-]+/%s`, moduleName)
	modulePatternRe := regexp.MustCompile(modulePattern)

	filepath.Walk(path, func(path string, f os.FileInfo, _ error) error {
		if f.IsDir() || strings.HasPrefix(path, ".") || !strings.HasSuffix(f.Name(), tfExtension) {
			return nil
		}
		data, _ := ioutil.ReadFile(path)
		if modulePatternRe.Match(data) {
			affectedComponents = append(affectedComponents, filepath.Dir(path))
		}
		return nil
	})
	return affectedComponents
}

func getModuleName(path string) string {
	listPaths := strings.Split(path, "/")
	return listPaths[len(listPaths)-1]
}

func getGitMasterDiff(tag string) ([]string, error) {
	cmd := exec.Command("git", "--no-pager", "diff", "HEAD", tag, "--name-only")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	listFolders := strings.Split(string(out), "\n")
	return listFolders, nil
}

func ContainsTFFiles(path string) bool {
	var found bool
	filepath.Walk(path, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), tfExtension) {
			found = true
			return nil
		}
		return nil
	})
	return found
}
