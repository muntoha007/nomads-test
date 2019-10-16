# Users
Modul Master Users

method | url | description
--- | --- | ---
[GET](#list) | [/users](#list) | Get All List Users
[POST](#new) | [/users](#new) | Post New Users
[GET](#list_id) | [/users/id](#list_id) | Get List Users by ID
[PUT](#update_id) | [/users/id](#update_id) | Update Users by ID
[DELETE](#delete_id) | [/users/id](#delete_id) | Delete Users by ID

## GET /users
<a name="list"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": [
        {
            "id": 1,
            "username": "jackyhtg",
            "is_active": true,
            "roles": [
                {
                    "ID": 1,
                    "Name": "superadmin"
                }
            ]
        }
    ]
}
```

## POST /users
<a name="new"></a>
### Body

```
{
    "username": "peterpan",
    "email": "peterpan@gmail.com",
    "password": "1234",
    "re_password": "1234",
    "is_active": true,
    "roles" : [
    	{
    		"id" : 1
    	}	
    ]
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 2,
        "username": "peterpan",
        "is_active": true,
        "roles": [
            {
                "ID": 1,
                "Name": ""
            }
        ]
    }
}
```

## GET /users/id
<a name="list_id"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 1,
        "username": "jackyhtg",
        "is_active": true,
        "roles": null
    }
}
```

## PUT /users/id
<a name="update_id"></a>
### Body
```
{
    "id": 2,
    "username": "gatholoco",
    "email": "gatholoco@gmail.com",
    "password": "1234",
    "re_password": "1234",
    "is_active": false,
    "roles" : [
    	{
    		"id": 1
    	}
    ]
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 2,
        "username": "gatholoco",
        "is_active": false,
        "roles": [
            {
                "ID": 1,
                "Name": ""
            }
        ]
    }
}
```

## DELETE /users/id
<a name="delete_id"></a>
### Response

```

```