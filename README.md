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
