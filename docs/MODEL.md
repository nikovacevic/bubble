[Home](https://github.com/nikovacevic/bubble) | [Model](./MODEL.md) | [Resources](./RESOURCES.md)

## Model

### Modules
- Auth: provides temporary access to a user
- User: provides access to view and edit user-related resources
- Venue: provides access to viewing, creating, and editing venues for sending messages
  - Page: public or private repository for posts, events, polls, etc.
  - Group: private repository for direct messages among select Users
  - Event: public or private calendar event, repository for Invitations and RSVPs (also a Message)
  - Poll: public or private respository for votes (also a Message)
- Message: provides users ability to send messages in certain venues
  - Post: public or private message belonging to a page or event
  - Direct: private message belonging to a group
  - Event: public or private calendar event (also a Venue)
  - Invitation: public or private invitation to join an event
  - RSVP: public or private message belonging to an event
  - Poll: public or private message with vote options (also a Venue)
  - Vote: public or private message belonging to a poll
- Media: provides an API for storing and retrieving media (files) by URI
  - Document
  - Image

_Note: see [pkg/bubble](https://github.com/nikovacevic/bubble/tree/master/pkg/bubble) for implementations of domain types._

#### Auth
Authenticate users using magic links. Respond to token-based authentication with

Request a magic link sent to the given
```
GET /auth/login?email={email}
```
Authenticate with OTP using the auth link
```
GET /auth/login?email={email}&token={token}
```
Log out of current session
```
POST /auth/logout
```

#### User
We should try to maintain as little personal information as possible. Perhaps just ID, email, and name to start. (And photo?)

How should access control work? Roles and resources? Just roles? How do we tie a user to a venue with permissions?

Get a list of users
```
GET /user
```
Add a new user (e.g. invite a user to join via email?)
```
POST /user
```
Update a user (keep edit history?)
```
PUT /user/{id}
```
Delete a user
```
DELETE /user/{id}
```

#### Venue
This is the main set of communication features: opening new pages and group chats; sending messages; creating and responding to events. Again, simple is best, and if possible it would be nice to keep all of it consolidated to a simple and robust API. (Perhaps we'll have to factor out some of the types that it's a stretch to call a "venue" or a "message", e.g. an event, like but we'll see.)

Get a list of venues by type
```
GET /venue?type={type}
```
Create a new venue
```
POST /venue
```
Update a venue (keep edit history?)
```
PUT /venue/{id}
```
Get the given venue and its messages
```
GET /venue/{id}
```
Send a message to the given venue
```
POST /venue/{id}/message
```

#### Message
Most message creation will happen through venues, but we need to allow interacting directly with messages (e.g. linking directly to a specific message; editing; deleting).

View a message with it's venue included for context
```
GET /message/{id}
```
Update a message (keep edit history?)
```
PUT /message/{id}
```
Delete a message
```
DELETE /message/{id}
```

#### Media
I don't love this name. Documents? Objects? Files?

Get a piece of media
```
GET /media/{id}
```
Upload a new piece of media
```
POST /media/{id}
```
Delete a piece of media
```
DELETE /media/{id}
```
