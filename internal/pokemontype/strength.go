package pokemontype

var effectiveElements = map[IType][]IType{
	FireType: {
		GrassType,
		IceType,
		BugType,
		SteelType,
	},
	WaterType: {
		FireType,
		GroundType,
		RockType,
	},
	ElectricType: {
		WaterType,
		FlyingType,
	},
	GrassType: {
		WaterType,
		GroundType,
		RockType,
	},
	IceType: {
		GrassType,
		GroundType,
		FlyingType,
		DragonType,
	},
	FightingType: {
		NormalType,
		IceType,
		RockType,
		DarkType,
		SteelType,
	},
	PoisonType: {
		GrassType,
		FairyType,
	},
	GroundType: {
		FireType,
		ElectricType,
		PoisonType,
		RockType,
		SteelType,
	},
	FlyingType: {
		GrassType,
		FightingType,
		BugType,
	},
	PsychicType: {
		FightingType,
		PoisonType,
	},
	BugType: {
		GrassType,
		PsychicType,
	},
	RockType: {
		FireType,
		IceType,
		FlyingType,
		BugType,
	},
	GhostType: {
		PsychicType,
		GhostType,
	},
	DragonType: {
		DragonType,
	},
	DarkType: {
		PsychicType,
		GhostType,
	},
	SteelType: {
		IceType,
		RockType,
		FairyType,
	},
	FairyType: {
		FightingType,
		DragonType,
		DarkType,
	},
}

func GetEffective(elements IType) []IType {
	return effectiveElements[elements]
}
