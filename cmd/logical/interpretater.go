package main

type (
	rulePart interface {
		print() string
	}
	rule struct {
		head rulePart
		tail rulePart
	}
	conjunction struct {
		functor   string
		arguments []rulePart
	}
	variable struct {
		name string
	}
	term struct {
		functor   string
		arguments []rulePart
	}
)

func (r rule) print() string {
	return r.head.print() + " :- " + r.tail.print()
}

func (c *conjunction) print() string {
	ret := c.arguments[0].print()
	for i := 1; i < len(c.arguments); i++ {
		ret += ", " + c.arguments[i].print()
	}
	return ret
}

func (v *variable) print() string {
	return v.name
}

func (v *term) print() string {
	ret := v.functor
	if len(v.arguments) != 0 {
		ret += "(" + v.arguments[0].print()
		for i := 0; i < len(v.arguments); i++ {
			ret += ", " + v.arguments[i].print()
		}
		ret += ")"
	}
	return ret
}

func true() *term {
	return &term{functor: "TRUE"}
}
