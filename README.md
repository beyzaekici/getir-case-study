## Getir Case

### _About_

This project is a web api. It has 3 endpoints.
These endpoints fetch the data in the provided MongoDB collection and return the result in the requested format.

##### _Requirements_

go 1.17

##### _Installation_

Clone this repo to your local machine using https://github.com/beyzaekici/getir-case-study

##### _Endpoints_

/in-memory 
/in-memory/ 
/records

##### **_Mongo Search_**

**URL** :   /records

**Method** : POST

**Request example**

```json
{
"startDate": "2016-01-21",
"endDate": "2016-03-02",
"minCount": 2900,
"maxCount": 3000
}
```

**Success Response**

**Code** : 200

**Response example**

```json
{
"code": 0,
"msg": "Success",
"records":
        [
                {
                    "createdAt": "2016-02-19T08:35:39.409+02:00",
                    "key": "kkxEdhft",
                    "totalCount": 2980
                }
        ]
}
```

##### _Set Data to InMemory Cache_

**URL** : /in-memory/

**Method** : POST

**Request example**

```json
{
"key": "example",
"value": "example"
}
```

**Success Response**

**Code** : 201

**Response example**

```json
{
"key": "example",
"value": "example"
}
```

##### _Get Data from InMemory Cache_

**Request example**
**URL** : /in-memory?key=example

**Method** : GET

**Success Response**

**Code** : 200

**Response example**

```json
{
"key": "example",
"value": "example"
}
```
