CREATE TABLE coffee_map_item (
    id text PRIMARY KEY,
    item_name text NOT NULL,
    item_type text,
    image_url text,
    review_url text,
    description text,

    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp
);
