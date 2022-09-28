-- name: CreateUser :execresult
INSERT INTO users (username)
VALUES
(?);

-- name: AddPassword :exec
INSERT INTO passwords (UserId, Password)
VALUES
(?, ?);
