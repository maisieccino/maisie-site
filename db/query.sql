-- name: GetItem :one
SELECT * FROM coffee_map_item
WHERE id = $1 LIMIT 1;
