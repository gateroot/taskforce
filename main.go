package main

import "github.com/hiroaki-sekine/taskforce/di"

func main() {
	cmd := di.InjectCommandLineDelivery()
	cmd.Exec()
}
