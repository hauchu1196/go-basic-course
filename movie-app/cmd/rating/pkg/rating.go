package models

type Rating struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	MovieID string `json:"movie_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
