-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS brands (
	"id" VARCHAR(255) NOT NULL PRIMARY KEY,
	"title" VARCHAR(255),
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS brands;
-- +goose StatementEnd
