This repo contains a web app to store and manage important dates such as birthdays or anniversaries. The app has two modules: category and event. 
You can create, update and delete a category. When creating an event you will need to give a name, date, select a category and add details if you want.
You can read, craete, edit and delete an event.
You can also search for an event by the name.  
gRPC is used to communicate with the services.
Please change the environment variables of the server and client accordingly in the config file.  
Find the config file of the server in **todo/env**.  
Find the config file of the client in **cms/env**.  
To run the DB migration:  
```
cd todo
go run migrations/migrate.go up
```
To run the server: 
```
go run todo/main.go
```  
To run the client: 
```
go run cms/main.go
```
To view the web app in the browser go to: localhost:8080 (or the port number you specified in the config file).  
