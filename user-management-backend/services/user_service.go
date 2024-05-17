package services

import (
	"user-management-backend/database"
	"user-management-backend/models"

	sq "github.com/Masterminds/squirrel"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	query := sq.Select("*").From("users")
	sql, args, err := query.ToSql()
	if err != nil {
		return users, err
	}
	rows, err := database.DB.Query(sql, args...)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(id string) (models.User, error) {
	var user models.User
	query := sq.Select("*").From("users").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return user, err
	}
	err = database.DB.QueryRow(sql, args...).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	query := database.PSQL.Insert("users").Columns("name", "email").Values(user.Name, user.Email).Suffix("RETURNING id")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	return database.DB.QueryRow(sql, args...).Scan(&user.ID)
}

func UpdateUser(id string, user *models.User) error {
	query := database.PSQL.Update("users").Set("name", user.Name).Set("email", user.Email).Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(sql, args...)
	return err
}

func DeleteUser(id string) error {
	query := database.PSQL.Delete("users").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(sql, args...)
	return err
}
