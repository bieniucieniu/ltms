package store

type Colection struct {
	namespaces map[string]*Namespace
}

func NewColection() *Colection {
	l := Colection{namespaces: map[string]*Namespace{}}
	return &l
}
