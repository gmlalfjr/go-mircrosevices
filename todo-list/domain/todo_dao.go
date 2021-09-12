package domain

import (
	"github.com/go-microservices/todo-list/repositories"
	errors "github.com/go-microservices/todo-list/utils"
)

func (t *Todo) TodoDao() *errors.RestErr {
	if err := repositories.Client.Model(&t).Select("description", "done", "created_at", "modified_at", "user_id").Create(&t); err.Error != nil {
		return errors.NewBadRequest("Failed Create Data")
	}

	return nil
}

func (t *Todo) TodoDaoUpdate() *errors.RestErr {
	if err := repositories.Client.Model(&t).Select("description", "done", "modified_at").Where(&t.Id).Updates(Todo{
		Description: t.Description,
		Done:        t.Done,
		ModifiedAt:  t.ModifiedAt,
	}); err.Error != nil {
		return errors.NewBadRequest("Failed Create Data")
	}

	return nil
}

func (t *Todo) GetTodoDetailDao() *errors.RestErr {
	if result := repositories.Client.Model(&t).Where("id = ?", t.Id).First(&t); result.Error != nil {
		return errors.NewInternalServerError("Something Bad Happened")
	}

	return nil
}

func (t *Todo) GetTodoDao() ([]Todo, *errors.RestErr) {
	results := make([]Todo, 0)
	result := repositories.Client.Table("todos").Where("user_id = ?", t.UserId).Where("done = ?", 0).Find(&results)
	if result.Error != nil {
		return nil, errors.NewInternalServerError("Error Get Data")
	}

	return results, nil
}

func (t *Todo) DeleteTodoDao() *errors.RestErr {
	result := repositories.Client.Table("todos").Delete(&Todo{}, &t.Id)
	if result.Error != nil {
		return errors.NewInternalServerError("Error Get Data")
	}

	return nil
}
