package rest

type AvailableNumber struct {
	// these properties should always come
	Name    string `json:"friendly_name"`
	Number  string `json:"phone_number"`
	Country string `json:"iso_country"`

	// only available for us & canada
	Lata       string `json:"lata"`
	RateCenter string `json:"rate_center"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
}

type AvailableNumbersResponse struct {
	Uri              string            `json:"uri"`
	AvailableNumbers []AvailableNumber `json:"available_phone_numbers"`
}
