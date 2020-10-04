package controllers

// Context はcontrollerで使うメソッドのinterface
type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
}