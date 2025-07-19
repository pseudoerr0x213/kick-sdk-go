package kicksdk

// const baseURL := "https://api.kick.com/public/v1/categories"

type GetCategoriesResponse struct {
	Data    []Category `json:"data"`
	Message string     `json:"message"`
}
