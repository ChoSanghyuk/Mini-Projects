package handler

type AssetHistReq struct {
	ID uint `json:"id" validate:"required"`
}

type TotalStatusResp struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type AddFundReq struct {
	Name string `json:"name" validate:"required"`
}

type AddAssetReq struct {
	Name      string  `json:"name" validate:"required"`
	Category  uint    `json:"category" validate:"required"`
	Currency  string  `json:"currency" validate:"required"`
	Top       float64 `json:"top" validate:"required"`
	Bottom    float64 `json:"bottom" validate:"required"`
	SellPrice float64 `json:"sel_price"`
	BuyPrice  float64 `json:"buy_price"`
	Path      string  `json:"path"`
}

type UpdateAssetReq struct {
	ID        uint    `json:"id" validate:"required"`
	Name      string  `json:"name"`
	Category  uint    `json:"category"`
	Currency  string  `json:"currency"`
	Top       float64 `json:"top"`
	Bottom    float64 `json:"bottom"`
	SellPrice float64 `json:"sel_price"`
	BuyPrice  float64 `json:"buy_price"`
	Path      string  `json:"path"`
}

type DeleteAssetReq struct {
	ID uint `json:"id" validate:"required"`
}

type MarketStatusParam struct {
	Date string `json:"date" validate:"date"`
}

type SaveMarketStatusParam struct {
	Status uint `json:"status" validate:"required"`
}

type SaveInvestParam struct {
	FundId  uint    `json:"fund_id"`
	AssetId uint    `json:"asset_id"`
	Price   float64 `json:"price"`
	Count   int     `json:"count"`
}
