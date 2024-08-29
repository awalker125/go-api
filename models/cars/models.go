package cars

type Cars struct {
	Model string `json:"model"`
	Parts []Part `json:"parts"`
}

// Represents individual parts
type Part struct {
	Name string `json:"name"`
}
