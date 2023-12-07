package core

type Component struct {
	ComponentName string
	Path          string
	Query         map[string]string
}
type Node struct {
	ComponentName string
	Path          string
	Query         map[string]string
}
type message struct {
	Header     map[string]string
	Body       interface{}
	Properties map[string]string
}

type Exchange struct {
	Properties map[string]string
	StepTrace  []message
	In         message
	Out        message
}
type RouteContext struct {
	
}
