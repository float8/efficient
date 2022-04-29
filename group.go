package efficient

type Group struct {
}

type GroupInterface interface {
	Execute(cxt Context)
}

func (g *Group) Execute(cxt Context) {}
