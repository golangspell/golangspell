package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// GetModuleName from the current application based in the go.mod file
func GetModuleName(currentPath string) string {
	filePath := fmt.Sprintf("%s%sgo.mod", currentPath, PlatformSeparator)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return ""
	}
	contentText := string(content)
	re := regexp.MustCompile("module (.*?)\n")
	match := re.FindStringSubmatch(contentText)
	if len(match) >= 2 {
		return strings.Trim(match[1], " ")
	}
	return ""
}
