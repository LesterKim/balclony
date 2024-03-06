package models

type Estate struct {
	MunicipalityID       int            `json:"municipality_id"`
	Owner                string         `json:"owner"`
	Address              string         `json:"address"`
	Block                string         `json:"block"`
	Lot                  string         `json:"lot"`
	EstateType           string         `json:"estate_type"`
	Hash                 string         `json:"hash"`
	RecentActivity       string         `json:"recent_activity"`
	Photos               []string       `json:"photos"`
	Slug                 string         `json:"slug"`
	MakePaymentURL       string         `json:"make_payment_url"`
	SmartContractAddress string         `json:"smart_contract_address"`
	Qualifier            string         `json:"qualifier"`
	IsVerified           bool           `json:"is_verified"`
	IsFeature            bool           `json:"is_feature"`
	Vacant               bool           `json:"vacant"`
	SaleDate             string         `json:"sale_date"`
	Lat                  StringOrNumber `json:"lat"`
	Lng                  StringOrNumber `json:"lng"`
}

type StringOrNumber struct {
	Value string
}

func (son *StringOrNumber) UnmarshalJSON(data []byte) error {
	if string(data) == "0" {
		son.Value = "0"
		return nil
	}

	// Remove quotes if the data is a string.
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}

	son.Value = string(data)
	return nil
}
