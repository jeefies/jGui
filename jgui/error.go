package jgui

type JguiError struct {
	msg string
}

func (je *JguiError) Error() string {
	return je.msg
}

func NewError(msg string) (*JguiError) {
	return &JguiError{msg}
}
