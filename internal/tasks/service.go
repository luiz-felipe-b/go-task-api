package tasks

type TaskService struct {
	Repo *TaskRepository
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.Repo.GetAll()
}

func (s *TaskService) CreateTask(task Task) (int, error) {
	return s.Repo.Create(task)
}
