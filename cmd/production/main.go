package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	backward = flag.Bool("backward", false, "Прямой ход")
	fileName = flag.String("f", "examples/production.json", "Файл с правилами")
)

func main() {
	rules, err := loadRules(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rules:")
	for k, v := range rules {
		fmt.Printf("%s: %v\n", k, v)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input first rule:")
		text, _, _ := reader.ReadLine()
		fmt.Println(proccessRuleForward(string(text), rules))
	}
}

func proccessRuleForward(rule string, rules map[string][]string) string {
	workRules := make(map[string][]string, len(rules))
	for k, v := range rules {
		workRules[k] = append([]string{}, v...)
	}
	workBuf := rule

mainLoop:
	for {
		for fact, goals := range workRules {
			if contains(fact, workBuf) {
				fmt.Println(fact, "->", goals)
				for _, goal := range goals {
					workBuf += goal
					delete(workRules, fact)
					continue mainLoop
				}
			}
		}
		return workBuf
	}
}

// находится ли факт в рабочей памяти
func contains(sub, check string) bool {
	for _, r := range sub {
		if !strings.ContainsRune(check, r) {
			return false
		}
	}
	return true
}

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
