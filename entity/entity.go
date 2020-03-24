package entity

type Driver struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Drivers struct {
	Drivers []Driver
}
