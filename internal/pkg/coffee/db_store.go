package coffee

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maisieccino/maisie-site/internal/pkg/db"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

type DBStore struct {
	*db.Queries
}

func NewDBStore(conn *pgxpool.Pool) Store {
	queries := db.New(conn)
	return &DBStore{queries}
}

func toMapItem(i db.CoffeeMapItem) MapItem {
	var loc Location
	if !i.Location.Empty() {
		loc.Latitude = i.Location.X()
		loc.Longitude = i.Location.Y()
	}

	return MapItem{
		ID:        i.ID.String(),
		Name:      i.ItemName,
		Type:      ItemType(i.ItemType.String),
		ImageURL:  i.ImageUrl.String,
		ReviewURL: i.ReviewUrl.String,
		Location:  loc,
	}
}

func (s *DBStore) Get(ctx context.Context, id string) (MapItem, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return MapItem{}, fmt.Errorf("parsing UUID: %w", err)
	}
	item, err := s.GetItem(ctx, u)
	if err != nil {
		return MapItem{}, fmt.Errorf("reading from DB: %w", err)
	}

	return toMapItem(item), err
}

func (s *DBStore) Create(context.Context, MapItem) error {
	return nil
}

func (s *DBStore) Update(context.Context, string, MapItem) error {
	return nil
}

func (s *DBStore) List(ctx context.Context) ([]MapItem, error) {
	items, err := s.ListItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("reading from DB: %w", err)
	}
	result := []MapItem{}
	for _, i := range items {
		result = append(result, toMapItem(i))
	}
	return result, nil
}

func (s *DBStore) SearchByArea(ctx context.Context, params SearchByAreaParams) ([]MapItem, error) {
	coords := []geom.Coord{
		{
			params.X0,
			params.Y0,
		},

		{
			params.X1,
			params.Y0,
		},

		{
			params.X1,
			params.Y1,
		},

		{
			params.X0,
			params.Y1,
		},

		{
			params.X0,
			params.Y0,
		},
	}
	p := geom.NewPolygon(geom.XY)
	p.MustSetCoords([][]geom.Coord{coords})
	p.SetSRID(4326)

	enc, err := ewkb.Marshal(p, ewkb.NDR)
	if err != nil {
		return nil, fmt.Errorf("encoding input: %w", err)
	}
	results, err := s.Queries.SearchByArea(ctx, enc)
	if err != nil {
		return nil, fmt.Errorf("reading from DB: %w", err)
	}

	items := []MapItem{}
	for _, i := range results {
		items = append(items, toMapItem(i))
	}
	return items, nil
}

func (s *DBStore) Delete(context.Context, string) error {
	return nil
}
