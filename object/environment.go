package object

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, val Object) Object
}

func NewEnvironment() Environment {
	s := make(map[string]Object)
	return &environment{store: s}
}

type environment struct {
	store map[string]Object
}

func (e *environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
