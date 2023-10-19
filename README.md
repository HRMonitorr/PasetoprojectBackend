# Paseto Project Backend
## _Progress Status_
Last Edited 19 Oktober 2023 10:58 PM UTC+7

## Endpoint

- Insertdata User To Mongo
- Paseto Generator token
- Password Hasher
- Function PasetoEncode With Role
- Function Paseto Decode Get Username
- Function Paseto Decode Get Rolename
- Endpoint Sign Up
- Endpoint Login
- Endpoint Getdata With Login Header

## Insert Data User

API Link : 

```sh
https://asia-southeast2-gis-project-401902.cloudfunctions.net/UserdataInsert
```
Body :

```sh
{
    "username" : "user",
    "password" : "pass"
}
```

Password akan dihash otomatis didalam fungsi
Response :
```sh
{
    "status": true,
    "message": "Berhasil Input data"
}
```

## Paseto Token

API Link : 

```sh
https://asia-southeast2-gis-project-401902.cloudfunctions.net/PASETO-Project
```
Body :

```sh
{
    "username" : "user",
    "password" : "pass"
}
```
Response : 
```sh
{
    "status": true,
    "token": "token",
    "message": "Selamat Datang"
}
```

## Password Hasher

API Link : 

```sh
https://asia-southeast2-gis-project-401902.cloudfunctions.net/Password-Hasher
```
Body :

```sh
{
    "username" : "user",
    "password" : "pass"
}
```
Response : 
```sh
{
    "status": true,
    "token": "$2a$12$AMrX.qRPNEBUaa2HhSP8dOte0Fu9wO7vJ19IPxarXkhFQNNEAA3HW",
    "message": "Berhasil Hash Password"
}
```

## Register

API Link : 

```sh
https://asia-southeast2-gis-project-401902.cloudfunctions.net/RegisterUser
```
Body :

```sh
{
    "username" : "user",
    "password" : "pass",
    "role" : "role"
}
```
Response : 
```sh
{
    "status": true,
    "message": "Berhasil Input data"
}
```
## Login

API Link : 

```sh
https://asia-southeast2-gis-project-401902.cloudfunctions.net/Login
```
Body :

```sh
{
    "username" : "user",
    "password" : "pass",
    "role" : "role"
}
```
Response : 
```sh
{
    "status": false,
    "token": "token",
    "message": "Selamat Datang"
}
```

## GetUser With Token

API Link : 

```sh
https://asia-southeast2-gis-project-401902.cloudfunctions.net/GetUserWithToken
```
Header :

```sh
Login : "TokenString"
```
Response : 
```sh
{
    "status": true,
    "message": "data User berhasil diambil",
    "data": [
        {
            "username": "data",
            "password": "data",
            "role": "role"
        },
        {
            "username": "data",
            "password": "data",
            "role": "role"
        }
    ]
}
```

### Important
Before using GetUser With Token Login terlebih dahulu untuk dapat token string