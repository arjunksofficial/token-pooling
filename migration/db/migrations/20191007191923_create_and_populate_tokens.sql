
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS tokens 
(
    token_id INT PRIMARY KEY ,
    token_data VARCHAR(20) , 
    count INT
);

INSERT INTO tokens values
    (1,'token1',0),
    (2,'token2',0),
    (3,'token3',0),
    (4,'token4',0),
    (5,'token5',0),
    (6,'token6',0),
    (7,'token7',0),
    (8,'token8',0),
    (9,'token9',0),
    (10,'token10',0),
    (11,'token11',0),
    (12,'token12',0)
;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE tokens;