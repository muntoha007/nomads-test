# User Login
Modul Master Users Login

method | url | description
--- | --- | ---
[POST](#login) | [/users](#login) | For Login Post Username and Password

## POST /login
<a name="login"></a>
### Body

```
{
    "username": "jackyhtg",
    "password": "12345678"
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphY2t5aHRnIiwiZXhwIjoxNTcxMDQ0ODgwfQ.LBElJxejpBVgwsw-VcE-pb1vi11P2FNPaKs2Z6255u8"
    }
}
```

