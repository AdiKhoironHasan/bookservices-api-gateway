# Microservices App (gRPC + REST)

![alt text](https://github.com/AdiKhoironHasan/bookservices-api-gateway/blob/main/var/app/image/system-design.png?raw=true)

<!-- ABOUT THE PROJECT -->
## About The Project

The Book Services project is a set of microservices that enables users to manage books and users data. It consists of three main services: API Gateway Services, Book Services, and User Services. The communication between these services is done through gRPC, using the ProtoBank contract. The gRPC protocol implements rtoken to ensure the security of services against unauthorized access.

External communication is done via REST API, which is processed through API Gateway Services, and then forwarded to the target service using gRPC. JWT is used to ensure the security of client-to-service communication.

Each service has its own PostgreSQL database. GORM is used to access these databases using ORM (Object Relational Mapping).

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

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
5. Generate Go code from the proto file.
```
make run-proto
```
6. To start the service with REST API.
```
go run main.go
```
7. To start the service with gRPC.
```
run go run main.go grpc:start
 ```

<!-- USAGE EXAMPLES -->
## Usage

To make it easier for you to try the endpoint service, we have provided a "Run in Postman" button which you can click to directly open Postman with the configured environment and collection.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/18402968-9b1f0b22-ebeb-481f-b205-022558c4f089?action=collection%2Ffork&collection-url=entityId%3D18402968-9b1f0b22-ebeb-481f-b205-022558c4f089%26entityType%3Dcollection%26workspaceId%3Da3c53e94-cbc8-4668-9027-b2122261f411)

By using Postman, you can easily execute endpoint services and see the response results in a format that is easy to read and understand. You can also test the various types of requests supported by the endpoint service, such as GET, POST, PUT, DELETE, etc.

For more information on how to use Postman, you can visit the official documentation page at https://learning.postman.com/docs/getting-started/introduction/.

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

<!-- CONTACT -->
## Contact

[Adi Khoiron Hasan](https://www.linkedin.com/in/adi-khoiron-hasan) - adikhoironhasan@gmail.com

### Project Link
- service-api-gateway [https://github.com/AdiKhoironHasan/bookservices-api-gateway](https://github.com/AdiKhoironHasan/bookservices-api-gateway)
- service-book [https://github.com/AdiKhoironHasan/bookservices-users](https://github.com/AdiKhoironHasan/bookservices-users)
- service-user [https://github.com/AdiKhoironHasan/bookservices-books](https://github.com/AdiKhoironHasan/bookservices-books)
- protobank [https://github.com/AdiKhoironHasan/bookservices-protobank](https://github.com/AdiKhoironHasan/bookservices-protobank)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
