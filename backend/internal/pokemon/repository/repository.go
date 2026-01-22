package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Upsert(ctx context.Context, data *pokemon.PokemonData) error {
	entity, err := newPokemonEntityFromModel(data)
	if err != nil {
		return err
	}
	if entity == nil {
		return nil
	}

	var baseStatus interface{}
	if len(entity.BaseStatus) > 0 {
		baseStatus = entity.BaseStatus
	}

	var types interface{}
	if len(entity.Types) > 0 {
		types = entity.Types
	}

	_, err = r.db.ExecContext(
		ctx,
		queryUpsertPokemon,
		entity.ID,
		entity.Name,
		baseStatus,
		types,
		entity.Sprites,
	)
	return err
}

func (r *Repository) Get(ctx context.Context, filter pokemon.GetFilter) (*pokemon.PokemonData, error) {
	var row *sql.Row
	if filter.ID != 0 {
		row = r.db.QueryRowContext(ctx, queryGetPokemonByID, filter.ID)
	} else if filter.Name != "" {
		row = r.db.QueryRowContext(ctx, queryGetPokemonByName, filter.Name)
	} else {
		return nil, fmt.Errorf("pokemon repository: empty get filter")
	}

	entity, err := scanPokemonRow(row)
	if err != nil {
		return nil, err
	}
	return entity.ToModel()
}

func (r *Repository) List(ctx context.Context, filter pokemon.ListFilter) ([]pokemon.PokemonData, error) {
	query, args := buildListQuery(filter)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []pokemon.PokemonData{}
	for rows.Next() {
		entity, err := scanPokemonRows(rows)
		if err != nil {
			return nil, err
		}
		model, err := entity.ToModel()
		if err != nil {
			return nil, err
		}
		if model != nil {
			results = append(results, *model)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func scanPokemonRow(row *sql.Row) (*PokemonEntity, error) {
	var entity PokemonEntity
	var baseStatus []byte
	var types []byte
	var sprites sql.NullString
	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&baseStatus,
		&types,
		&sprites,
	)
	if err != nil {
		return nil, err
	}

	entity.BaseStatus = baseStatus
	entity.Types = types
	if sprites.Valid {
		entity.Sprites = sprites.String
	}
	return &entity, nil
}

func scanPokemonRows(rows *sql.Rows) (*PokemonEntity, error) {
	var entity PokemonEntity
	var baseStatus []byte
	var types []byte
	var sprites sql.NullString
	err := rows.Scan(
		&entity.ID,
		&entity.Name,
		&baseStatus,
		&types,
		&sprites,
	)
	if err != nil {
		return nil, err
	}

	entity.BaseStatus = baseStatus
	entity.Types = types
	if sprites.Valid {
		entity.Sprites = sprites.String
	}
	return &entity, nil
}

func buildListQuery(filter pokemon.ListFilter) (string, []interface{}) {
	query := strings.Builder{}
	query.WriteString(queryListPokemons)

	args := []interface{}{}
	argIndex := 1

	whereClauses := []string{}
	search := strings.TrimSpace(filter.Search)
	if search != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("name ILIKE $%d", argIndex))
		args = append(args, fmt.Sprintf("%%%s%%", search))
		argIndex++
	}

	elementType := strings.TrimSpace(filter.ElementType)
	if elementType != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("types ? $%d", argIndex))
		args = append(args, strings.ToLower(elementType))
		argIndex++
	}

	if len(whereClauses) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(whereClauses, " AND "))
	}

	query.WriteString(" ORDER BY id")

	if filter.Limit > 0 {
		query.WriteString(fmt.Sprintf(" LIMIT $%d", argIndex))
		args = append(args, filter.Limit)
		argIndex++
	}
	if filter.Offset > 0 {
		query.WriteString(fmt.Sprintf(" OFFSET $%d", argIndex))
		args = append(args, filter.Offset)
	}

	return query.String(), args
}
