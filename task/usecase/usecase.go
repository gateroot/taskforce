package usecase

import (
	"errors"
	"fmt"
	"log"
	domain "taskforce/domain/task"
	"taskforce/task"
)

type TaskUsecase struct {
	gateway task.Gateway
	factory domain.Factory
}

func NewTaskUsecase(gateway task.Gateway, factory domain.Factory) task.Usecase {
	return &TaskUsecase{gateway, factory}
}

func (uc *TaskUsecase) New(title, description string) error {
	id, err := uc.gateway.Create(title, description)
	if err != nil {
		return errors.New(fmt.Sprintf("create task failed: %v", err))
	}
	log.Printf("created task successfully (id: %d)", id)
	return nil
}

func (uc *TaskUsecase) Start(taskId int) error {
	t, err := uc.gateway.Read(taskId)
	if err != nil {
		return errors.New(fmt.Sprintf("read task failed: %v", err))
	}
	t = uc.factory.Start(t)
	err = uc.gateway.Update(t)
	if err != nil {
		return errors.New(fmt.Sprintf("update task failed: %v", err))
	}
	log.Print("updated task successfully")
	return nil
}

func (uc *TaskUsecase) Stop(taskId int) error {
	t, err := uc.gateway.Read(taskId)
	if err != nil {
		return errors.New(fmt.Sprintf("read task failed: %v", err))
	}
	t = uc.factory.Stop(t)
	err = uc.gateway.Update(t)
	if err != nil {
		return errors.New(fmt.Sprintf("update task failed: %v", err))
	}
	log.Print("updated task successfully")
	return nil
}

func (uc *TaskUsecase) Pause(taskId int) error {
	t, err := uc.gateway.Read(taskId)
	if err != nil {
		return errors.New(fmt.Sprintf("pause task failed: %v", err))
	}
	t = uc.factory.Pause(t)
	err = uc.gateway.Update(t)
	if err != nil {
		return errors.New(fmt.Sprintf("update task failed: %v", err))
	}
	log.Print("paused task successfully")
	return nil
}

func (uc *TaskUsecase) Complete(taskId int) error {
	t, err := uc.gateway.Read(taskId)
	if err != nil {
		return errors.New(fmt.Sprintf("read task failed: %v", err))
	}
	t = uc.factory.Complete(t)
	err = uc.gateway.Update(t)
	if err != nil {
		return errors.New(fmt.Sprintf("update task failed: %v", err))
	}
	log.Print("updated task successfully")
	return nil
}

func (uc *TaskUsecase) Close(taskId int) error {
	t, err := uc.gateway.Read(taskId)
	if err != nil {
		return errors.New(fmt.Sprintf("read task failed: %v", err))
	}
	t = uc.factory.Close(t)
	err = uc.gateway.Update(t)
	if err != nil {
		return errors.New(fmt.Sprintf("update task failed: %v", err))
	}
	log.Print("updated task successfully")
	return nil
}

func (uc *TaskUsecase) Delete(taskId int) error {
	err := uc.gateway.Delete(taskId)
	if err != nil {
		return errors.New(fmt.Sprintf("delete task failed: %v", err))
	}
	log.Printf("deleted task successfully (id: %d)", taskId)
	return nil
}
