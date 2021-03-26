package models

type ReviewPhrase struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Phrase    string `json:"phrase" bson:"phrase"`
	Frequency int    `json:"frequency" bson:"frequency"`
}
