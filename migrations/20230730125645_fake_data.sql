-- +goose Up
-- +goose StatementBegin
INSERT INTO USERS(name) VALUES ('mr. Foo');
INSERT INTO USERS(name) VALUES ('ms. Foo');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM USERS;
-- +goose StatementEnd
