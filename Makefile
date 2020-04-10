.PHONY: build clean

clean:
	rm -rf ./functions

build:
	mkdir -p functions
	# Users Endpoints
	# getonebyauthid
	go get ./srv/endpoints/users/getone/...
	go build -o functions/users-getone ./srv/endpoints/users/getone/...
	#Weights Endpoints
	# getallbyuserid
	go get ./srv/endpoints/weights/getall/...
	go build -o functions/weights-getall ./srv/endpoints/weights/getall/...
