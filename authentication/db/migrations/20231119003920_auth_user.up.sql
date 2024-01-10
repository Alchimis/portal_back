CREATE TABLE auth_user
(
    id serial primary key,
    login character varying(256) NOT NULL,
    salt character varying(256) NOT NULL,
    password character varying(256) NOT NULL
);