package sql

import (
	"context"
	"portal_back/role/api/internalapi/model"
	"portal_back/role/impl/app/role"
	"portal_back/role/impl/domain"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	conn *pgx.Conn
}

func rowsToArray(rows pgx.Rows) []domain.Role {
	var roles []domain.Role
	for rows.Next() {
		var role domain.Role
		rows.Scan(&role.Id, &role.Title, &role.Description, &role.RoleType)
		roles = append(roles, role)
	}
	return roles
}

// AssignRoleToUser implements role.RoleRepository.
func (repo *repository) AssignRoleToUser(context context.Context, roleId, userId int) error {
	query := `
		INSERT INTO employee_roles(accountid, roleid)
		SELECT $1, $2
		WHERE 
			NOT EXISTS(
				SELECT  accountid, roleid FROM employee_roles
				WHERE accountid=$1 AND roleid=$2
			)
	`
	_, err := repo.conn.Query(context, query, userId, roleId)
	return err
}

// GetAllRoles implements role.RoleRepository.
func (repo *repository) GetAllRoles(context context.Context) ([]domain.Role, error) {
	query := `
		SELECT role.id, role.title, role.description, role.roletype 
		FROM role
	`
	rows, err := repo.conn.Query(context, query)
	defer rows.Close()

	roles := rowsToArray(rows)
	if err == pgx.ErrNoRows {
		return []domain.Role{}, nil
	} else if err != nil {
		return nil, err
	}
	return roles, nil
}

// GetUserRoles implements role.RoleRepository.
func (repo *repository) GetUserRoles(context context.Context, userId int) ([]domain.Role, error) {
	query := `
		SELECT role.id, role.title, role.description, role.roletype  FROM role
		RIGHT JOIN employee_roles ON role.id=employee_roles.roleid
		AND employee_roles.accountid=$1
	`
	rows, err := repo.conn.Query(context, query, userId)
	defer rows.Close()

	roles := rowsToArray(rows)
	if err == pgx.ErrNoRows {
		return []domain.Role{}, nil
	} else if err != nil {
		return nil, err
	}
	return roles, nil
}

// RemoveRoleFromUser implements role.RoleRepository.
func (repo *repository) RemoveRoleFromUser(context context.Context, roleId, userId int) error {
	query := `
		DELETE FROM  employee_roles
		WHERE employee_roles.accountid=$1 
		AND employee_roles.roleid=$2
		RETURNING accountid
	`
	var id int
	return repo.conn.QueryRow(context, query, userId, roleId).Scan(&id)
}

func (repo *repository) IsUserHasRole(context context.Context, accountId int, roleType model.RoleType) (bool, error) {
	query := `
		SELECT COUNT(*) FROM employee_roles
		LEFT JOIN role ON employee_roles.roleid=role.id
		WHERE employee_role.accountid=$1 AND role.roletype=$2
	`
	var count int
	err := repo.conn.QueryRow(context, query, accountId, roleType).Scan(&count)
	if err == pgx.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return count == 0, nil
}

func NewRepository(conn *pgx.Conn) role.RoleRepository {
	return &repository{conn: conn}
}
