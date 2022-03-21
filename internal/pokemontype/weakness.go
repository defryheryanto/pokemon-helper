package pokemontype

var weakToElements = map[IType][]IType{
	FireType: {
		WaterType,
		GroundType,
		RockType,
	},
	WaterType: {
		ElectricType,
		GrassType,
	},
	ElectricType: {
		GroundType,
	},
	GrassType: {
		FireType,
		IceType,
		PoisonType,
		FlyingType,
		BugType,
	},
	IceType: {
		FireType,
		FightingType,
		RockType,
		SteelType,
	},
	FightingType: {
		FlyingType,
		PsychicType,
		FairyType,
	},
	PoisonType: {
		GroundType,
		PsychicType,
	},
	GroundType: {
		WaterType,
		GrassType,
		IceType,
	},
	FlyingType: {
		ElectricType,
		IceType,
		RockType,
	},
	PsychicType: {
		BugType,
		GhostType,
		DarkType,
	},
	BugType: {
		FireType,
		FlyingType,
		RockType,
	},
	RockType: {
		WaterType,
		GrassType,
		FightingType,
		GroundType,
		SteelType,
	},
	GhostType: {
		GhostType,
		DarkType,
	},
	DragonType: {
		IceType,
		DragonType,
		FairyType,
	},
	DarkType: {
		FightingType,
		BugType,
		FairyType,
	},
	SteelType: {
		FireType,
		FightingType,
		GroundType,
	},
	FairyType: {
		PoisonType,
		SteelType,
	},
}

func getWeakness(element IType) []IType {
	return weakToElements[element]
}
