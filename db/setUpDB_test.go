package db_test

import (
	"flashCardProject/db"
	"testing"
)

func TestSetUpDatabase(t *testing.T) {
	_, err := db.SetUpDatabase()
	if err != nil {
		t.Errorf("SetUpDatabase() error =%v", err)
	}
}
