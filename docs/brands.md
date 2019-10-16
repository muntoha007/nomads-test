# Brands
Modul Master Brands

method | url | description
--- | --- | ---
[GET](#list) | [/brands](#list) | Get All List Brands
[POST](#new) | [/brands](#new) | Post New Brands
[GET](#list_id) | [/brands/id](#list_id) | Get List Brands by ID
[PUT](#update_id) | [/brands/id](#update_id) | Update Brands by ID
[DELETE](#delete_id) | [/brands/id](#delete_id) | Delete Brands by ID

## GET /brands
<a name="list"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": [
        {
            "id": 1,
            "name": "Honda"
        }
    ]
}
```

## POST /brands
<a name="new"></a>
### Body

```
{
    "name": "Suzuki"
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 2,
        "name": "Suzuki"
    }
}
```

## GET /brands/id
<a name="list_id"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 2,
        "name": "Suzuki"
    }
}
```

## PUT /brands/id
<a name="update_id"></a>
### Body
```
{
	"id" : 2,
    "name": "Suzuku"
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 2,
        "name": "Suzuku"
    }
}
```

## DELETE /brands/id
<a name="delete_id"></a>
### Response

```

```