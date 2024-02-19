# Tern + SQLc

Proof of concept test for Tern + SQLc with PostgreSQL

### for creating user 
- `POST`
```sh
curl -X POST -H "Content-Type: application/json" -d '{                                                                                                                                                                       9GiB/15GiB
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com",
      "password": "securepassword"
  }' http://localhost:8080/users
```
