package coffee

import (
	"context"
	"fmt"

	"github.com/maisieccino/maisie-site/internal/pkg/db"
)

type DBStore struct {
	*db.Queries
}

func (s *DBStore) Get(ctx context.Context, id string) (MapItem, error) {
	item, err := s.GetItem(ctx, id)
	if err != nil {
		return MapItem{}, fmt.Errorf("reading from DB: %w", err)
	}

	return MapItem{
		ID:         item.ID,
		Name:       item.Name,
		Type:       ItemType(item.ItemType.String),
		ImageURL:   item.ImageUrl.String,
		ReviewURL:  item.ReviewUrl.String,
		LocationID: 0,
		Location:   Location{},
	}, err
}
