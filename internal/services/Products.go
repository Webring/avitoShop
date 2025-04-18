package services

import "errors"

func ProductPrice(item string) (uint, error) {
	prices := map[string]uint{
		"t-shirt":    80,
		"cup":        20,
		"book":       50,
		"pen":        10,
		"powerbank":  200,
		"hoody":      300,
		"umbrella":   200,
		"socks":      10,
		"wallet":     50,
		"pink-hoody": 500,
	}

	price, exists := prices[item]
	if !exists {
		return 0, errors.New("product not found")
	}
	return price, nil
}
