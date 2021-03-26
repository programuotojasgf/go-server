package models

type ReviewPhrase struct {
	ID        string `json:"_id"`
	Phrase    string `json:"phrase" bson:"phrase"`
	Frequency int    `json:"frequency" bson:"frequency"`
}
