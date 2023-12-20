CREATE TABLE sections
(
    id            SERIAL PRIMARY KEY,
    title         CHARACTER VARYING(256) NOT NULL,
    thumbnail_url CHARACTER VARYING(256) NOT NULL,
    is_favorite   BOOLEAN                NOT NULL,
    company_id    INTEGER                NOT NULL
);