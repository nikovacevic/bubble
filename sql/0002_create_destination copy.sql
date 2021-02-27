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

INSERT INTO destination_type {

}

/*
(=) destination name
(-) field
(+) permission
(>) messages, "?" indicating optional
*/

/*

# Top-level types

= Document
(can be viewed as page)
- author
- title
- content
- slug
- tags
+ read: [0, 1, 2, 3, 4]
+ comment?: [1, 2, 3, 4]
+ write: [1, 2, 3, 4]
+ moderate: [2, 3, 4]
> comments?

= Venue
(can be viewed as timeline, calendar, search, etc.)
- owner
- name
- description
- slug
- tags
+ read: [0, 1, 2, 3, 4]
+ write: [1, 2, 3, 4]
+ moderate: [2, 3, 4]
> posts
> events
> polls

= Group
(can be viewed as timeline, search)
- owner
- name
- description
- slug
- tags
+ moderate: [2, 3, 4]
> direct

# Intermediate-level types

= Post

= Event

= Poll

*/

/*

Destination API
===============
List destinations     GET    /dst/
Create destination    POST   /dst/ {destination}
Edit destination      PUT    /dst/:id {destination}
Delete destination    DELETE /dst/:id
Explore destination   GET    /dst/:id/:view
                      GET    /dst/:id/timeline
                      GET    /dst/:id/calendar
                      GET    /dst/:id/search?query=$1&sort=$2
Message destination   POST   /dst/:id/message {message}

Message API
===========
Read message          GET    /msg/:id
Edit message          PUT    /msg/:id
Delete message        DELETE /msg/:id

*/

INSERT INTO migration (step) VALUES (2);
