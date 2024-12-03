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
BEGIN;
ALTER TABLE attendance
RENAME COLUMN verse TO metaverse;

ALTER TABLE attendance
RENAME COLUMN public_key TO address;
END;

-- Migration #7
BEGIN;
ALTER TABLE attendance
RENAME COLUMN id TO uuid;

ALTER TABLE attendance
ADD COLUMN id int;
COMMIT;

-- Migration #8
BEGIN;
UPDATE attendance
SET location = CASE 
    WHEN location IS NULL THEN '137,-2'
    ELSE location || '137,-2'
END;
COMMIT;
 
-- Migration #9
BEGIN;
ALTER TABLE attendance
ADD COLUMN entrance_status VARCHAR(5) CHECK (entrance_status IN ('ENTER', 'EXIT')) DEFAULT 'ENTER';

COMMENT ON COLUMN attendance.entrance_status IS 'Status indicating whether the person entered or exited.';
COMMIT;
