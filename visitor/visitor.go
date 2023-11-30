package visitor

type Visitor interface {
	Target() string
	Visit() ([]string, error)
}
