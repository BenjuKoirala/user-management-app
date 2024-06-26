package services

import (
	"database/sql"
	"testing"
	"user-management-backend/database"
	"user-management-backend/models"
	//"user-management-backend/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestServices(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Services Suite")
}

var _ = ginkgo.Describe("User Service", func() {
	var (
		db   *sql.DB
		mock sqlmock.Sqlmock
		err  error
	)

	ginkgo.BeforeEach(func() {
		db, mock, err = sqlmock.New()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		database.DB = db
	})

	ginkgo.AfterEach(func() {
		db.Close()
	})

	ginkgo.Describe("GetUsers", func() {
		ginkgo.It("should return a list of users", func() {
			rows := sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "John", "john@example.com").
				AddRow(2, "Alice", "alice@example.com")

			mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

			users, err := GetUsers()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(users).To(gomega.HaveLen(2))
			gomega.Expect(users[0].ID).To(gomega.Equal(1))
			gomega.Expect(users[0].Name).To(gomega.Equal("John"))
			gomega.Expect(users[0].Email).To(gomega.Equal("john@example.com"))
			gomega.Expect(users[1].ID).To(gomega.Equal(2))
			gomega.Expect(users[1].Name).To(gomega.Equal("Alice"))
			gomega.Expect(users[1].Email).To(gomega.Equal("alice@example.com"))
		})
	})

	ginkgo.Describe("CreateUser", func() {
		ginkgo.It("should create a new user", func() {
			user := &models.User{
				Name:  "Alice",
				Email: "alice@example.com",
			}

			mock.ExpectQuery("^INSERT INTO users (.+) VALUES (.+) RETURNING id").WithArgs("Alice", "alice@example.com").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			err := CreateUser(user)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(user.ID).To(gomega.Equal(1))
		})
	})

	ginkgo.Describe("UpdateUser", func() {
		ginkgo.It("should update an existing user", func() {
			user := &models.User{
				ID:    1,
				Name:  "Alice",
				Email: "alice@example.com",
			}

			mock.ExpectExec("^UPDATE users SET name = \\$1, email = \\$2 WHERE id = \\$3").
				WithArgs("Alice", "alice@example.com", int64(1)).
				WillReturnResult(sqlmock.NewResult(0, 1))

			err := UpdateUser("1", user)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})

	ginkgo.Describe("DeleteUser", func() {
		ginkgo.It("should return an error if user does not exist", func() {
			mock.ExpectExec("^DELETE FROM users WHERE id = \\$1").
				WithArgs(int64(1)).
				WillReturnResult(sqlmock.NewResult(0, 0))

			err := DeleteUser("1")
			gomega.Expect(err).To(gomega.MatchError("no rows in result set"))
		})
	})
})
