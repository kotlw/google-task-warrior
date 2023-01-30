package config

import "fmt"

// Default app config.
func Default() (*Config, error) {

	c := &Config{
		Taskwarrior: Taskwarrior{
			Taskrc: Taskrc{
				Path: "~/.taskrc",
			},
			TaskData: TaskData{
				Path: "~/.task",
			},
		},
	}

	if err := postprocess(c); err != nil {
		return nil, fmt.Errorf("postprocess: %w", err)
	}

	return c, nil
}
