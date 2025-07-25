// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/twpayne/go-geom"
)

type CoffeeMapItem struct {
	ID        uuid.UUID
	ItemName  string
	ItemType  pgtype.Text
	ImageUrl  pgtype.Text
	ReviewUrl pgtype.Text
	Summary   pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Location  geom.Point
}
