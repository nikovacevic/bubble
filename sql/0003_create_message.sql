/*
Create message tables. A "message" is a super-type for a communication that you
can send to a destination. Different message types can have different meanings
based on their types and the destinations to which they are sent.
*/

CREATE TABLE IF NOT EXISTS message_type (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
);

CREATE TABLE IF NOT EXISTS message (
    id BIGSERIAL PRIMARY KEY,
    message_type_id INT NOT NULL,
    from_account_id INT NOT NULL,
    to_destination_id BIGINT NOT NULL,
    FOREIGN KEY (message_type_id) REFERENCES message_type (id),
    FOREIGN KEY (from_account_id) REFERENCES account (id),
    FOREIGN KEY (to_destination_id) REFERENCES to_destination (id)
);

/* TODO: INSERT "post" (sends to a page) */
/* TODO: INSERT "comment" (sends to a post, event, or poll) */
/* TODO: INSERT "direct" (sends to a group) */
/* TODO: INSERT "rsvp" (sends to an event) */
/* TODO: INSERT "vote" (sends to a poll) */

INSERT INTO migration (step) VALUES (3);
