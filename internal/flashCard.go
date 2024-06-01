package internal

type FlashCard struct {
	ID       uint `gorm:"primaryKey"`
	Question string
	Answer   string
	Status   string
}
