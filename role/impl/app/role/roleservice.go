package role

import (
	"context"
	"portal_back/role/api/internalapi/model"
	"portal_back/role/impl/domain"
)

type RoleService interface {
	GetAllRoles(context context.Context) ([]domain.Role, error)
	GetUserRoles(context context.Context, userId int) ([]domain.Role, error)
	AssignRoleToUser(context context.Context, roleId, userId int) error
	RemoveRoleFromUser(context context.Context, roleId, userId int) error
	IsUserHasRole(context context.Context, accountId int, roleType model.RoleType) (bool, error)
}

type service struct {
	repo RoleRepository
}

func NewService(repo RoleRepository) RoleService {
	return &service{repo: repo}
}

func (service *service) GetAllRoles(context context.Context) ([]domain.Role, error) {
	return service.repo.GetAllRoles(context)
}

func (service *service) GetUserRoles(context context.Context, userId int) ([]domain.Role, error) {
	return service.repo.GetUserRoles(context, userId)
}

func (service *service) AssignRoleToUser(context context.Context, roleId, userId int) error {
	return service.repo.AssignRoleToUser(context, roleId, userId)
}

func (service *service) RemoveRoleFromUser(context context.Context, roleId, userId int) error {
	return service.repo.RemoveRoleFromUser(context, roleId, userId)
}

func (service *service) IsUserHasRole(context context.Context, accountId int, roleType model.RoleType) (bool, error) {
	return service.repo.IsUserHasRole(context, accountId, roleType)
}
