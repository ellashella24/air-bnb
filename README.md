# air-bnb
a rest-ful api project 

Run project with: 
```
go run main.go
```

## Stack-tech :dart:
- [x] RESTful API Using Go, Echo, Gorm, MySQL
- [x] AWS for service api

## Open Endpoints

Open endpoints require no Authentication.

* Register : `POST /users/register`
* Login : `POST /users/login/`

## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the request. A Token can be acquired from the Login view above.

### Current User related

Each endpoint manipulates or displays information related to the User whose Token is provided with the request:

- Get user profile data by User ID : `GET /users/profile`
- Update user data by User ID : `PUT /users`
- Delete user data by User ID : `DELETE /users`

### City related

Each endpoint manipulates or displays information related to the Homestay whose Token is provided with the request:

- Get all city data : `GET /city`
- Get city data by ID : `GET /city/:id`

### Homestay related

Each endpoint manipulates or displays information related to the Homestay whose Token is provided with the request:

- Get all homestay data : `GET /homestays`
- Get all homestay data by HostID : `GET /homestays/host`
- Get all homestay data by City : `GET /homestays/:search`
- Create homestay : `POST /homestay/create`
- Update homestay : `PUT /homestay/update`
- Delete homestay : `DELETE /homestay/delete`

### Booking related

Each endpoint manipulates or displays information related to the Booking whose Token is provided with the request:

- Create booking : `POST /booking`
- Get booking history data by UserID : `GET /history/:id`
- Get booking history data by HostID : `GET /recap/:id`
- Checkout booking : `PUT /booking/checkout`
- Reschedule booking : `PUT /booking/reschedule`


## Endpoints that require Check Role isAdmin
The endpoint below requires checking that the currently logged in user role is admin

### Current User related

Each endpoint manipulates or displays information related to the User whose Token is provided with the request:

- Get all user data: `GET /users`

### City related

Each endpoint manipulates or displays information related to the City whose Token and the role is admin that provided with the request:

- Create city : `POST /city/create`
- Update city : `PUT /city/update`
- Delete city : `DELETE /city/delete`

