[Home](https://github.com/nikovacevic/bubble) | [Docs](https://github.com/nikovacevic/bubble/blob/master/docs/)

## Spec

### Auth
```
Request magic link     GET     /auth/login?email=$1
Log in                 POST    /auth/login {email,token}
Log out                POST    /auth/logout
```

### Account
Get a list of accounts
```
List accounts          GET     /account
Create account         POST    /account {account}
Edit account           PUT     /account/:id {account}
Delete account         DELETE  /account/:id
```

### Destination
```
List destinations      GET     /dst/
Create destination     POST    /dst/ {destination}
Edit destination       PUT     /dst/:id {destination}
Delete destination     DELETE  /dst/:id
Explore destination    GET     /dst/:id/:view
                       GET     /dst/:id/timeline
                       GET     /dst/:id/calendar
                       GET     /dst/:id/search?query=$1&sort=$2
Message destination    POST    /dst/:id/message {message}
```

### Message
```
Read message           GET     /msg/:id
Edit message           PUT     /msg/:id
Delete message         DELETE  /msg/:id
```
