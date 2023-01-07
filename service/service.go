package service

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"

	"github.com/pp3times/assessment/models"
)

type Expenses struct {
	ID uint `gorm:"primary key;autoIncrement" json:"id"`

	Title string `json:"title"`

	Amount float64 `json:"amount"`

	Note string `json:"note"`

	Tags []string `json:"tags" gorm:"serializer:json"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {

	app.Post("/expenses", r.CreateExpense)

	app.Get("/expenses/:id", r.GetExpenseByID)

}

// GET Expense By ID

func (r *Repository) GetExpenseByID(context *fiber.Ctx) error {

	id := context.Params("id")

	expenseModel := &models.Expenses{}

	if id == "" {

		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{

			"message": "id cannot be empty",
		})

		return nil

	}

	err := r.DB.Where("id = ?", id).First(expenseModel).Error

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not get expense"})

		return err

	}

	context.Status(http.StatusOK).JSON(expenseModel)

	return nil

}

// CREATE EXPENSE

func (r *Repository) CreateExpense(context *fiber.Ctx) error {

	expense := Expenses{}

	err := context.BodyParser(&expense)

	if err != nil {

		context.Status(http.StatusUnprocessableEntity).JSON(

			&fiber.Map{"message": "request failed"})

		return err

	}

	validator := validator.New()

	err = validator.Struct(Expenses{})

	if err != nil {

		context.Status(http.StatusUnprocessableEntity).JSON(

			&fiber.Map{"message": err},
		)

		return err

	}

	err = r.DB.Create(&expense).Error

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not create expense"})

		return err

	}

	// context.Status(http.StatusOK).JSON(&fiber.Map{

	// 	"message": "expense has been successfully added",
	// })

	context.Status(http.StatusCreated).JSON(expense)

	return nil

}
