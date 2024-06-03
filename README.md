# FlashCardProject

FlashCardProject is a simple Go application designed to manage flashcards. It provides functionalities to set up a database and display flashcards in a tabular format.


## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/mhoarb/FlashCardProject.git
    cd FlashCardProject
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

## Usage
 ```bash
go buld
go install
```

### Setting Up the Database

The `SetUpDatabase` function in the `db` package sets up a MySQL database connection.

```go
package main

import (
    "fmt"
    "FlashCardProject/db"
)

func main() {
    db, err := db.SetUpDatabase()
    if err != nil {
        fmt.Println("Error setting up database:", err)
        return
    }
    fmt.Println("Database setup successfully:", db)
}
```
## Displaying Flashcards

### The DisplayTable function in the db package displays flashcards in a tabular format using the tablewriter package.
```go
package main

import (
    "FlashCardProject/db"
)

func main() {
    flashCards := []db.FlashCard{
        {Question: "What is Go?", Answer: "A programming language"},
        {Question: "Who created Go?", Answer: "Google"},
    }
    db.DisplayTable(flashCards)
}
```
## Project Structure
```bash
FlashCardProject/
├── cmd/
├── db/
│   ├── db.go
│   ├── db_test.go
│   ├── flashcard.go
│   └── flashcard_test.go
├── internal/
├── go.mod
├── go.sum
└── main.go
```
 cmd/: Contains command-line interface related files.
 db/: Contains database-related code and tests.
 internal/: Contains internal application logic.
 main.go: Entry point of the application.



