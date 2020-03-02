package config

import (
	"fmt"
	"sort"
)

type Config struct {
	Bind   map[string]*BindEntry
	Groups map[string][]string
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

func (c *Config) BindKeys() []string {
	keys := make([]string, 0, len(c.Bind))
	for key := range c.Bind {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
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

func (c *Config) GroupedBindEntries(names []string) ([]*BindEntry, error) {
	bs := make([]*BindEntry, 0, len(names))
	for _, name := range names {
		g, ok := c.Groups[name]
		if !ok {
			return nil, &configReadError{fmt.Sprintf("No group entry: %s", name)}
		}
		b, err := c.BindEntries(g)
		if err != nil {
			return nil, err
		}
		bs = append(bs, b...)
	}
	return bs, nil
}

func (c *Config) GroupKeys() []string {
	keys := make([]string, 0, len(c.Groups))
	for key := range c.Groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
