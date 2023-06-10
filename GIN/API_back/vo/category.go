package vo

type CategoryList struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	CanteenId    uint   `json:"canteen_id"`
}
