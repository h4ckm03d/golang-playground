### Auth:

- login
- check login
- valid -> fetch data dari db -> generate token -> return token
- invalid -> return 401

# Request secret/admin/user area

- request with jwt
- validate jwt -> data email, data user dll - context

# sample request

```bash
curl -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJjbGFzcyI6ImJjYyJ9.3qJKWhNYdwCLFCib8R0ey4BRQur6WBUkceeTcLUxRC0' localhost:5000/secret/hello
```
