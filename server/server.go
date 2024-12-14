package server

type Server interface {
	Start()
	Use(interface{})
}
