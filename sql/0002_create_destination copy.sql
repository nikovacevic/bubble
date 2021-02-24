/*
Create destination tables. A "destination" is a super-type for a place to which
you can go, and to which you can send something. Different destinations can
have different rules, accepting different interactions.
*/

CREATE TABLE IF NOT EXISTS destination_type (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
);

CREATE TABLE IF NOT EXISTS destination (
    id BIGSERIAL PRIMARY KEY,
    destination_type_id INT NOT NULL,
    owner_account_id INT NOT NULL,
    FOREIGN KEY (destination_type_id) REFERENCES destination_type (id)
    FOREIGN KEY (owner_account_id) REFERENCES account (id)
);

/* TODO: INSERT "page" */
/* TODO: INSERT "group" */
/* TODO: INSERT "event" */
/* TODO: INSERT "poll" */

INSERT INTO migration (step) VALUES (2);
