package models

/* Relacion tiene la relacion entre dos usuarios */
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
