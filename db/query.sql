-- name: GetItem :one
SELECT * FROM coffee_map_item
WHERE id = $1
LIMIT 1;

-- name: ListItems :many
SELECT * FROM coffee_map_item
ORDER BY id;

-- name: SearchByArea :many
SELECT * FROM
coffee_map_item
WHERE ST_WITHIN(location, ST_GEOMFROMEWKB($1::bytea));

-- name: CreateItem :one
INSERT INTO coffee_map_item (
    id,
    item_name,
    item_type,
    image_url,
    review_url,
    summary,
    location
) VALUES (
    gen_random_uuid(),
    $1, $2, $3, $4, $5, $6
) RETURNING *;
