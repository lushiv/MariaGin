# MariaGin-Swaggerize
This is a boilerplate base repository for a Gin Golang API with Swagger documentation and MariaDB MySQL, including all the required packages and necessary helpers for a complete backend project using Gin. 

## Implemented Stacks

![Go](https://img.shields.io/badge/-Go-05122A?style=flat&logo=go)&nbsp;
![Gin](https://img.shields.io/badge/-Gin-05122A?style=flat&logo=gin)&nbsp;
![MariaDB](https://img.shields.io/badge/-MariaDB-05122A?style=flat&logo=mariadb)&nbsp;
![JWT](https://img.shields.io/badge/-JWT-05122A?style=flat&logo=jwt)&nbsp;
![SendGrid](https://img.shields.io/badge/-SendGrid-05122A?style=flat&logo=sendgrid)&nbsp;
![Swagger](https://img.shields.io/badge/-Swagger-05122A?style=flat&logo=swagger)&nbsp;
![Git](https://img.shields.io/badge/-Git-05122A?style=flat&logo=git)&nbsp;
![Liquibase](https://img.shields.io/badge/-Liquibase-05122A?style=flat&logo=liquibase)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;
![Script](https://img.shields.io/badge/-Script.sh-05122A?style=flat&logo=gnu-bash)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![MySql](https://img.shields.io/badge/-MySql-05122A?style=flat&logo=MySql)&nbsp;
![Redis](https://img.shields.io/badge/-Redis-05122A?style=flat&logo=Redis)&nbsp;
![rabbitmq](https://img.shields.io/badge/-rabbitmq-05122A?style=flat&logo=rabbitmq)&nbsp;
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
