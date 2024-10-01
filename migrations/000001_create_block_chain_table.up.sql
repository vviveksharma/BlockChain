CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;


CREATE TABLE IF NOT EXISTS block_tbl (
    id UUID PRIMARY KEY,
    prev_hash character varying(100),
    created_at timestamptz NOT NULL,
    data character varying(100) NOT NULL,
    hash character varying(100) NOT NULL,
    nonce int8 NOT NULL
);

COMMIT;