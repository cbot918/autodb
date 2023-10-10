
-- name: Getuser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: Listusers :many
SELECT * FROM users
ORDER BY name;

-- name: Createuser :one
INSERT INTO users (
	name, bio
) VALUES (
	$1, $2
)
RETURNING *;

-- name: Deleteuser :exec
DELETE FROM users
WHERE id = $1;
