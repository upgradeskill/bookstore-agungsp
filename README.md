# Task 3 - Majoo Golang Bootcamp

<hr>

## Build repo to executable file

`goexec build`
<small style="color: red;"><i>\*This command only for Windows users</i></small>

## Build and start server

`goexec run`
<small style="color: red;"><i>\*This command only for Windows users</i></small>

## Login
|Method|Endpoint|
|---|---|
|<b style="color: blue;">POST</b>| `/login` |
> <b style="color: green;">Response</b>
> `{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"}`
>&nbsp;

**Use the token on bearer authentication**

## Endpoint list
|Method|Endpoint|Desctiption|Status|
|---|---|---|---|
| <b style="color: blue;">GET</b>  | `/books`  | Get all data  | <span style="color: green;">Accessible</span>  |
| <b style="color: blue;">GET</b>  | `/books/:id`  | Get specific data with id  | <span style="color: green;">Accessible</span>  |
| <b style="color: blue;">POST</b>  | `/books`  | Create data  | <span style="color: red;">Restricted</span>  |
| <b style="color: blue;">PUT</b>  | `/books/:id`  | Update data  | <span style="color: red;">Restricted</span>  |
| <b style="color: blue;">DELETE</b>  | `/books/:id`  | Delete data  | <span style="color: red;">Restricted</span>  |
