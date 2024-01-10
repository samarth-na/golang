package resource

import "fmt"

type MenuItem struct {
	ID       int
	Title    string
	Category string
	Price    float64
	Img      string
	Desc     string
}

func main() {
	menu := []MenuItem{
		{
			ID:       1,
			Title:    "buttermilk pancakes",
			Category: "breakfast",
			Price:    15.99,
			Img:      "./images/item-1.jpeg",
			Desc:     "I'm baby woke mlkshk wolf bitters live-edge blue bottle, hammock freegan copper mug whatever cold-pressed",
		},
		{
			ID:       2,
			Title:    "diner double",
			Category: "lunch",
			Price:    13.99,
			Img:      "./images/item-2.jpeg",
			Desc:     "vaporware iPhone mumblecore selvage raw denim slow-carb leggings gochujang helvetica man braid jianbing. Marfa thundercats",
		},
		{
			ID:       3,
			Title:    "godzilla milkshake",
			Category: "shakes",
			Price:    6.99,
			Img:      "./images/item-3.jpeg",
			Desc:     "ombucha chillwave fanny pack 3 wolf moon street art photo booth before they sold out organic viral.",
		},
		// ... (repeat the structure for the remaining items)
	}

	// Print the menu items
	for _, item := range menu {
		fmt.Printf("ID: %d\nTitle: %s\nCategory: %s\nPrice: $%.2f\nImg: %s\nDesc: %s\n\n",
			item.ID, item.Title, item.Category, item.Price, item.Img, item.Desc)
	}
}
