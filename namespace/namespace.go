package namespace

type Namespace struct {
	prefix string
}

func NewNamespace(namespace string) *Namespace {
	return &Namespace{
		prefix: namespace,
	}
}

func (n *Namespace) Get(name ...string) string {
	if len(name) > 0 && name[0] != "" {
		return n.prefix + "-" + name[0]
	}

	return n.prefix
}
