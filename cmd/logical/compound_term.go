package main

// CompoundTerm -
type CompoundTerm struct {
	atom *Atom
	args []Term
}

func newCompoundTerm(atom *Atom, args []Term) *CompoundTerm {
	return &CompoundTerm{
		atom: atom,
		args: args,
	}
}

func (c *CompoundTerm) unify(other Term) bool {
	ct, ok := other.(*CompoundTerm)
	if !ok {
		return false
	}
	if c.atom.String() != ct.atom.String() {
		return false
	}

	if len(c.args) != len(ct.args) {
		return false
	}

	for i, arg := range c.args {
		if arg.unify(ct.args[i]) {
			return false
		}
	}

	return true
}

func (c *CompoundTerm) String() string {
	ret := c.atom.String() + "("
	for i, arg := range c.args {
		ret += arg.String()
		if i < len(c.args)-1 {
			ret += ", "
		}
	}
	return ret + ")"
}
