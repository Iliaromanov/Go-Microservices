# gRPC_Docker-ChatApp

Simple chat app built using Docker and gRPC

## Docker Notes

## General Notes

- article on server side streaming using gRPC: https://itnext.io/streaming-data-with-grpc-2eb983fdee11 

  (this microservice uses streaming as can be seen in the `Broadcast` service defined in the .proto file)
  
  In brief: data streaming is a way of alerting the client when the server has new data available rather than the client constantly having to send requests to the server regardless of whether new data is available or not
