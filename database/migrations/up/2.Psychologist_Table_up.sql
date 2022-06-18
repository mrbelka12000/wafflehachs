CREATE TABLE IF NOT EXISTS Psychologists (
ID  serial primary key,
FirstName text not null,
LastName text not null,
Nickname text Unique not null,
Avatarurl text,
Email text Unique not null,
Age integer not null,
Password text not null
)