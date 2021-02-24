/*
Create account and role tables. An account is a medium for interacting with
every other concept. A role is a permission that allows an account to have
a broad type of access to most, though perhaps not all, concepts. Each
account always starts with no roles, with the exception of the first account
to be created, which starts with the 'owner' role.
*/

CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW()
);

/* https://x-team.com/blog/storing-secure-passwords-with-postgresql/ */
/*
    INSERT INTO users (email, password) VALUES (
        'johndoe@mail.com',
        crypt('johnspassword', gen_salt('bf'))
    );

    SELECT id
    FROM users
    WHERE email = 'johndoe@mail.com'
    AND password = crypt('johnspassword', password);
*/

CREATE TABLE role (
   id SERIAL PRIMARY KEY,
   name TEXT UNIQUE NOT NULL
);

/* TODO */
INSERT INTO role (name) VALUES
    ('owner'),         /* all */
    ('administrator'), /* read, write, edit, administrate */
    ('moderator'),     /* read, write, edit */
    ('member'),        /* read, write */
    ('visitor');       /* read */

CREATE TABLE account_role (
  account_id INT NOT NULL,
  role_id INT NOT NULL,
  created TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY (account_id, role_id),
  FOREIGN KEY (account_id) REFERENCES account (id),
  FOREIGN KEY (role_id) REFERENCES role (id)
);

INSERT INTO migration (step) VALUES (1);
