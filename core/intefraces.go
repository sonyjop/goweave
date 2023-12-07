package core

type CompnentManager interface {
	//CreateNode(component Component) Node
	GetComponent(componentName string) (Component, error)
	CreateEndPoint(uri string) (Component, error)
}

type EndPointManager interface {
	GetRouteContext() (RouteContext, error)
	SetRouteContext(context RouteContext)
}
