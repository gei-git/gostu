package service

import (
	"fmt"

	"github.com/gei-git/todo-cli/internal/model"
	"github.com/gei-git/todo-cli/internal/storage"
)

type TaskService struct {
	store *storage.TaskStore
}

func NewTaskService() (*TaskService, error) {
	store, err := storage.NewTaskStore()
	if err != nil {
		return nil, fmt.Errorf("初始化存储失败: %w", err)
	}
	return &TaskService{store: store}, nil
}

func (s *TaskService) Close() error {
	return s.store.Close()
}

// AddTask 添加新任务
func (s *TaskService) AddTask(title string) (*model.Task, error) {
	if title == "" {
		return nil, fmt.Errorf("任务标题不能为空")
	}

	task := model.NewTask(title)
	if err := s.store.Create(task); err != nil {
		return nil, fmt.Errorf("保存任务失败: %w", err)
	}
	return task, nil
}

// ListTasks 获取所有任务
func (s *TaskService) ListTasks() ([]*model.Task, error) {
	return s.store.List()
}
