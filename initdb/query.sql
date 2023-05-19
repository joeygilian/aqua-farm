CREATE TABLE farm (
 id SERIAL PRIMARY KEY,
 name VARCHAR NOT NULL,
);

INSERT INTO farm (id, name)
VALUES 
(1, 'farm1'),
(2, 'farm2')