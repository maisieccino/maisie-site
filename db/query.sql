-- name: GetItem :one
SELECT * FROM coffee_map_item
WHERE id = $1
LIMIT 1;

-- name: ListItems :many
SELECT * FROM coffee_map_item
ORDER BY id;

-- name: CreateItem :one
INSERT INTO coffee_map_item (
  id, name,
  item_type, image_url, review_url
) VALUES  (
  $1, $2, $3, $4, $5
) RETURNING *;
