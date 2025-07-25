package models

type UserSW struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
	Position string `json:"position"`
	Password string `json:"password"`
}

type ShopSW struct {
	Id          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	CityId      int    `json:"city_id"`
	About       string `json:"about"`
	Email       string `json:"email"`
	TIN         string `json:"TIN"`
	CompanyName string `json:"company_name"`
	CompanyLogo string `json:"company_logo"`
	LegalEntity string `json:"legal_entity"`
}

type ReportUpdate struct {
	Id             int    `json:"id"`
	Image          string `json:"image"`
	Description    string `json:"description"`
	Crops          int    `json:"crops"`
	PlantDiseases  int    `json:"plant_diseases"`
	DisDescription string `json:"dis_description"`
}

type Migration struct {
	Id        int    `json:"id"`
	Migration string `json:"migration"`
	Batch     int    `json:"batch"`
}
