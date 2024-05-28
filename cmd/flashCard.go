package main

type IFlashCard interface {
	CreateFlashCard() error
	ListFlashCard() error
	StartPlay()
	Statistics()
	Restart()
	Exit()
}

type FlashCard struct {
	Question string
	Answer   string
}
