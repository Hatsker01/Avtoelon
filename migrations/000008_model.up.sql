CREATE TABLE models
( 
    id BIGSERIAL,
    marc_id BIGINT,
    name text ,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)