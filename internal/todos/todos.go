package todos

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Service - the struct for our todos service
type Service struct {
	DB *gorm.DB
}

// Todo - defines the Todo model
type Todo struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Description string    `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	CreatedAt   time.Time `json:"created" db:"created_at"`
	UpdatedAt   time.Time `json:"updated" db:"updated_at"`
}

// TodoService - the interface for our Todo service
type TodoService interface {
	GetTodo(ID uuid.UUID) (Todo, error)
	CreateTodo(Todo Todo) (Todo, error)
	UpdateTodo(ID uuid.UUID, newTodo Todo) (Todo, error)
	DeleteTodo(ID uuid.UUID) error
	GetAllTodos() ([]Todo, error)
}

// NewService - returns new todos service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetTodo - retrieves a Todo by their ID from the db
func (s *Service) GetTodo(ID uuid.UUID) (Todo, error) {
	var todo Todo
	if result := s.DB.First(&todo, ID); result.Error != nil {
		return Todo{}, result.Error
	}
	return todo, nil
}

// CreateTodo - adds a new Todo to the database
func (s *Service) CreateTodo(todo Todo) (Todo, error) {
	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	if result := s.DB.Create(&todo); result.Error != nil {
		return Todo{}, result.Error
	}
	return todo, nil
}

// UpdateTodo - updates a Todo by ID with new Todo info
func (s *Service) UpdateTodo(ID uuid.UUID, newTodo Todo) (Todo, error) {
	todo, err := s.GetTodo(ID)
	if err != nil {
		return Todo{}, err
	}

	todo.UpdatedAt = time.Now()
	if result := s.DB.Model(&todo).Updates(newTodo); result.Error != nil {
		return Todo{}, result.Error
	}

	return todo, nil
}

// DeleteTodo - deletes a Todo from the database by ID
func (s *Service) DeleteTodo(ID uuid.UUID) error {
	if result := s.DB.Delete(&Todo{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllTodos - retrieves all Todos from the db
func (s *Service) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	if result := s.DB.Find(&todos); result.Error != nil {
		return []Todo{}, result.Error
	}
	return todos, nil
}
