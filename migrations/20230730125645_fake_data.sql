-- +goose Up
-- +goose StatementBegin
INSERT INTO USERS(id, name) VALUES (100,'foo');
INSERT INTO USERS(id, name) VALUES (101,'faa');

INSERT INTO LABELS(id, name) VALUES (1,'B');
INSERT INTO LABELS(id, name) VALUES (2,'C');


   

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM USERS;
DELETE FROM LABELS;



-- +goose StatementEnd
