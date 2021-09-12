package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Todo struct {
	Id     int64  `json:"id"`
	UserId string `json:"userId"`
	// Description string    `json:"description" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Done        int8      `json:"done"`
}

func (t *Todo) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
