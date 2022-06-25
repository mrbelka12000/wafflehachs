CREATE TABLE IF NOT EXISTS Rooms(
ID  serial primary key,
ClientId integer not null,
PsychoId integer not null,
Expires timestamp not null,
IsStarted boolean not null,
FOREiGN KEY(PsychoID) REFERENCES Psychologists(ID),
FOREiGN KEY(ClientID) REFERENCES Clients(ID)
)