# book-grpc-users

Install protoc complier in the machine from the link below for windows machine:
https://github.com/protocolbuffers/protobuf/releases/tag/v28.3

Then while coding a grpc project, start with .proto file then code the file with proto buf specification. Compile them using the protoc compiler and generate two files using below command.

protoc --go_out=../chat --go_opt=paths=source_relative --go-grpc_out=../chat --go-grpc_opt=paths=source_relative users.proto 

Include the option in the proto file as it is required.