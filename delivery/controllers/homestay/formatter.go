package homestay

type FormReqUpdate struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type FormReqCreate struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	CityId uint   `json:"cityid"`
}
