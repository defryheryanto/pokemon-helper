package pokemontype

type IType interface {
	WeakAgainst() []IType
	StrongAgainst() []IType
}

type Type string

func (t Type) WeakAgainst() []IType {
	return GetWeakness(t)
}

func (t Type) StrongAgainst() []IType {
	return GetEffective(t)
}
