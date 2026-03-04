package loader

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
}

var (
	productCategories = []string{
		"Electronics", "Clothing", "Books", "Home & Garden", "Sports",
		"Toys", "Beauty", "Automotive", "Food", "Health",
	}

	productAdjectives = []string{
		"Premium", "Professional", "Ultra", "Compact", "Deluxe",
		"Wireless", "Smart", "Portable", "Vintage", "Modern",
		"Classic", "Heavy-Duty", "Lightweight", "Eco-Friendly", "Budget",
	}

	productNouns = map[string][]string{
		"Electronics":   {"Headphones", "Speaker", "Charger", "Cable", "Monitor", "Keyboard", "Mouse", "Webcam", "Tablet", "Laptop"},
		"Clothing":      {"T-Shirt", "Jeans", "Jacket", "Dress", "Shoes", "Hat", "Socks", "Sweater", "Shorts", "Coat"},
		"Books":         {"Novel", "Textbook", "Cookbook", "Biography", "Guide", "Manual", "Journal", "Comic", "Magazine", "Anthology"},
		"Home & Garden": {"Lamp", "Chair", "Table", "Vase", "Planter", "Mirror", "Rug", "Curtains", "Cushion", "Clock"},
		"Sports":        {"Ball", "Racket", "Weights", "Mat", "Bike", "Helmet", "Gloves", "Shoes", "Bag", " Bottle"},
		"Toys":          {"Puzzle", "Game", "Doll", "Car", "Robot", "Blocks", "Plush", "Kit", "Model", "Set"},
		"Beauty":        {"Cream", "Serum", "Lotion", "Shampoo", "Perfume", "Makeup", "Brush", "Mask", "Oil", "Spray"},
		"Automotive":    {"Filter", "Wiper", "Cover", "Mat", "Cleaner", "Tool", "Light", "Charger", "Mount", "Holder"},
		"Food":          {"Snacks", "Drinks", "Sauce", "Spice", "Coffee", "Tea", "Honey", "Chocolate", "Cereal", "Nuts"},
		"Health":        {"Vitamins", "Supplement", "Bandage", "Thermometer", "Mask", "Cream", "Spray", "Gel", "Pills", "Tea"},
	}

	brands = []string{
		"TechPro", "HomeEssentials", "Active", "StyleMaxGear", "NaturePure",
		"UrbanStyle", "ClassicLine", "BudgetBuy", "PremiumPlus", "SmartChoice",
		"EliteBrand", "BasicFit", "ProGrade", "EcoLife", "FutureTech",
	}
)

func GenerateProducts(count int) []Product {
	rand.Seed(time.Now().UnixNano())

	products := make([]Product, count)

	for i := 0; i < count; i++ {
		category := productCategories[rand.Intn(len(productCategories))]
		adjective := productAdjectives[rand.Intn(len(productAdjectives))]
		noun := productNouns[category][rand.Intn(len(productNouns[category]))]
		brand := brands[rand.Intn(len(brands))]

		products[i] = Product{
			ID:          int64(i + 1),
			Name:        fmt.Sprintf("%s %s %s", brand, adjective, noun),
			Description: generateProductDescription(category, adjective, noun),
			Category:    category,
			Brand:       brand,
			Price:       generatePrice(),
		}
	}

	return products
}

func generateProductDescription(category, adjective, noun string) string {
	templates := []string{
		"The %s %s is perfect for everyday use. High quality and reliable performance.",
		"Upgrade your %s experience with this %s %s. Designed for durability and ease of use.",
		"Looking for a %s %s? This premium %s offers exceptional value.",
		"Discover the %s %s - your ideal companion for all %s needs.",
		"This %s %s combines style and functionality at an affordable price.",
	}

	template := templates[rand.Intn(len(templates))]
	return fmt.Sprintf(template, adjective, noun, category)
}

func generatePrice() float64 {
	base := 9.99 + rand.Float64()*990.01
	return float64(int(base*100)) / 100
}
