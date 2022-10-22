# 04 - Testing

Time for testing has come.

For this exercise, you are given a basic app that answers to the `/hostname`
endpoint, returning the hostname of the computer is located at. Your task here is
to write tests for the only two methods you will find: `NewServer` and `HandleHostname`.

To write those tests you need to use the package `httptest` and its functions
`NewRecorder` and `NewRequest`. The tests should test basic things as content
returned and HTTP code at least.

The `TestHandleHostname` should test the endpoint directly, while the
`TestNewServer` should test the `/` (root) endpoint.
