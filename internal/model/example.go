package model

//go:generate klctl gen model2 Example --database DB
// @def primary ID
// @def unique_index I_name_id Name
// @def index I_ff_user UserID FF
type Example struct {
	PrimaryID
	Name    string `db:"F_name,default='',size=100" json:"name"` // Name
	CarType int    `db:"F_car_type,default=0" json:"CarType"`
	FF      string `db:"f_FF,size=100"`
	RefUser
}
