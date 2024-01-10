CREATE TABLE department (
                            id SERIAL PRIMARY KEY NOT NULL,
                            name VARCHAR(256) NOT NULL,
                            parentDepartmentId INT,
                            companyId INT NOT NULL,
                            supervisorId INT NOT NULL,
                            FOREIGN KEY(parentDepartmentId) REFERENCES department(id),
                            FOREIGN KEY(companyId) REFERENCES company(id),
                            FOREIGN KEY(supervisorId) REFERENCES employeeAccount(id)
);