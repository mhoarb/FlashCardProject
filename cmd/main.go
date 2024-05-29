package main

import (
	"flashCardProject/internal"
)

func main() {
	internal.NewFlashCard("what is your name", "mmd")
	internal.NewFlashCard("why is not be a human", "because the its hot")
	internal.DisplayTable()
}
