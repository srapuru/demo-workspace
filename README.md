# demo-workspace
# Please follow below steps to use the repo


1) clone the repo using the following command, It will clone into the directory  'demo-workspace'

git clone https://github.com/srapuru/demo-workspace.git    

2) cd demo-workspace/src

3) Look at the protobuf file demo-workspace/src/proto/Sample.proto for the sample GRPC call;
It has one grpc call   Hello ( input as (HelloRequest and reponse as HelloResponse )

4) make clean;make 
 --generates  protobuf stubs in the directory called 'generated' 

5) demo-workspace/bin directory contains the executables 
server --- GRPC server 
client ----GRPC client
server listens on the port localhost:9999 for the requests 

6) Added the following Dockerfiles 
	./src/client/Dockerfile - to create a docker server image 
	./src/server/Dockerfile - to create a docker client image 


7) cd demo-workspace; execute the following command 
	docker-compose up -d

Output :
Building with native build. Learn about native build in Compose here: https://docs.docker.com/go/compose-native-build/
Creating network "demo-workspace_default" with the default driver
Creating server_container ... done
Creating client_container ... done

=== docker ps # command shows as follows, it means server still up and running, client executes and stopped 

 docker ps
CONTAINER ID        IMAGE               COMMAND                 CREATED             STATUS              PORTS                    NAMES
0f592959b211        server:1.0          "/bin/sh -c ./server"   31 seconds ago      Up 29 seconds       0.0.0.0:9999->9999/tcp   server_container


===== docker exec -it 0f592959b211 bash # to enter into the server image and see the logs 
root@0f592959b211:/opt# cat server.log
2021/02/06 23:58:35 ==== Starting the server
2021/02/06 23:58:36 Received: GRPC

== docker-compose  down
Stopping server_container ... done
Removing client_container ... done
Removing server_container ... done
Removing network demo-workspace_default




