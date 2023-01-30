package taskwarrior

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadRC - read taskrc configuration file into map.
func ReadRC(path string) (map[string]string, error) {
	res := make(map[string]string)

	dat, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		s := strings.Split(line, "=")
		res[s[0]] = s[1]
	}

	return res, nil
}
