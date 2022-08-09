package model

//go:generate klctl gen model2 User --database DB
// @def primary ID
// @def unique_index I_user_id UserID
type User struct {
	PrimaryID
	RefUser
}

type RefUser struct {
	UserID string `db:"F_user_id,size=100"`
}
