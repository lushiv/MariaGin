package restaurants

// Restaurant represents a restaurant entity.
type TblRestaurant struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float64 `json:"rating"`
}
