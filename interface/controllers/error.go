package controllers

// Error はMessageを持つ構造体
type Error struct {
	Message string
}

// NewError はerrの中身を持つErrorのアドレスを返す関数
func NewError(err error) *Error {
	return &Error {
		Message: err.Error(),
	}
}