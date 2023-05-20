CREATE TABLE farm (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    status BOOLEAN NOT null DEFAULT true,
    created_date TIMESTAMPTZ DEFAULT NOW(),
    updated_date TIMESTAMPTZ,
    deleted_date TIMESTAMPTZ
);

CREATE TABLE pond (
 id SERIAL PRIMARY KEY,
 farm_id INT,
 name VARCHAR NOT null, 
 status BOOLEAN NOT null DEFAULT true,
 created_date TIMESTAMPTZ DEFAULT NOW(),
 updated_date TIMESTAMPTZ,
 deleted_date TIMESTAMPTZ,
 constraint fk_farm
 foreign key(farm_id)
 references farm(id)
);