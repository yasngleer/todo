You can access the api from:
https://todo-production-8c32.up.railway.app/
admin e-mail: admin@admin.com
admin password: admin
user e-mail: example@example.com
user password: example

### Installation 
#### With Docker
Clone this repo and enter the directory

``` 
docker build -t "todoapp" .
docker run -p 8081:80 todoapp
``` 

You can use any rest client to access the api from localhost:8081

### TODO
- Complete missing error handling
- Write tests
- Use swagger
- Make todostore interface
- Write mock todostore instance
- Use validator for request


### Endpoints and Usage
           
- POST   /api/users           
    
    Create-User
    ``` 
    {
	"email":"client1@gmail.com",
	"password":"password"
    }
    ``` 
- POST   /api/users/login
    
    Login
    ``` 
    {
	"email":"client1@gmail.com",
	"password":"password"
    }
    ```
- POST   /api/todo

    Add Todo
    ``` 
    {
	"Name":"Todo Name"
    }
    ```
- GET    /api/todo/:id        
    
    Get Todo by ID

- GET    /api/todo/

    Get all todos

- PUT    /api/todo/:todoid    

    Update todo by id
    ``` 
    {
	"Name":"NewTodoName"
    }
    ```
- DELETE /api/todo/:todoid    
    
    Delete todo by id

- POST   /api/todo/:id/step   
    
    Add new step to todo
    ``` 
    {
	"Context":"Step 4"
    }
    ```
- PUT    /api/todo/:todoid/step/:stepid 
    
    Update step
    ``` 
    {
	"Context":"Step 4",
    "Completed":true
    }
    ```
- DELETE /api/todo/:todoid/step/:stepid 
    
    Delete todo 