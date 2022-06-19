CREATE TABLE IF NOT EXISTS session (
    ID serial primary key,
    uuid  text not null,
    expires_at timestamp not null
)