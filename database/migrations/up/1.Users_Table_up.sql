CREATE TABLE IF NOT EXISTS Users (
ID  serial primary key,
FirstName text not null,
LastName text not null,
UserName text Unique not null,
Avatarurl text,
Email text Unique not null,
Age integer not null,
Password text not null
)