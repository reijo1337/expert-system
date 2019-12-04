package main

// Rule описывает правила для продукционной модели
type Rule struct {
	Fact string
	Goal string
}

// Парсинг файла с правилами
type fileRule struct {
	Fact []string `json:"fact"`
	Goal string   `json:"goal"`
}
