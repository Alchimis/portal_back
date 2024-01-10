CREATE TABLE employee_department (
                                     accountId INT NOT NULL,
                                     departmentId INT NOT NULL,
                                     FOREIGN KEY(accountId) REFERENCES employeeAccount(id),
                                     FOREIGN KEY(departmentId) REFERENCES department(id),
                                     PRIMARY KEY (accountId, departmentId)
);