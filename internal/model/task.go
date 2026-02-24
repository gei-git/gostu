package model

import "time"

// Priority 任务优先级（企业常用枚举写法）
type Priority string

const (
	PriorityLow    Priority = "low"    // 低优先级
	PriorityMedium Priority = "medium" // 中优先级（默认）
	PriorityHigh   Priority = "high"   // 高优先级
)

// Task 任务模型（核心数据结构，后续所有存储、API 都基于它）
type Task struct {
	ID          uint       `json:"id"`                    // 任务唯一ID（BoltDB 会自动生成）
	Title       string     `json:"title"`                 // 任务标题（必填）
	Description string     `json:"description,omitempty"` // 任务描述（可选）
	Priority    Priority   `json:"priority"`              // 优先级
	DueDate     *time.Time `json:"due_date,omitempty"`    // 截止日期（指针，可为空）
	Tags        []string   `json:"tags,omitempty"`        // 标签列表（如 ["工作", "紧急"]）
	Done        bool       `json:"done"`                  // 是否已完成
	CreatedAt   time.Time  `json:"created_at"`            // 创建时间
	UpdatedAt   time.Time  `json:"updated_at"`            // 最后更新时间
}

// NewTask 创建一个新任务的工厂函数（企业最佳实践）
func NewTask(title string) *Task {
	now := time.Now()
	return &Task{
		Title:     title,
		Priority:  PriorityMedium, // 默认中等优先级
		Done:      false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
