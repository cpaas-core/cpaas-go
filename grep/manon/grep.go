package grep

import (
	"fmt"
	"os/exec"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	args := append(flags, pattern)
	var result []string
	args = append(args, files...)
	cmd := exec.Command("grep", args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	strOutput := strings.TrimSpace(string(stdout))

	if strOutput != "" {
		result = append(result, strings.Split(strOutput, "\n")...)
	}
	return result
}
