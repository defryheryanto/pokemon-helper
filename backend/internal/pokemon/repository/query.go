package repository

const (
	queryUpsertPokemon = `
		INSERT INTO pokemons (id, name, base_status, types, sprites)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name,
			base_status = EXCLUDED.base_status,
			types = EXCLUDED.types,
			sprites = EXCLUDED.sprites,
			updated_at = NOW()
	`
	queryGetPokemonByID = `
		SELECT id, name, base_status, types, sprites
		FROM pokemons
		WHERE id = $1
	`
	queryGetPokemonByName = `
		SELECT id, name, base_status, types, sprites
		FROM pokemons
		WHERE LOWER(name) = LOWER($1)
	`
	queryListPokemons = `
		SELECT id, name, base_status, types, sprites
		FROM pokemons
	`
)
