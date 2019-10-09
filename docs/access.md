# Access
modul master akses

method | url | description
--- | --- | ---
[GET](#list) | [/access](#list) | Get All List Access

## GET /access
<a name="list"></a>
### Response
```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": [
        {
            "id": 1,
            "name": "root",
            "alias": "root"
        },
        {
            "id": 2,
            "parent_id": 1,
            "name": "health",
            "alias": "health"
        },
        {
            "id": 3,
            "parent_id": 2,
            "name": "GET /health",
            "alias": "check.Health"
        },
        {
            "id": 4,
            "parent_id": 1,
            "name": "users",
            "alias": "users"
        },
        {
            "id": 5,
            "parent_id": 4,
            "name": "GET /users",
            "alias": "users.List"
        },
        {
            "id": 6,
            "parent_id": 4,
            "name": "GET /users/{id}",
            "alias": "users.View"
        },
        {
            "id": 7,
            "parent_id": 4,
            "name": "POST /users",
            "alias": "users.Create"
        },
        {
            "id": 8,
            "parent_id": 4,
            "name": "PUT /users/{id}",
            "alias": "users.Update"
        },
        {
            "id": 9,
            "parent_id": 4,
            "name": "DELETE /users/{id}",
            "alias": "users.Delete"
        },
        {
            "id": 10,
            "parent_id": 1,
            "name": "roles",
            "alias": "roles"
        },
        {
            "id": 11,
            "parent_id": 10,
            "name": "GET /roles",
            "alias": "roles.List"
        },
        {
            "id": 12,
            "parent_id": 10,
            "name": "GET /roles/{id}",
            "alias": "roles.View"
        },
        {
            "id": 13,
            "parent_id": 10,
            "name": "POST /roles",
            "alias": "roles.Create"
        },
        {
            "id": 14,
            "parent_id": 10,
            "name": "PUT /roles/{id}",
            "alias": "roles.Update"
        },
        {
            "id": 15,
            "parent_id": 10,
            "name": "DELETE /roles/{id}",
            "alias": "roles.Delete"
        },
        {
            "id": 16,
            "parent_id": 10,
            "name": "POST /roles/{id}/access/{access_id}",
            "alias": "roles.Grant"
        },
        {
            "id": 17,
            "parent_id": 10,
            "name": "DELETE /roles/{id}/access/{access_id}",
            "alias": "roles.Revoke"
        },
        {
            "id": 18,
            "parent_id": 1,
            "name": "access",
            "alias": "access"
        },
        {
            "id": 19,
            "parent_id": 18,
            "name": "GET /access",
            "alias": "access.List"
        }
    ]
}
```