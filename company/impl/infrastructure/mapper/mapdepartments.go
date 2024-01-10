package mapper

import (
	frontendapi "portal_back/company/api/frontend"
	"portal_back/company/impl/domain"
)

func MapDepartments(departments []domain.DepartmentInfo) []frontendapi.DepartmentInfo {
	var result []frontendapi.DepartmentInfo
	for _, d := range departments {
		result = append(result, frontendapi.DepartmentInfo{Id: d.Id, Name: d.Name})
	}
	return result
}
