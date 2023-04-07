## REAT API

This is a REST API service written in Go that provides several endpoints to perform:

# Endpoints
1. /rest/substr/find

This endpoint is used to find the maximum substring that does not contain any repeating characters. The input string should consist of Latin letters (uppercase and lowercase) and digits.
Endpoint body:
*json*
`{
    "text": "asdasd"
}`

2. /rest/email/check

This endpoint is used to search for a string in the following format: "Email: email", where any number of white spaces (including line breaks) can be present instead of "", and the $email should be a string that looks like a valid email.
Endpoint body:
*json*
 `{
     "emails": ["email@mail.ru"]
 }`

3. /rest/iin/check

An IIN checker is a similar functionality that searches for a sequence of digits that is a valid Individual Identification Number (IIN) instead of email addresses.
Endpoint body:
*json*
`
{
    "iin": 940623125434
}
`

Unit tests for each endpoint are available in separate files. These tests can be run using the go test command.
`make test`

To run the server, use the following command:
`make run`

The server will start listening on port 8080 by default. This can be changed by setting the PORT environment variable in config.yml before running the server.