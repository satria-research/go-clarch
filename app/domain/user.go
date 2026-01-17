package domain

type User struct {
	Id             string `json:"id" bson:"_id"`
	Username       string `json:"username" bson:"username"`
	Fullname       string `json:"fullname" bson:"fullname"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
	FavoritePhrase string `json:"favorite_phrase" bson:"favorite_phrase"`
}
