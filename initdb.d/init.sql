CREATE TABLE IF NOT EXISTS tasks (
    id INT NOT NULL,
    name VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO tasks (id, name) VALUES
    (1, 'hoge'),
    (2, 'fuga'),
    (3, 'foo'),
    (4, 'bar')
;
