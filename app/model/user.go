package model

type UserModel struct {
	UserId    string `json:"UserId"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	Status    string `json:"Status"`
	CreatedAt int64  `json:"CreatedAt"`
	UpdatedAt int64  `json:"UpdatedAt"`
}
