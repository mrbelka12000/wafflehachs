CREATE TABLE IF NOT EXISTS Psychologists (
ID  integer unique,
FOREiGN KEY(ID) REFERENCES users(ID),
Busymode text
)