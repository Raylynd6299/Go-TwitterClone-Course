package models

// Tweet este es la estructura del tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
