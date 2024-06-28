package services

import (
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"strconv"
	"user-management-backend/database"
	"user-management-backend/models"

	sq "github.com/Masterminds/squirrel"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	query := sq.Select("*").From("users")
	sql, args, err := query.ToSql()
	if err != nil {
		log.Errorf("Error while obtaining user info", err.Error())
		return users, err
	}
	rows, err := database.DB.Query(sql, args...)
	if err != nil {
		log.Errorf("Error while obtaining database record", err.Error())
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Errorf("Error while obtaining database info", err.Error())
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	query := database.PSQL.Insert("users").Columns("name", "email").Values(user.Name, user.Email).Suffix("RETURNING id")
	sql, args, err := query.ToSql()
	if err != nil {
		log.Errorf("Error while creating user info", err.Error())
		return err
	}
	err = database.DB.QueryRow(sql, args...).Scan(&user.ID)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				log.Errorf("Error while creating user info: %s", errors.New("username already exists"))
				return errors.New("username already exists")
			}
		}
		log.Errorf("Error while creating user info", err.Error())
		return err
	}
	return nil
}

func UpdateUser(id string, user *models.User) error {
	// Convert the id from string to int64
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Errorf("Error while converting id to int64", err.Error())
		return err
	}

	query := database.PSQL.Update("users").Set("name", user.Name).Set("email", user.Email).Where(sq.Eq{"id": userID})
	sql, args, err := query.ToSql()
	if err != nil {
		log.Errorf("Error while updating user info", err.Error())
		return err
	}
	_, err = database.DB.Exec(sql, args...)
	if err != nil {
		log.Errorf("Error while updating user info", err.Error())
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	// Convert the id from string to int64
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Errorf("Error while converting id to int64", err.Error())
		return err
	}

	query := database.PSQL.Delete("users").Where(sq.Eq{"id": userID})
	sql, args, err := query.ToSql()
	if err != nil {
		log.Errorf("Error while deleting user info", err.Error())
		return err
	}
	result, err := database.DB.Exec(sql, args...)
	if err != nil {
		log.Errorf("Error while deleting user info", err.Error())
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Error while getting rows affected", err.Error())
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows in result set")
	}
	return nil
}
