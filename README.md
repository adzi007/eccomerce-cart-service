# Ecommerce Cart Service
A microservice responsible for managing user cart data in an e-commerce system. Built with Go and Fiber using Clean Architecture principles.

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)


---

## 🚀 Features

- Add products to cart
- Update product quantity in cart
- Remove products or clear cart
- Retrieve current cart data per user
- Communicates with Product Service (via REST) to validate product data
- Caches frequent cart reads using Redis

---

## 🧰 Tech Stack

- **Language**: Golang (Go 1.21+)
- **Framework**: [Fiber](https://gofiber.io/)
- **Architecture**: Clean Architecture (Handler → Usecase → Repository)
- **Database**: MySQL
- **Cache**: Redis
- **Communication**: REST HTTP (to Product Service) and gRPC server to receive request from order service
- **Containerization**: Docker + Docker Compose

---

## 📁 Project Structure (Clean Architecture)
```md
├───cart_proto
├───config
│   └───database
├───docs
├───internal
│   ├───domain
│   ├───handler
│   │   └───grpc_handler
│   ├───migration
│   ├───model
│   │   ├───entity
│   │   ├───request
│   │   └───response
│   ├───repository
│   │   ├───cache
│   │   └───product_service_repo
│   └───usecase
├───pkg
│   ├───cachestore
│   ├───logger
│   └───utils
├───redis.conf
├───server
└───tmp
```
## Getting Started

### Prerequisites
- Docker
- Go 1.21+
- <a href="https://github.com/adzi007/ecommerce-products-service" target="_blank">Ecommerce Product Service</a>

## Running Locally (Docker)

1. Clone the project
```bash
git clone https://github.com/adzi007/eccomerce-cart-service.git
cd ecommerce-cart-service
```
2. CD into the ecommerce-cart-service directory and create an .env file or edit from .env.example following with fields bellow
```
PORT_APP=5000
DB_HOST=ecommerce-mysql-db
DB_PORT=3306
DB_USERNAME=YOUR_DB_USERNAME
DB_PASSWORD=YOUR_DB_PASSWORD
DB_NAME=ecommerce_app
URL_PRODUCT_SERVICE=http://ecommerce-products-service:3000/products
REDIS_HOST=ecommerce-cart-redis
REDIS_PORT=6379
```

3. Build container
```
docker-compose up --build
```

The App will be running at `http://localhost:5000`

## Database Migration

1. Execute migration database
```
docker exec -it ecommerce-cart-service /migrate
```
## API Documentation
<a href="https://www.postman.com/grey-satellite-91338/ms-ecommerce-projects/overview" target="_blank">Postman Collections</a>

