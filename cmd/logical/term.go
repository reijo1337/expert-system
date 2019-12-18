package main

type (
	// Term терм у пролога
	Term interface {
		unify(other Term) bool
		String() string
	}

	// TrueTerm -
	TrueTerm struct {
	}

	// Atom -
	Atom struct {
		name string
	}
)

// TrueTerm реализация интерфейса
func (t *TrueTerm) unify(other Term) bool {
	return true
}

func (t *TrueTerm) String() string {
	return "TRUE"
}

// Atom реализация интерфейса
func (a *Atom) unify(other Term) bool {
	return a.name == other.String()
}

func (a *Atom) String() string {
	return a.name
}
