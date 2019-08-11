D=docker
BUF_DIR = service
CL_DIR = greeter_client
SRV_DIR = greeter_server
STACK = scale
all:
	make clean
	make protoc
	make .server_image
	make .client_image
	make .swarm_init
	make stack_deploy
clean:
	sudo docker service rm $(STACK)_app $(STACK)_client $(STACK)_proxy
	rm -r $(BUF_DIR)/service.pb.go $(SRV_DIR)/service.pb.go $(CL_DIR)/service.pb.go
protoc:
	protoc --go_out=plugins=grpc:. $(BUF_DIR)/service.proto
	cp $(BUF_DIR)/service.pb.go $(CL_DIR)/
	cp $(BUF_DIR)/service.pb.go $(CL_DIR)/
.swarm_init:
	docker swarm init
stack_deploy:
	sudo docker stack deploy --compose-file=docker-compose.yml $(STACK)
	sleep 10
.server_image:
	sudo $(D) build -t app_service:latest $(SRV_DIR)/
.client_image:
	sudo $(D) build -t app_service:latest $(CL_DIR)/
