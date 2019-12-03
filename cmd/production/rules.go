package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// загрузка правил из файла
func loadRules(filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open rules file: %v", err)
	}
	var rules []Rule
	if err := json.NewDecoder(file).Decode(&rules); err != nil {
		return nil, fmt.Errorf("parsing rules file: %v", err)
	}
	out := make(map[string][]string, len(rules))
	for _, rule := range rules {
		out[rule.Fact] = append(out[rule.Fact], rule.Goal)
	}
	return out, nil
}
