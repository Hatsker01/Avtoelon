CREATE TABLE cars
(
    id UUID PRIMARY KEY,
    category_id BIGINT,
    model_id BIGINT,
    body_id BIGINT,
    date DATE,
    price BIGINT,
    auction BOOLEAN,
    enginee FLOAT,
    oil_id INT,
    transmission_id INT,
    milage BIGINT,
    color_id BIGINT,
    drive_unit_id INT,
    outside_id INT[],
    optic_id INT[],
    salon_id INT[],
    media_id INT[],
    options_id INT[],
    additionally_id INT[],
    add_info TEXT,
    region_id INT,
    city_id INT,
    phone VARCHAR(10),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);