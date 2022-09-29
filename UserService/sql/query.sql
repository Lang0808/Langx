-- name: CreateUser :execresult
INSERT INTO users (username)
VALUES
(?);

-- name: AddPassword :exec
INSERT INTO passwords (UserId, Password)
VALUES
(?, ?);

-- name: GetUser :one
SELECT * FROM users WHERE username=?;

-- name: GetPassword :one
SELECT * FROM passwords WHERE UserId=?;
