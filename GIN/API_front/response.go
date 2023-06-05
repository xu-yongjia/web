package api1

import "gintest/DBstruct"

// DataList 带有总数的Data结构
// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 用于产生返回出错信息
func ERRRESPONSE(reason string, errcode int) Response {
	return Response{
		Status: errcode,
		Msg:    reason,
	}
}

// 用于返回成功信息以及附带的数据
func SUCCESSRESPONSE(data interface{}) Response {
	return Response{
		Msg:    "OK",
		Status: 200,
		Data:   data,
	}
}

func SUCCESSRESPONSE_NODATA() Response {
	return Response{
		Msg:    "OK",
		Status: 200,
	}
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Msg:    "ok",
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}

// Product 商品序列化器
type Product_json struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CanteenID     int    `json:"canteen_id"`
	CategoryID    int    `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	AvgScore      string `json:"avg_score"`
	CreatedAt     int64  `json:"created_at"`
}

func BuildProduct(item DBstruct.Product) Product_json {
	return Product_json{
		ID:            item.ID,
		Name:          item.Name,
		CanteenID:     item.CanteenID,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		AvgScore:      item.Score,
		CreatedAt:     item.CreatedAt.Unix(),
	}
}

// BuildProducts 序列化商品列表
func BuildProducts(items []DBstruct.Product) (products []Product_json) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
