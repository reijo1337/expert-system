package main

// Variable -
type Variable struct {
	name     string
	instance Term
}

func (v *Variable) unify(other Term) bool {
	if v.instance != v {
		return v.instance.unify(other)
	}
	v.instance = other
	return true
}

func (v *Variable) String() string {
	ret := v.name
	if v.instance != v {
		ret += " -> " + v.instance.String()
	}
	return ret
}

func (v *Variable) reset() {
	v.instance = v
}
