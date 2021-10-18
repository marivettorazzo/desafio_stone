package domain

import (
	"time"
)

type Transfer struct {
	ID                     int       `json:"id,omitempty"`
	Amount                 float64   `json:"Amount,omitempty"`
	Account_origin_id      int       `json:"Account_origin_id,omitempty"`
	Account_destination_id int       `json:"Account_destination_id,omitempty"`
	Created_at_            time.Time `json:"Created_a,omitempty"`
}
