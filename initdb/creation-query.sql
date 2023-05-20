CREATE TABLE farm (
 id SERIAL PRIMARY KEY,
 name VARCHAR NOT NULL,
);

CREATE TABLE pond (
 id SERIAL PRIMARY KEY,
 farm_id INT,
 name VARCHAR NOT null, 
 constraint fk_farm
 foreign key(farm_id)
 references farm(id)
);