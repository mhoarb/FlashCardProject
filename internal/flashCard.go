package internal

// FlashCard represents a flash card entity with ID, Question, Answer, and Status fields.
// ID is the primary key field for database operations, annotated with `gorm:"primaryKey"`.
type FlashCard struct {
	ID       uint `gorm:"primaryKey"`
	Question string
	Answer   string
	Status   string
}
