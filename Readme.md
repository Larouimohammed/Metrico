#run docker image for mongo data base
docker run --name mongodb -d mongo:latest -p 27017:27017
#install go client driver mongodb library for go
go get go.mongodb.org/mongo-driver/mongo
#install gin library for go
go get github.com/gin-gonic/gin
#runnig Metriko app{Server+API+Agent}
make run