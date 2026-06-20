# notification-service


curl --location 'http://localhost:8080/api/v1/notifications' \
--header 'Content-Type: application/json' \
--data '{
    "userId": "user-123",
    "channel": "EMAIL",
    "message": "Welcome to our platform!"
}'