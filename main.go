package main

import "taskforce/di"

func main() {
	cmd := di.InjectCommandLineDelivery()
	cmd.Exec()
}
