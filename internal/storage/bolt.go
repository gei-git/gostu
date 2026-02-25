package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gei-git/todo-cli/internal/model" // ← 注意：如果你的 go.mod module 名不是 gostu，请改成你的实际模块名！

	"go.etcd.io/bbolt"
)

const (
	bucketName = "tasks"
	dbFile     = "tasks.db"
)

// TaskStore BoltDB 任务存储实现（企业级单例模式）
type TaskStore struct {
	db *bbolt.DB
}

// NewTaskStore 创建或打开 BoltDB 文件
func NewTaskStore() (*TaskStore, error) {
	db, err := bbolt.Open(dbFile, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	// 创建 bucket（如果不存在）
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("创建 bucket 失败: %w", err)
	}

	return &TaskStore{db: db}, nil
}

// Close 关闭数据库连接
func (s *TaskStore) Close() error {
	return s.db.Close()
}

// Create 添加新任务
func (s *TaskStore) Create(task *model.Task) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		// 自动生成 ID（最大 ID + 1）
		id, _ := b.NextSequence()
		task.ID = uint(id)

		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put([]byte(fmt.Sprintf("%d", task.ID)), buf)
	})
}

// List 获取所有任务
func (s *TaskStore) List() ([]*model.Task, error) {
	var tasks []*model.Task

	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.ForEach(func(k, v []byte) error {
			var t model.Task
			if err := json.Unmarshal(v, &t); err != nil {
				return err
			}
			tasks = append(tasks, &t)
			return nil
		})
	})
	return tasks, err
}

// MarkDone 标记任务为已完成
func (s *TaskStore) MarkDone(id uint) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		key := []byte(fmt.Sprintf("%d", id))

		v := b.Get(key)
		if v == nil {
			return fmt.Errorf("任务 ID %d 不存在", id)
		}

		var task model.Task
		if err := json.Unmarshal(v, &task); err != nil {
			return err
		}

		task.Done = true
		task.UpdatedAt = time.Now()

		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(key, buf)
	})
}
