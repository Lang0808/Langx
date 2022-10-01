-- name: CreateBlog :execresult
INSERT INTO blogs ( authorId, title, content, active, timeCreate)
VALUES
(?, ?, ?, ?, ?);
