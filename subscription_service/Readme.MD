Subscription Service Documentation
Overview
The subscription_service is a microservice responsible for managing subscriptions. It provides functionalities to create, retrieve, update, delete, and list subscriptions via gRPC.

Prerequisites
Go 1.18+
Docker
Docker Compose
Setup
Clone the Repository

Setup
Clone the Repository
git clone https://github.com/OkoroKelvin/Product-Microservice-Lithium/tree/main/subscription_service
cd subscription_service


Environment Variables
Create a .env file in the subscription_service directory with the following content:

DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=subscriptions
DB_PORT=5432
GRPC_PORT=:8081
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=subscriptions

Build and Run with Docker Compose

docker-compose up --build


Usage
gRPC Endpoints
The subscription_service exposes the following gRPC endpoints:

CreateSubscription: Creates a new subscription.
GetSubscription: Retrieves a subscription by its ID.
UpdateSubscription: Updates an existing subscription.
DeleteSubscription: Deletes a subscription by its ID.
ListSubscriptions: Lists all subscriptions with pagination.


Development
Generate gRPC Code
To generate gRPC code, run the following command:
make generate


Generate Mocks
To generate mocks for testing, run the following command:
make mocks