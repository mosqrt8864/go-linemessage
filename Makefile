.PHONY: mongo-build mongo-clean app-build app-clean

#mongo 
mongo-build:
	sudo docker build -t mongodbimg ./build/package/docker/mongoDB/
	sudo docker run -d --publish 27017:27017 --name mongodb mongodbimg
mongo-clean:
	sudo docker stop mongodb
	sudo docker rm mongodb

#linemessage application
app-build:
	sudo docker build -t appimg -f build/package/docker/linemessage/Dockerfile .
	sudo docker run -d --link mongodb:mongodb --publish 8080:8080 --name app appimg
app-clean:
	sudo docker stop app
	sudo docker rm app