package main

import "regexp"

var (
	tokenRegexp   = regexp.MustCompile(`[A-Za-z0-9_]+|:\-|[()\.,]`)
	atomNameRegex = regexp.MustCompile(`^[A-Za-z0-9_]+$`)
	variableRegex = regexp.MustCompile(`^[A-Z_][A-Za-z0-9_]*$`)
)

func getTokenList(inputText string) []string {
	return tokenRegexp.FindAllString(inputText, -1)
}

type parser struct {
	tokens    []string
	current   string
	iterIndex int
	scope     map[string]rulePart
	finished  bool
}

func newParser(inputText string) *parser {
	tokens := getTokenList(inputText)
	p := &parser{
		tokens:    tokens,
		iterIndex: -1,
		finished:  false,
	}
	p.parseNext()
	return p
}

func (p *parser) parseNext() {
	p.iterIndex++
	p.current = p.tokens[p.iterIndex]
	p.finished = p.iterIndex == len(p.tokens)-1
}

func (p *parser) parseRules() {
	for !p.finished {
		p.scope = make(map[string]rulePart)

	}
}

func (p *parser) parseRule() rulePart {
	head := p.parseTerm()
	if p.current == "." {
		p.parseNext()
		return &rule{head: head, tail: true()}
	}

	if p.current != ":-" {
		panic("Expected :- in rule but got " + p.current)
	}
	p.parseNext()

	arguments := make([]rulePart, 0)
	for p.current != "." {
		arguments = append(arguments, p.parseTerm())
		if p.current != "," && p.current != "." {
			panic("Expected , or . in term but got " + p.current)
		}
		if p.current == "," {
			p.parseNext()
		}
	}
	p.parseNext()

	if len(arguments) > 1 {
		return &conjunction{arguments: arguments}
	}

	return &rule{head: head, tail: arguments[0]}
}

// Парсим терм
func (p *parser) parseTerm() rulePart {
	//  Нарвались на конъюкцию
	if p.current == "(" {
		p.parseNext()
		arguments := make([]rulePart, 0)
		for p.current != ")" {
			arguments = append(arguments, p.parseTerm())
			if p.current != "," && p.current != ")" {
				panic("Expected , or ) in term but got " + p.current)
			}
			if p.current == "," {
				p.parseNext()
			}
		}
		p.parseNext()
		return &conjunction{arguments: arguments}
	}

	functor := p.parseAtom()
	// Нарвались на переменную
	if variableRegex.MatchString(functor) {
		if functor == "_" {
			return &variable{name: functor}
		}
		variab, ok := p.scope[functor]
		if !ok {
			p.scope[functor] = &variable{name: functor}
			variab = p.scope[functor]
		}
		return variab
	}

	// Если нет аргументов в обработке - вернем атом, который как терм, без аргументов
	if p.current != "(" {
		return &term{functor: functor}
	}
	p.parseNext()

	// Парсим аргументы терма
	arguments := make([]rulePart, 0)
	for p.current != ")" {
		arguments = append(arguments, p.parseTerm())
		if p.current != "," && p.current != ")" {
			panic("Expected , or ) in term but got " + p.current)
		}
		if p.current == "," {
			p.parseNext()
		}
	}
	p.parseNext()

	return &term{functor: functor, arguments: arguments}
}

func (p *parser) parseAtom() string {
	name := p.current
	if !atomNameRegex.MatchString(name) {
		panic("Invalid Atom Name: " + name)
	}
	p.parseNext()
	return name
}
