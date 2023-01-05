package contorller

type Route struct {
	Method      string
	Path        string
	HandlerFunc func(req interface{}) (res interface{}, err error)
}

type Struct struct {
	Routes []Route
}
