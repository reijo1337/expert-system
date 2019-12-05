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
	rules, backRules, err := loadRules(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Правила в Базе знаний:")
	for k, v := range rules {
		fmt.Printf("%s -> %v\n", k, v)
	}

	fmt.Println()
	fmt.Println("БЗ в обратном порядке")
	for k, v := range backRules {
		fmt.Println(k, "<-", v)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Введите факты:")
	text, _, _ := reader.ReadLine()
	fmt.Println()
	wq := proccessForward(text, rules)
	if len(wq) != 0 {
		fmt.Println()
		fmt.Println("Итоговая рабочая память:")
		fmt.Println(wq[len(wq)-1])
	}

	fmt.Println()
	fmt.Println("Введите факт для проверки:")
	text, _, _ = reader.ReadLine()
	fmt.Println()
	proccessBack(string(text), backRules)
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

func proccessBack(goal string, backRules map[string][]string) {
	if _, ok := backRules[goal]; !ok {
		fmt.Println("В БЗ нет правила для вывода факта", goal)
		return
	}

	workRules := make(map[string][]string, len(backRules))
	for k, v := range backRules {
		workRules[k] = append([]string{}, v...)
	}
	finalFacts := []string{}
	stack := []string{goal}

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rightFacts := workRules[top]
		delete(workRules, top)

		for _, rightFact := range rightFacts {
			var tmp []string
			json.Unmarshal([]byte(rightFact), &tmp)
			for _, newGoal := range tmp {
				if _, ok := workRules[newGoal]; !ok && !strings.Contains(strings.Join(finalFacts, " "), newGoal) {
					finalFacts = append(finalFacts, newGoal)
					continue
				}
				stack = append(stack, newGoal)
			}
			fmt.Println(top, "<-", rightFact)
		}
	}
	if len(finalFacts) != 0 {
		fmt.Println()
		fmt.Println("Итоговый вывод:")
		for _, fact := range finalFacts {
			fmt.Println(fact)
		}
	}
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

// загрузка правил из файла
func loadRules(filePath string) (map[string][]string, map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("open rules file: %v", err)
	}
	var rules []fileRule
	if err := json.NewDecoder(file).Decode(&rules); err != nil {
		return nil, nil, fmt.Errorf("parsing rules file: %v", err)
	}
	out := make(map[string][]string, len(rules))
	back := make(map[string][]string, len(rules))
	for _, rule := range rules {
		sort.Strings(rule.Fact)
		data, _ := json.Marshal(rule.Fact)
		out[string(data)] = append(out[string(data)], rule.Goal)
		back[rule.Goal] = append(back[rule.Goal], string(data))
	}
	return out, back, nil
}
