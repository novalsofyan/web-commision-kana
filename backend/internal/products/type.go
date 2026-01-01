package products

import "backend-web-commision-kana/internal/utils/jsonresp"

type ReqAdminProduct struct {
	ProductName  string `json:"product_name"`
	ProductPrice int32  `json:"product_price"`
}

type ResAdminProduct struct {
	Msg string `json:"msg"`
}

type ReqAdminProducts struct {
	Products []ReqAdminProduct `json:"products"`
}

type ResUserProduct struct {
	ID           int32  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductPrice int32  `json:"product_price"`
}

type ResUserProducts struct {
	Products []ResUserProduct `json:"products"`
}

type Res struct {
	JSONres jsonresp.JSONResponder
}
