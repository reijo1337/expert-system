package main

const (
	// FACT -
	FACT = iota
	// RULE -
	RULE
	// QUERY -
	QUERY
)

// Clause - запись в базе знаний
type Clause struct {
	head       *CompoundTerm
	goals      []*Goal
	variables  []*Variable
	clauseType int //fact, rule, or query?
}

func newClause(head *CompoundTerm, body []*Goal) *Clause {
	return &Clause{
		head:  head,
		goals: body,
	}
}
