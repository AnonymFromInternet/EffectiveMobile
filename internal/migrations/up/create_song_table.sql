CREATE TABLE IF NOT EXISTS song (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    release_date DATETIME,
    artist_id INTEGER,
    song_text varchar(255),
    link varchar(255),
    FOREIGN KEY (artist_id) REFERENCES artist(id)
)
