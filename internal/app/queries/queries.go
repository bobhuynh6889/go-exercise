package queries

import (
	"fmt"

	db "github.com/bobhuynh6889/go-exercise/internal/app/databases"
	"github.com/bobhuynh6889/go-exercise/internal/app/models"
	"github.com/rs/zerolog/log"
)

// type UserQueries struct {
// 	*sql.DB
// }

type UserQueries struct {
	DB db.Database
}

// get all users
func (q *UserQueries) GetUsers() ([]models.User, error) {
	query := "SELECT * FROM users"

	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// declare user Model
	users := []models.User{}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.PassWord, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Error().Err(err).Msg(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return users, nil
}

// get user by Id
func (q *UserQueries) GetUserById(id int) ([]models.User, error) {
	users := []models.User{}
	query := "SELECT * FROM users WHERE id = $1"

	rows, err := q.DB.Query(query, id)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.PassWord, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Error().Err(err).Msg(err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return users, nil
}

// get user by Name
func (q *UserQueries) GetUserByName(UserName string) ([]models.User, error) {
	users := []models.User{}
	query := "SELECT * FROM users WHERE user_name = $1"

	fmt.Print(UserName)

	rows, err := q.DB.Query(query, UserName)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.PassWord, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Error().Err(err).Msg(err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return users, nil
}

// create an user
func (q *UserQueries) CreateUser(ctx *models.User) error {
	query := `
			INSERT INTO users (id, user_name, pass_word, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
		`
	_, err := q.DB.Exec(query, ctx.ID, ctx.UserName, ctx.PassWord, ctx.CreatedAt, ctx.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return err
}

// delete an user
func (q *UserQueries) DeleteUser(id int) error {
	query := `
        DELETE FROM users
        WHERE id = $1
    `

	_, err := q.DB.Exec(query, id)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return err
}

// update an user
func (q *UserQueries) UpdateUser(id int, ctx *models.User) ([]models.User, error) {
	users := []models.User{}
	query := `
        UPDATE users
        SET user_name = $2, pass_word = $3, updated_at = $4
        WHERE id = $1
    `
	_, err := q.DB.Exec(query, id, ctx.UserName, ctx.PassWord, ctx.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	users, err = q.GetUserById(id)

	return users, nil
}
