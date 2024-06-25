-- +goose Up
-- Step 1: Create a new sequence
CREATE SEQUENCE user_id_seq START 1;

-- Step 2: Set the default value of id to use the sequence
ALTER TABLE "users"
  ALTER COLUMN id SET DEFAULT nextval('user_id_seq');

-- Step 3: Add a NOT NULL constraint to id
ALTER TABLE "users"
  ALTER COLUMN id SET NOT NULL;

-- +goose Down
-- Step 1: Drop the default value and NOT NULL constraint
ALTER TABLE "users"
  ALTER COLUMN id DROP DEFAULT,
  ALTER COLUMN id DROP NOT NULL;

-- Step 2: Drop the sequence
DROP SEQUENCE user_id_seq;