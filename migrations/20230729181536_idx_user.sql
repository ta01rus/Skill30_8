-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX IDX_UNQ_USER_NAME ON USERS(NAME);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IDX_UNQ_USER_NAME;
-- +goose StatementEnd
