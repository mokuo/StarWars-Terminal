package main

type Terminal interface {
	Setup()
	Cmd() string
	Args() []string
}
