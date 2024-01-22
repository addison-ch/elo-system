-- name: CreatePicture :one
INSERT INTO pictures (
  url,
  matches,
  rating
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPicture :one
SELECT * FROM pictures
WHERE id = $1 LIMIT 1;

-- name: ListPictures :many
SELECT * FROM pictures
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetRandomPair :many
SELECT * FROM pictures
ORDER BY random()
LIMIT 2;

-- name: UpdatePicture :one
UPDATE pictures
SET rating = $2
WHERE id = $1
RETURNING *;

-- name: DeletePicture :exec
DELETE FROM pictures
WHERE id = $1;

-- name: GetUsersPictures :many
SELECT * FROM pictures
WHERE user_id = $1 LIMIT 3;