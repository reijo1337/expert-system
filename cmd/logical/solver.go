package main

type solver struct {
	rules []interface{}
}

func newSolver(rulesText string) *solver {
	return &solver{}
}
