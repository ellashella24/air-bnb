package homestay

type FormReqUpdate struct {
	Name  string
	Price int
}

type FormReqCreate struct {
	Name   string
	Price  int
	CityId uint
}
