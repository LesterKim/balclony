package models

type Listing struct {
	UID                  string    `json:"uid"`
	IsVerified           bool      `json:"is_verified"`
	Photos               []string  `json:"photos"`
	EstateType           string    `json:"estate_type"`
	Hash                 string    `json:"hash"`
	Owner                string    `json:"owner"`
	Address              string    `json:"address"`
	Block                string    `json:"block"`
	Lot                  string    `json:"lot"`
	LastTaxPaymentAmount string    `json:"last_tax_payment_amount"`
	LastTaxPaymentDate   string    `json:"last_tax_payment_date"`
	Lat                  string    `json:"lat"`
	Lng                  string    `json:"lng"`
	HouseClass           string    `json:"house_class"`
	YearBuilt            string    `json:"year_built"`
	SaleInformation      []Sale    `json:"sale_information"`
	IsOwner              bool      `json:"is_owner"`
	Comments             []string  `json:"comments"`
	History              []TaxList `json:"history"`
	// ... include other fields as necessary
}

type Sale struct {
	ID               int     `json:"id"`
	EstateUID        string  `json:"estate_uid"`
	Date             string  `json:"date"`
	County           string  `json:"county"`
	District         string  `json:"district"`
	Book             string  `json:"book"`
	DeedNote         *string `json:"deed_note"`
	RTFee            *string `json:"r_t_fee"`
	Grantor          *string `json:"grantor"`
	RTFExempt        *string `json:"r_t_f_exempt"`
	Page             string  `json:"page"`
	DeedRecorder     *string `json:"deed_recorder"`
	Price            string  `json:"price"`
	NU               string  `json:"nu"`
	Grantee          *string `json:"grantee"`
	YearSameAsDeed   *string `json:"year_same_as_deed"`
	Buildings        *string `json:"buildings"`
	PropertyLocation *string `json:"property_location"`
	YearBuilt        *string `json:"year_built"`
	Remarks          *string `json:"remarks"`
	Land             *string `json:"land"`
	Total            *string `json:"total"`
	FloorArea        *string `json:"floor_area"`
	Ratio            string  `json:"ratio"`
	Block            *string `json:"block"`
	Lot              *string `json:"lot"`
	Qual             *string `json:"qual"`
	Class            *string `json:"class"`
	CL4Type          *string `json:"cl_4_type"`
	Condo            *string `json:"condo"`
	NonusableCode    *string `json:"nonusable_code"`
	SerialNumber     *string `json:"serial_Number"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type TaxList struct {
	EstateUID      string  `json:"estate_uid"`
	Year           string  `json:"year"`
	Owner          string  `json:"owner"`
	Street         string  `json:"street"`
	CityStateZip   string  `json:"city_state_zip"`
	LandAssmnt     string  `json:"land_assmnt"`
	BuildingAssmnt string  `json:"building_assmnt"`
	Exempt         string  `json:"exempt"`
	TotalAssmnt    string  `json:"total_assmnt"`
	Assessed       string  `json:"assessed"`
	Imp            *string `json:"imp"`
	PropClass      *string `json:"prop_class"`
	AdditionalLots *string `json:"additional_lots"`
	Deductions     *string `json:"deductions"`
}
