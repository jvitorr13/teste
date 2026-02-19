package usecase

import "taskMetrics/pkg/database"

type UseCases struct {
	Create *CreateUserUseCase
	List   *ListUsersUseCase
}

func NewUseCases(process *database.ProcessManager) *UseCases {
	return &UseCases{
		Create: NewCreateUserUseCase(process),
		List:   NewListUsersUseCase(process),
	}
}
