package blabberBot

import "time"

//Broadcast ...
type Broadcast struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	CreateDate time.Time `json:"createDate"`
	ExpiryDate time.Time `json:"expiryDate"`
}
