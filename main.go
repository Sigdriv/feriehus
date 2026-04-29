package main

import "github.com/Sigdriv/feriehus/handler"

func main() {
	srv := handler.CreateHandler()

	srv.CreateGinGroup()
}
