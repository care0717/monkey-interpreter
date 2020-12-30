package object

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, val Object) Object
}

func NewEnclosedEnvironment(outer Environment) Environment {
	env := newEnvironment()
	env.Outer = outer
	return env
}

func NewEnvironment() Environment {
	return newEnvironment()
}

func newEnvironment() *environment {
	s := make(map[string]Object)
	return &environment{Store: s}
}

type environment struct {
	Store map[string]Object
	Outer Environment
}

func (e *environment) Get(name string) (Object, bool) {
	obj, ok := e.Store[name]
	if !ok && e.Outer != nil {
		obj, ok = e.Outer.Get(name)
	}
	return obj, ok
}

func (e *environment) Set(name string, val Object) Object {
	e.Store[name] = val
	return val
}
