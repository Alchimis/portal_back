package mapper

import (
	frontendapi "portal_back/company/api/frontend"
	"portal_back/company/impl/domain"
)

func MapRoles(departments []domain.RoleInfo) []frontendapi.RoleInfo {
	var result []frontendapi.RoleInfo
	for _, d := range departments {
		result = append(result, frontendapi.RoleInfo{Id: d.Id, Name: d.Name})
	}
	return result
}
