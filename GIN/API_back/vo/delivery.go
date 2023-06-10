package vo

type DeliveryList struct {
	DeliverId     uint   `json:"deliver_id"`
	Truename      string `json:"truename"`
	Phone         string `json:"phone"`
	CanteenNumber int    `json:"canteen_number"`
}
