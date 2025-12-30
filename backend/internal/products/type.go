package products

import "backend-web-commision-kana/internal/utils/jsonresp"

type ReqAdminProducts struct {
	ProductName  string `json:"product_name"`
	ProductPrice int32  `json:"product_price"`
}

type ResProduct struct {
	ID           int32  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductPrice int32  `json:"product_price"`
}

type ResProducts struct {
	Products []ResProduct `json:"products"`
}

type Res struct {
	JSONres jsonresp.JSONResponder
}
