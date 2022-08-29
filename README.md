### A library management application with Golang, Gin, Jwt and MongoDb
### CRUD Restful API (backend server) with hexagonal architecture

&nbsp;
### This application allows a library_user to see collection of books , rent a book
### while an admin can create, modify and delete a book from the store 
###

&nbsp;

#### Set your .env variables or use the default variables. 
#### set up mongodb on your computer,
#### run the application with 
```GO 

make run
 ```
 #### Access Different endpoints 
 ```GO
 POST "/admin/addbook"
GET "/admin/getbook/:book_id"
PATCH "/admin/editbook/:book_id"
GET "/user/getallbooks"
GET "/"
DELETE "/admin/deletebook/:_id"
POST "/user/signup"
POST "/user/signin" 
```

#### and other endpoints you can find in 
```GO
router/http.go
```