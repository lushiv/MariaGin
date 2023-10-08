
<a href="#"><img width="20%" height="auto" src="https://i.ibb.co/bXhyRw7/Screenshot-from-2023-10-02-17-53-14.png" height="100"/></a>
## Swaggerize Go Gin Backend With MariaDB
MariaGin is your comprehensive starting point for developing robust and feature-rich backend applications with the Go programming language, powered by the Gin web framework. This boilerplate repository comes fully equipped with Swagger documentation and MariaDB MySQL integration, offering a streamlined experience for building scalable and well-documented APIs.
## Key Features:
![Go](https://img.shields.io/badge/-Go-05122A?style=flat&logo=go)&nbsp;
![Gin](https://img.shields.io/badge/-Gin-05122A?style=flat&logo=gin)&nbsp;
![MariaDB](https://img.shields.io/badge/-MariaDB-05122A?style=flat&logo=mariadb)&nbsp;
![Swagger](https://img.shields.io/badge/-Swagger-05122A?style=flat&logo=swagger)&nbsp;
![JWT](https://img.shields.io/badge/-JWT-05122A?style=flat&logo=jwt)&nbsp;
![Redis](https://img.shields.io/badge/-Redis-05122A?style=flat&logo=Redis)&nbsp;
![Liquibase](https://img.shields.io/badge/-Liquibase-05122A?style=flat&logo=liquibase)&nbsp;
![rabbitmq](https://img.shields.io/badge/-RabbitMq-05122A?style=flat&logo=rabbitmq)&nbsp;
![SendGrid](https://img.shields.io/badge/-SendGrid-05122A?style=flat&logo=e)&nbsp;
![Firebase](https://img.shields.io/badge/-Firebase-05122A?style=flat&logo=firebase)&nbsp;
![AWS_S3](https://img.shields.io/badge/-AWS_S3-05122A?style=flat&logo=AmazonS3)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;
![Script](https://img.shields.io/badge/-Script.sh-05122A?style=flat&logo=gnu-bash)&nbsp;

##  Local Machine Setup (Linux)
Follow these steps to set up the project on your local machine:
- Clone project from `git clone https://github.com/lushiv/MariaGin`
- Go the `cd MariaGin`
- Install all Dependencies using `go get ./...`
- After that Make the script executable using `chmod +x start_server.sh`
- After that make `.env` file from `.env.sample` and change file basis on your configurations
- Set up database migration by following the database migration steps
- Now Run your server using the script: `./start_server.sh`
- Now application is running on : [http://localhost:3000/docs/index.html#/](http://localhost:3000/docs/index.html#/)

### For Database Migration
 - Go to the dir `cd MariaGin/db/db_migration/maria_gin_db`
 - Next open the file `liquibase.properties`
 - Now change these things basis on your creation:
 -  <span style="color:red">YOUR_DATABASE_NAME</span>
 -  <span style="color:red">YOUR_DB_USERNAME</span>
 -  <span style="color:red">YOUR_DB_PASSWORD</span>

```
changeLogFile: maria_gin_db/changelog-master.xml
driver: com.mysql.cj.jdbc.Driver
url: jdbc:mysql://localhost:3306/YOUR_DATABASE_NAME?autoReconnect=true&useSSL=false&maxReconnects=10&allowPublicKeyRetrieval=true&createDatabaseIfNotExist=true
username: YOUR_DB_USERNAME
password: YOUR_DB_PASSWORD
logLevel=DEBUG
classpath: ./mysql-connector-java-8.0.19.jar

```
- Now that all database configurations have been updated, you can proceed to run `./db_migration.sh`
