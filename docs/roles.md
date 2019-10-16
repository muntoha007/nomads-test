# Roles
Modul Master Roles

method | url | description
--- | --- | ---
[GET](#list) | [/roles](#list) | Get All List Roles
[POST](#new) | [/roles](#new) | Post New Roles
[GET](#list_id) | [/roles/id](#list_id) | Get List Roles by ID
[PUT](#update_id) | [/roles/id](#update_id) | Update Roles by ID
[DELETE](#delete_id) | [/roles/id](#delete_id) | Delete Roles by ID

# Role Access
Modul Roles Access
method | url | description
--- | --- | ---
[POST](#new_role_access) | [/roles/id/access/id](#new_role_access) | Create Roles Access by ID Role and ID Role Access
[DELETE](#delete_role_access) | [/roles/id/access/id](#delete_role_access) | Delete Roles Access by ID Role and ID Role Access

## GET /roles
<a name="list"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": [
        {
            "id": 1,
            "name": "superadmin"
        }
    ]
}
```

## POST /roles
<a name="new"></a>
### Body

```
{
    "name": "user"
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 2,
        "name": "user"
    }
}
```

## GET /roles/id
<a name="list_id"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 1,
        "name": "superadmin"
    }
}
```

## PUT /roles/id
<a name="update_id"></a>
### Body
```
{
	"id" : 2,
    "name": "news-admin"
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 2,
        "name": "news-admin"
    }
}
```

## DELETE /roles/id
<a name="delete_id"></a>
### Response

```

```

## POST /roles/id/access/id
<a name="new_role_access"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": null
}
```

## DELETE /roles/id/access/id
<a name="delete_role_access"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": null
}
```