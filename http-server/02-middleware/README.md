# 02 - Middleware

**Note**: The starting point can be the solution to the exercise 01, where you had
to create a single file HTTP API with some basic endpoints. If you did not submit
it, feel free to grab a solution from someone else to implement this. The file
placed in this exercise also contain a simple case for you to apply what is requested.

In this exercise we will take a look at Middleware in HTTP APIs. Simply put, middleware
is code that runs in the middle of the request and the expected response. For example,
authentication and authorization can be seen as a middleware, since it happens after
the request and before the response is returned to the user that requested it.

Middleware can be whatever you want. In this exercise we will implement two pieces of
middleware: one specific for routes, and one generic for the whole API.

In our case we can see the middleware as "just another function" and we will
implement it this way. Basically it will receive parameters and a `http.HandlerFunc`
function, and it will return a `http.HandlerFunc`. It can take additional paramters
too if we want.

## Route Middleware

Create two methods that will make a route return error if the method sent to it
is not allowed:

* `AllowMethods` function: accepts a list of strings that represents methods,
if the incoming request does not match any of the methods listed, returns an error.

* `handleFunc` function: similar to allow methods, but it accepts a single method
and if the request does not match, it returns an error.

## Global Middleware

Create one method that logs some information about the request into the console.
For example`LoggingMiddleware`.

You probably used the `http.ListenAndServe` method to run your HTTP server. When
using this global middleware.

```
http.ListenAndServe(":5000", LoggingMiddleware(http.DefaultServeMux.ServeHTTP))
```

This will make it work.
