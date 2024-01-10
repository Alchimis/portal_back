CREATE TABLE employeeAccount (
                                 id SERIAL PRIMARY KEY NOT NULL,
                                 userId INT NOT NULL,
                                 companyId INT NOT NULL,
                                 firstName VARCHAR(256) NOT NULL,
                                 secondName VARCHAR(256) NOT NULL,
                                 surname VARCHAR(256) NOT NULL,
                                 telephoneNumber VARCHAR(256) NOT NULL,
                                 avatarUrl VARCHAR(256) NOT NULL,
                                 dateOfBirth DATE NOT NULL,
                                 job VARCHAR(256) NOT NULL,
                                 UNIQUE (userId, companyId),
                                 FOREIGN KEY(userId) REFERENCES auth_user(id),
                                 FOREIGN KEY(companyId) REFERENCES company(id)
);