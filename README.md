# go-gin-api-boilerplate
This is a boilerplate base repository for a Gin Golang API with Swagger documentation and MariaDB MySQL, including all the required packages and necessary helpers for a complete backend project using Gin. 

##  Local Machine Setup (Linux)
Follow these steps to set up the project on your local machine:
- Clone project from `git clone https://github.com/lushiv/go-gin-api-boilerplate.git`
- Go the `cd go-gin-api-boilerplate`
- Install all Dependencies using `go get ./...`
- After that Make the script executable using `chmod +x start_server.sh`
- After that make `.env` file from `.env.sample` and change file basis on your configurations
- Now Run your server using the script: `./start_server.sh`
- Now application is running on : [http://localhost:3000/docs/index.html#/](http://localhost:3000/docs/index.html#/)

### For Database Migration
 - Go to the dir `cd go-gin-api-boilerplate/db/db_migration/go_gin_api_boilerplate_db`
 - Next open the file `liquibase.properties`
 - Now change these things basis on your creation:
 -  <span style="color:red">YOUR_DATABASE_NAME</span>
 -  <span style="color:red">YOUR_DB_USERNAME</span>
 -  <span style="color:red">YOUR_DB_PASSWORD</span>

```
changeLogFile: go_gin_api_boilerplate_db/changelog-master.xml
driver: com.mysql.cj.jdbc.Driver
url: jdbc:mysql://localhost:3306/YOUR_DATABASE_NAME?autoReconnect=true&useSSL=false&maxReconnects=10&allowPublicKeyRetrieval=true&createDatabaseIfNotExist=true
username: YOUR_DB_USERNAME
password: YOUR_DB_PASSWORD
logLevel=DEBUG
classpath: ./mysql-connector-java-8.0.19.jar

```
- Now that all database configurations have been updated, you can proceed to run `./db_migration.sh`
