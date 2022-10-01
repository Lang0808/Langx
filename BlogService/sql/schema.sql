CREATE TABLE blogs (
    id INTEGER AUTO_INCREMENT,
    authorId INTEGER NOT NULL,
    title NVARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    active BIT NOT NULL,
    timeCreate BIGINT,
    PRIMARY KEY(id),
    INDEX(active, timeCreate),
    INDEX(active, authorId, timeCreate)
);