# Vanilla HTTP Server (Single file)

**Author**: Hector Martinez

**Note**: keep everything into one file, we will work on followup exercises to refactor
the code and split it into different files.


In this exercise you have to create a HTTP server with the standard golang
library (no third party libraries) that accepts a series of verbs on some
endpoints and that should return certain information.

Endpoints and verbs:

| HTTP Verb | Endpoint | Return |
| --- | --- | --- |
| GET | /hostname | The hostname of the machine running the server |
| GET | /nslookup/:hostname | Does a DNS lookup on the hostname and returns the IPs for that host server by the DNS server |
| GET | /hash/?text=:text&algorithm=:algorithm | Returns a JSON with three keys: `hash` which contains the hash of the text created by the algorithm specified, `text` that contains the original text given to the endpoint, and `algorithm` with the algorithm that generated the hash.
| GET | /headers | Returns the headers received from the client |

All endpoints should return the HTTP error for method not supported if other
verbs are requested (PUT, POST, PATCH, etc.)
