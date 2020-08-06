package lib

import (
	"fmt"
	"os/exec"
	"strings"
)

func PrintComponents(path, tag string) error {
	fileList, err := getGitMasterDiff(tag)
	if err != nil {
		return err
	}
	fmt.Println(fileList)
	return nil
}

func getGitMasterDiff(tag string) ([]string, error) {
	diffOp := fmt.Sprintf("%s HEAD", tag)
	cmd := exec.Command("git", "--no-pager", "diff", diffOp, "--name-only")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	listFolders := strings.Split(string(out), "\n")
	return listFolders, nil
}
