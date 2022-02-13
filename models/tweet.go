package models

/* Tweet caputra el mensaje del tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
