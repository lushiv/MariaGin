package restaurants

type ErrorResponse struct {
	Message string `json:"message"`
}

type CommonResponse struct {
	Message string `json:"message"`
}

type AddRestaurantRequest struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float64 `json:"rating"`
}

type UpdateRestaurantRequest struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float64 `json:"rating"`
}

type DeleteRestaurantRequest struct {
	ID int `json:"id"`
}

type GetRestaurantsResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float64 `json:"rating"`
}

