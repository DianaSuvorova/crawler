package main

type processor interface {
	url() (string)
	process(availSpaceInQueue int) ([]processor)
}
