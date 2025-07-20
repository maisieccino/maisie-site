package coffee

type MapItem struct {
	ID        string   `db:"id"`
	Name      string   `db:"name"`
	Type      ItemType `db:"type"`
	ImageURL  string   `db:"image_url" json:"image_url"`
	ReviewURL string   `db:"review_url" json:"review_url"`
	Location
}

type ItemType string

const (
	CoffeeShopItem ItemType = "CoffeeShop"
	RoasterItem    ItemType = "Roaster"
	RestaurantItem ItemType = "Restaurant"
)

type Location struct {
	// TODO: Set address later.
	// Address   string  `db:"address"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
}
