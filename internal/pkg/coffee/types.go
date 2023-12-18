package coffee

type MapItem struct {
	ID        int
	Name      string
	Type      ItemType
	ImageURL  string
	ReviewURL string
	Location
}

type ItemType string

const (
	CoffeeShopItem ItemType = "CoffeeShop"
	RoasterItem    ItemType = "Roaster"
	RestaurantItem ItemType = "Restaurant"
)

type Location struct {
	Address   string
	Latitude  float32
	Longitude float32
}
