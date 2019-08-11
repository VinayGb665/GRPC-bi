# How to use!
To jump into one of the client nodes and start sending messages run :
	
    sudo sh ./run_end_to_end 
This does a make in the background pulling a lot of data/docker-images so wait for the make to complete
**To build locally :** 

 
 1. Install go plugins for GRPC and protobuf
				 `go get -u google.golang.org/grpc `
Protogen-go installation instructions [here](https://grpc.io/docs/quickstart/go/)
	
 2. Compile protobuf in the service directory which generates equivalent go code  `protoc --go_out`
 3. Build the necessary docker images at client and server
	  `docker image build -t app_service greeter_server/`
	 `docker image build -t app_service greeter_server/`
 4. Init the swarm and deploy the stack with a suitable name
	  `docker stack deploy --compose-file=docker-compose.yml scale ` 
 5. Attach to one of the client containers to send messages
	  `docker ps #Find out the container id`
	  `docker attach cid # attached to tty `

Done

    

# Prerequisites

 1. Docker, Docker-compose
 2. Go, Make, git 
