# 03 - Modules

Before going into more complex things, we will learn about go modules.

*Note*: if you do not have previous exercise, take the result at
`/http-server/hemartin/02-middleware/main.go` to start this exercise.

For this exercise, you should create a module called `app` that has a `NewHandler`
function that returns something you can pass to `http.ListenAndServe()` seconds
parameter to handle the server requests. You are not expected to change the code
on the `main.go` file.

If you have more time to spend on this, split the code into multiple files
multiple files for the `app` module so you gain more experience working with
modules.
