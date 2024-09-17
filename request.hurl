# Post a TODO
POST https://todo.com/create
x-on-behalf-of: acc_12345
x-api-version: 2024-06-30
{
  "text": "go shoping"
  # "some": "comment"
}
HTTP 200
[Captures]
id: jsonpath "$['id']"
sca-token: header sca-token-session