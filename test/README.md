<!-- Token -->
curl --location 'http://localhost:9497/0.1/token' \
--header 'Accent-Session-Type: mobile' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--header 'Authorization: Basic xxxxxxx' \
--data '{
  "access_type": "online",
  "backend": "accent_user",
  "expiration": 7200
}'