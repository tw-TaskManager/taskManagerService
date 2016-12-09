
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE Task_Manager(
    "id" serial PRIMARY key NOT NULL,
    "task" text NOT NULL
);

-- +goose Down
drop table task_manager
-- SQL section 'Down' is executed when this migration is rolled back

