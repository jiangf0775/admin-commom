package types

type User struct {
	Id     uint64   `json:"id"`
	Name   string   `json:"name"`
	RoleId []uint64 `json:"roleId"`
}
