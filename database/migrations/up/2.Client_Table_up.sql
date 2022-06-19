CREATE TABLE IF NOT EXISTS clients (
ID  integer unique,
FOREiGN KEY(ID) REFERENCES users(ID),
Warnings integer
)