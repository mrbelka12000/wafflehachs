CREATE TABLE IF NOT EXISTS Clients (
ID  serial primary key,
FirstName text not null,
LastName text not null,
UserName text Unique not null,
Avatarurl text,
Email text Unique not null,
Age integer not null,
Warnings integer,
Password text not null
)