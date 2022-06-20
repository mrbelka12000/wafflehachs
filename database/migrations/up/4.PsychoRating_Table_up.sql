CREATE TABLE IF NOT EXISTS PsychoRate(
    PsychoID integer not null,
    ClientID integer not null,
    Rate integer not null,
    Comment text,
    Anonym BOOLEAN not null,
    FOREiGN KEY(PsychoID) REFERENCES Psychologists(ID),
    FOREiGN KEY(ClientID) REFERENCES Clients(ID)
)