package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
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
	fmt.Println("Правила в Базе знаний:")
	for k, v := range rules {
		fmt.Printf("%s -> %v\n", k, v)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("Введите факты:")
		text, _, _ := reader.ReadLine()
		fmt.Println()
		wq := proccessForward(text, rules)
		fmt.Println()
		fmt.Println("Итоговая рабочая память:")
		fmt.Println(wq[len(wq)-1])

	}
}

// Прямой ход для набора фактов
func proccessForward(baseRulesJSON []byte, rules map[string][]string) []string {
	workRules := make(map[string][]string, len(rules))
	for k, v := range rules {
		workRules[k] = append([]string{}, v...)
	}
	var workQueue []string

	// Кладем исходные факты в рабочую очередь
	var baseRules []string
	json.Unmarshal(baseRulesJSON, &baseRules)
	for _, rule := range baseRules {
		workQueue = addToWorkQueue(workQueue, rule)
	}

	fmt.Println()
	l := len(workQueue)
	fmt.Println("Выводы:")
	for i := 0; i < l; i++ {
		out, ok := workRules[workQueue[i]]
		if !ok {
			continue
		}
		for _, fact := range out {
			fmt.Println(workQueue[i], "->", fact)
			workQueue = addToWorkQueue(workQueue, fact)
		}

		l = len(workQueue)
	}

	return workQueue
}

func addToWorkQueue(workQueue []string, rule string) []string {
	l := len(workQueue)
	inputRule, _ := json.Marshal([]string{rule})
	// Добавляем в очередь факт
	workQueue = append(workQueue, string(inputRule))
	// Добавляем в очередь комбинацию факта с уже имеющимеся
	for i := 0; i < l; i++ {
		var tmp []string
		json.Unmarshal([]byte(workQueue[i]), &tmp)
		tmp = append(tmp, rule)
		sort.Strings(tmp)
		tmpNew, _ := json.Marshal(tmp)
		workQueue = append(workQueue, string(tmpNew))
	}
	return workQueue
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
	var rules []fileRule
	if err := json.NewDecoder(file).Decode(&rules); err != nil {
		return nil, fmt.Errorf("parsing rules file: %v", err)
	}
	out := make(map[string][]string, len(rules))
	for _, rule := range rules {
		sort.Strings(rule.Fact)
		data, _ := json.Marshal(rule.Fact)
		out[string(data)] = append(out[string(data)], rule.Goal)
	}
	return out, nil
}
