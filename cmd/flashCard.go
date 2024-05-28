package main

type IFlashCard interface {
	CreateFlashCard() error
	ListFlashCard() error
	StartPlay()
	Statistics()
	Restart()
	Exit()
}
