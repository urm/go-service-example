package tasks

import "github.com/urm/go-service-example/internal/ports"

type TaskService struct {
	storage ports.Storage
	sender  ports.MessageSender
}

func New(storage ports.Storage, sender ports.MessageSender) *TaskService {
	return &TaskService{storage: storage, sender: sender}
}
