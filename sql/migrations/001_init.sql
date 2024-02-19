CREATE TABLE "users" (
    ID UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255)
);

---- create above / drop below ----

DROP TABLE "users";
