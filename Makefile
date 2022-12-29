.PHONY: mongo-build mongo-clean

#mongo 
mongo-build:
	sudo docker build -t mongodbimg ./build/package/docker/mongoDB/
	sudo docker run -d --publish 27017:27017 --name mongodb mongodbimg
mongo-clean:
	sudo docker stop mongodb
	sudo docker rm mongodb