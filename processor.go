package main

type processor interface {
	url() (string)
	process() ([]processor)
}
