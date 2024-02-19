package coffee

type MapItem struct {
	ID         int      `db:"id"`
	Name       string   `db:"name"`
	Type       ItemType `db:"type"`
	ImageURL   string   `db:"image_url" json:"image_url"`
	ReviewURL  string   `db:"review_url" json:"review_url"`
	LocationID int      `db:"location_id" json:"-"`
	Location
}

type ItemType string

const (
	CoffeeShopItem ItemType = "CoffeeShop"
	RoasterItem    ItemType = "Roaster"
	RestaurantItem ItemType = "Restaurant"
)

type Location struct {
	Address   string  `db:"address"`
	Latitude  float32 `db:"latitude"`
	Longitude float32 `db:"longitude"`
}
