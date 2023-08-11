# 虎符(oauth2)

### Oauth2 Server

Credentials
```shell
http://localhost/token?grant_type=client_credentials&client_id=asdgasdfrvdsvasd&client_secret=dfgbgsbasdas
```
```json
{
    "code": 200,
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvcmciOiJnYmFzZGZzZmJzZ2FzZGYiLCJ1aWQiOiJhc2RnYXNkZnJ2ZHN2YXNkIiwibmJmIjoxNjkxNzM4NjU0LCJpYXQiOjE2OTE3Mzg2NTQsImV4cCI6MTY5NDMzMDY1NCwianRpIjoiOTM3M2NkYTAtYzQ0ZS00ZDA0LTliMWYtODFjYzM4MDdlYzE3IiwiaXNzIjoid3d3LmRlbW8uY29tIiwic3ViIjoiX19hc2RnYXNkZnJ2ZHN2YXNkIiwiYXVkIjpbIiJdfQ.SBjic7eOaV80QP9WnKpYJ2sAkzJhBD9QqtVikXHJ3SM",
    "tokenType": "bearer",
    "expiresIn": 2592000
}
```

Token Check
```shell
http://localhost/validate
Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvcmciOiJnYmFzZGZzZmJzZ2FzZGYiLCJ1aWQiOiJhc2RnYXNkZnJ2ZHN2YXNkIiwibmJmIjoxNjkxNzM4NjU0LCJpYXQiOjE2OTE3Mzg2NTQsImV4cCI6MTY5NDMzMDY1NCwianRpIjoiOTM3M2NkYTAtYzQ0ZS00ZDA0LTliMWYtODFjYzM4MDdlYzE3IiwiaXNzIjoid3d3LmRlbW8uY29tIiwic3ViIjoiX19hc2RnYXNkZnJ2ZHN2YXNkIiwiYXVkIjpbIiJdfQ.SBjic7eOaV80QP9WnKpYJ2sAkzJhBD9QqtVikXHJ3SM
```
```json
{
    "code": 200
}
```