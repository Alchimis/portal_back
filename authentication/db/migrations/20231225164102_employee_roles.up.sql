CREATE TABLE employee_roles (
                                accountId INT NOT NULL,
                                roleId INT NOT NULL,
                                FOREIGN KEY(accountId) REFERENCES employeeAccount(id),
                                FOREIGN KEY(roleId) REFERENCES role(id),
                                PRIMARY KEY(accountId, roleId)
);