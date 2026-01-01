package products

import "backend-web-commision-kana/internal/utils/jsonresp"

type ReqAdminProducts struct {
	ProductName  string `json:"product_name"`
	ProductPrice int32  `json:"product_price"`
}

type ResAdminProduct struct {
	Msg string `json:"msg"`
}

type Res struct {
	JSONres jsonresp.JSONResponder
}
