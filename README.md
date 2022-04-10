## Routes
***/login***
Accepts username/password and returns auth cookie valid for an hour

***/logout***
Clears server record of cookie

***/dashboard***
Returns data object with recent matches and local events

***/rendezvous/{id}***
Returns information for a match

***/event***
Data object with local events

***/profile***
User profile edit screen

## Event Generation

Every day at 3pm, find clusters of people, add them to event, send event to mobile
