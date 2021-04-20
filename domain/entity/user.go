package entity

type User struct {
	UID         string `json:"uid"`
	DisplayName string `json:"display_name"`
	AccessToken string `json:"access_token"`
}
