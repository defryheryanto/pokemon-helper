package pokemontype

type IType interface {
	WeakAgainst() []IType
	StrongAgainst() []IType
}

type Type string

func (t Type) WeakAgainst() []IType {
	return getWeakness(t)
}

func (t Type) StrongAgainst() []IType {
	return getEffective(t)
}
