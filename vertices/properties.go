package vertices

type Properties interface {
	Property(name string) interface{}
	DeleteProperty(name string)
	PropertiesCount() int
	SetProperty(name string, property interface{})
}
