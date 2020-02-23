package config

import (
	"fmt"
)

type Config struct {
	Bind map[string]*BindEntry
}

type BindEntry struct {
	Src   string
	Links []string
}

type configReadError struct {
	message string
}

func (err *configReadError) Error() string {
	return err.message
}

func (c *Config) BindEntries(names []string) ([]*BindEntry, error) {
	bs := make([]*BindEntry, 0, len(names))
	for _, name := range names {
		b, ok := c.Bind[name]
		if !ok {
			return nil, &configReadError{fmt.Sprintf("No bind entry: %s", name)}
		}
		bs = append(bs, b)
	}
	return bs, nil
}
