# Vehicle
Modul Master Vehicle

method | url | description
--- | --- | ---
[GET](#list) | [/vehicles](#list) | Get All List Vehicles
[POST](#new) | [/vehicles](#new) | Post New Vehicles
[GET](#list_id) | [/vehicles/id](#list_id) | Get List Vehicles by ID
[PUT](#update_id) | [/vehicles/id](#update_id) | Update Vehicles by ID
[DELETE](#delete_id) | [/vehicles/id](#delete_id) | Delete Vehicles by ID

## GET /vehicles
<a name="list"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": [
        {
            "id": 1,
            "brand_id": 1,
            "vehicle_name": "Vario 125x"
        }
    ]
}
```

## POST /vehicles
<a name="new"></a>
### Body

```
{
	"brand_id": 1,
    "vehicle_name": "Supra X"
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 4,
        "brand_id": 1,
        "vehicle_name": "Supra X"
    }
}
```

## GET /vehicles/id
<a name="list_id"></a>
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 4,
        "brand_id": 1,
        "vehicle_name": "Supra X"
    }
}
```

## PUT /vehicles/id
<a name="update_id"></a>
### Body
```
{
	"id" : 4,
	"brand_id" : 1,
    "vehicle_name": "Supra X 125cc"
}
```
### Response

```
{
    "status_code": "REBEL-200",
    "status_message": "OK",
    "data": {
        "id": 4,
        "brand_id": 1,
        "vehicle_name": "Supra X 125cc"
    }
}
```

## DELETE /vehicles/id
<a name="delete_id"></a>
### Response

```

```