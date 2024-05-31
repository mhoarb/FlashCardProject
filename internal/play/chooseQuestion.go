package play

import (
	"bufio"
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
)

func ChooseQuestion() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}
	slog.Info("please write the question you want to train")
	reader := bufio.NewReader(os.Stdin)
	question, _ := reader.ReadString('\n')
	question = question[:len(question)-1]

	slog.Info("please write the Answer of this question")
	answer, _ := reader.ReadString('\n')
	answer = answer[:len(answer)-1]

	r := gin.Default()
	r.GET("/question", func(c *gin.Context) {
		var questions []internal.FlashCard
		db.Find(&questions)
		c.JSON(http.StatusOK, questions)
	})

}
