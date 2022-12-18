CREATE TYPE access AS ENUM (
    'banned', 
    'admin', 
    'read'
);

CREATE TABLE privileges (
    user_id integer REFERENCES users (id),
    permission access
);
