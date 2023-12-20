CREATE TABLE user_sections_prefs
(
    user_id    INTEGER NOT NULL,
    section_id INTEGER NOT NULL,
    FOREIGN KEY (section_id) REFERENCES sections (id),
    PRIMARY KEY (user_id, section_id)
)