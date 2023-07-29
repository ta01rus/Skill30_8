-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE VIEW TASK_VIEW AS
SELECT  A.ID, TITLE , A.ASSIGNED_ID , A.AUTHOR_ID , A.CONTENT , A.OPENED, A.CLOSED
        ,C.NAME AS ASSIGNED_NAME, B.NAME AS AUTHOR_NAME
FROM TASKS A
LEFT JOIN USERS B ON B.ID = A.AUTHOR_ID 
LEFT JOIN USERS C ON C.ID = A.ASSIGNED_ID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW TASK_VIEW;
-- +goose StatementEnd
