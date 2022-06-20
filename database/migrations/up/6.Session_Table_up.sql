CREATE TABLE IF NOT EXISTS session (
    userid integer unique not null,
    uuid  text unique not null,
    FOREiGN KEY (userid) REFERENCES users(id),
    expires_at timestamp not null
)