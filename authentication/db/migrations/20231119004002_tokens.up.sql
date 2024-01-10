CREATE TABLE tokens
(
    id serial primary key,
    user_id integer NOT NULL,
    token character varying(256) NOT NULL
);