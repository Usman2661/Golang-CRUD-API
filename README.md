# Golang-CRUD-API
- [Todo Crud API using Gin,Gorm and PostgreSQL Database]

# Instructions to run 

- [1.Install golang from the golangs official website]

2. Clone the project in go directory i.e. go/src/github.com/{username}
3. Run command - go get to get all the dependances
4. For Hot Reload make sure Gin is installed using go get github.com/codegangsta/gin. Then copy the gin.exe file from User/go/bin into the C:/Go/bin folder
5. Create a file named .env in the root of the project with variables host,dbport,user,dbname,passsword,sslmode=disable,CorsUrl etc
6. Finally run the hot reload server using gin --appPort 8080 --all -i run main.go