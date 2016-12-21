
-- +goose Up
CREATE TABLE task_manager(
    "id" serial PRIMARY key NOT NULL,
    "task" text NOT NULL,
    "userId" text NOT NULL
);

-- SQL in section 'Up' is executed when this migration is applied


-- +goose Down
drop table task_manager;
-- SQL section 'Down' is executed when this migration is rolled back

