CREATE TYPE roleTypes AS ENUM('VIEWER', 'EDITOR', 'ADMIN', 'OWNER');

CREATE TABLE role (
                      id SERIAL PRIMARY KEY NOT NULL,
                      title character varying(256) NOT NULL,
                      description character varying(256) NOT NULL,
                      roleType roleTypes NOT NULL
);