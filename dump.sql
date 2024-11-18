CREATE TABLE users
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    public_key bytea[] NOT NULL,
    PRIMARY KEY (id)
);

-- Migration #2
ALTER TABLE IF EXISTS users
    ADD COLUMN location text;

COMMENT ON COLUMN users.location
    IS 'decentraland location (x,y)';

-- Migration #3
ALTER TABLE users RENAME TO attendance;

-- Migration #4
ALTER TABLE attendance ALTER COLUMN public_key TYPE TEXT USING public_key::TEXT;

-- Migration #5
ALTER TABLE attendance
ADD COLUMN verse VARCHAR(10) CHECK (verse IN ('dcl', 'hyperfy', 'irl'));

-- Migration #6
ALTER TABLE attendance
RENAME COLUMN verse TO metaverse;

ALTER TABLE attendance
RENAME COLUMN public_key TO address;
