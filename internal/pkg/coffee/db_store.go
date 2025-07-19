package coffee

import (
	"context"
	"database/sql"
	"errors"
)

type DBStore struct {
	sql.DB
}

func (s *DBStore) Get(ctx context.Context, id string) (MapItem, error) {
	if id == "" {
		return MapItem{}, errors.New("no id specified")
	}

	row := s.QueryRowContext(ctx, `SELECT * FROM coffee_map WHERE ID=$1`, id)
	if row == nil {
		return MapItem{}, errors.New("reading from database")
	}

	return MapItem{}, nil
}
