package pokemontype

type IType interface {
	WeakAgainst() []*Type
	StrongAgainst() []*Type
}

type Type string

func (t Type) WeakAgainst() []*Type {
	return nil
}

func (t Type) StrongAgainst() []*Type {
	return nil
}
