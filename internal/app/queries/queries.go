// package queries

// // BookQueries struct for queries from Book model.
// type UserQueries struct {
// 	*sqlx.DB
// }

// GetBooks method for getting all books.
// func (q *UserQueries) GetBooks() ([]models.Book, error) {
// 	// Define books variable.
// 	books := []models.Book{}

// 	// Define query string.
// 	query := `SELECT * FROM books`

// 	// Send query to database.
// 	err := q.Select(&books, query)
// 	if err != nil {
// 		// Return empty object and error.
// 		return books, err
// 	}

// 	// Return query result.
// 	return books, nil
// }
