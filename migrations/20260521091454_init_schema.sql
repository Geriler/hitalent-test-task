-- +goose Up
CREATE TABLE departments (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    parent_id BIGINT REFERENCES departments(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX departments_name_parent_id_unique
ON departments(name, parent_id)
WHERE parent_id IS NOT NULL;

CREATE UNIQUE INDEX departments_name_null_parent_unique
ON departments(name)
WHERE parent_id IS NULL;

CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    department_id BIGINT NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
    full_name VARCHAR(200) NOT NULL,
    position VARCHAR(200) NOT NULL,
    hired_at DATE,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE employees;

DROP INDEX departments_name_null_parent_unique;

DROP INDEX departments_name_parent_id_unique;

DROP TABLE departments;
