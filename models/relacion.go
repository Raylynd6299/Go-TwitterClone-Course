package models

// Relacion : es el modelo para grabar la relacion entre dos usuarios
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRealcionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
