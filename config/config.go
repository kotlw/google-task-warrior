package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/kotlw/google-task-warrior/internal/taskwarrior"
)

type (
	// Config - main config structure.
	Config struct {
		Taskwarrior Taskwarrior
	}

	// Taskwarrior - parameters related to cli app.
	Taskwarrior struct {
		Taskrc   Taskrc
		TaskData TaskData
	}

	// Taskrc - parameters related to taskrc file.
	Taskrc struct {
		Path string
		File map[string]string
	}

	// TaskData - parameters related to task data folder.
	TaskData struct {
		Path string
	}
)

func overwriteStrIfEnv(targetValue *string, envKey string) {
	if envValue := os.Getenv(envKey); envValue != "" {
		*targetValue = envValue
	}
}

func postprocess(c *Config) (err error) {
	m := map[string]*string{
		"TASKRC":   &c.Taskwarrior.Taskrc.Path,
		"TASKDATA": &c.Taskwarrior.TaskData.Path,
	}

	for k, v := range m {
		overwriteStrIfEnv(v, k)
	}

	for _, v := range m {
		*v = strings.Replace(*v, "~", "$HOME", 1)
		*v = os.ExpandEnv(*v)
	}

	c.Taskwarrior.Taskrc.File, err = taskwarrior.ReadRC(c.Taskwarrior.Taskrc.Path)
	if err != nil {
		return fmt.Errorf("taskwarrior.ReadRC: %w", err)
	}
	if os.Getenv("TASKDATA") == "" && c.Taskwarrior.Taskrc.File["data.location"] != "" {
		c.Taskwarrior.TaskData.Path = c.Taskwarrior.Taskrc.File["data.location"]
	}

	return nil
}
