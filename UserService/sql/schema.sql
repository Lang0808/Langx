CREATE TABLE users (
    id INTEGER AUTO_INCREMENT,
    username NVARCHAR(50) NOT NULL,
    isAdmin BIT NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(username),
    INDEX(isAdmin)
);

CREATE TABLE passwords (
    UserId INTEGER,
    Password NVARCHAR(50) NOT NULL,
    PRIMARY KEY(UserId)
);