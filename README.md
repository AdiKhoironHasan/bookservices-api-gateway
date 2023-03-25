# Microservices App (gRPC + REST)

![alt text](https://github.com/AdiKhoironHasan/bookservices-api-gateway/blob/main/var/app/image/system-design.png?raw=true)

<!-- ABOUT THE PROJECT -->
## About The Project

The BookServices project is a set of microservices that enables users to manage books and user data. The project consists of three main services: bookservices-api-gateway, bookservices-books, and bookservices-users. In the BookServices project, Coreography pattern is used as part of the [Saga Pattern](https://microservices.io/patterns/data/saga.html). In the project, bookservices-api-gateway acts as the central communication and control point, forwarding external client requests and responses to internal microservices. Coreography pattern is used to ensure that each microservice can handle its own transactional part independently, while bookservices-api-gateway coordinates and orchestrates them to achieve the overall goal.

Internal communication between services is done through gRPC, using the ProtoBank contract from bookservices-protobank. The gRPC protocol implements tokens to ensure the security of services against unauthorized access. External communication is done via REST API, which is processed through bookservices-api-gateway, and then forwarded to microservices using gRPC. JWT is used to ensure the security of external communication.

Each service has its own database using [PostgreSQL](https://www.postgresql.org/docs), and [GORM](https://gorm.io/docs) is used to access these databases using Object Relational Mapping (ORM).

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* PosgreSQL Database
* Go 1.20 or newer
* Go protobuf compiler (protoc)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Installation
1. Begin by cloning the repository to your local machine.
2. Set up a PostgreSQL database and create the necessary tables.
3. Create an environment file by duplicating the .env.example file and filling in the required values for each service according to your needs.
4. Install the necessary dependencies.
```
go get all
```
5. Migrate database.
```
go run main.go db:migrate
```
6. Generate Go code from the proto file.
```
make run-proto
```
7. To start the service with REST API.
```
go run main.go
```
8. To start the service with gRPC.
```
run go run main.go grpc:start
 ```

<!-- USAGE EXAMPLES -->
## Usage

To make it easier for you to try the endpoint service, please use the "Run in Postman" button which you can click to directly open Postman with the configured environment and collections.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/18402968-9b1f0b22-ebeb-481f-b205-022558c4f089?action=collection%2Ffork&collection-url=entityId%3D18402968-9b1f0b22-ebeb-481f-b205-022558c4f089%26entityType%3Dcollection%26workspaceId%3Da3c53e94-cbc8-4668-9027-b2122261f411)

By using [Postman](https://learning.postman.com/docs/getting-started/introduction/), you can easily execute endpoint services and see the response results in a format that is easy to read and understand. You can also test the various types of requests supported by the endpoint service, such as GET, POST, PUT, DELETE, etc.

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "ENCHANCHEMENT".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- THANK YOU -->
## Thank You
Many Thanks to [Ach Jailani](https://github.com/achjailani) who let me use his grpc boilerplate for this project.

<!-- CONTACT -->
## Contact

[Adi Khoiron Hasan](https://www.linkedin.com/in/adi-khoiron-hasan) - adikhoironhasan@gmail.com

### Project Link
- go-simple-grpc [https://github.com/achjailani/go-simple-grpc](https://github.com/achjailani/go-simple-grpc)
- service-api-gateway [https://github.com/AdiKhoironHasan/bookservices-api-gateway](https://github.com/AdiKhoironHasan/bookservices-api-gateway)
- service-book [https://github.com/AdiKhoironHasan/bookservices-users](https://github.com/AdiKhoironHasan/bookservices-users)
- service-user [https://github.com/AdiKhoironHasan/bookservices-books](https://github.com/AdiKhoironHasan/bookservices-books)
- protobank [https://github.com/AdiKhoironHasan/bookservices-protobank](https://github.com/AdiKhoironHasan/bookservices-protobank)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
