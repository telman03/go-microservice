# Go Microservices Project

This project is a microservices-based architecture built using **Golang, Fiber, Kafka, PostgreSQL, and Docker**. It consists of multiple services that communicate asynchronously via Kafka, including a **User Service** and a **Notification Service**.

## 📌 Features
- **User Service:** Manages user accounts and provides user details via REST API.
- **Notification Service:** Listens for Kafka messages and sends notifications via email.
- **Kafka Integration:** Uses Kafka for asynchronous messaging between microservices.
- **PostgreSQL Database:** Stores user and notification data.
- **Dockerized Setup:** Services run in Docker containers for easy deployment.

## 🏗️ Microservices Overview

### 1️⃣ **User Service** (`user-service`)
- Manages user accounts.
- Exposes API endpoints to retrieve user details.
- Built with **Fiber** (Go Web Framework).
- Stores user data in **PostgreSQL**.

### 2️⃣ **Notification Service** (`notification-service`)
- Listens for order-related events from Kafka.
- Fetches user details from the **User Service** via REST API.
- Stores notifications in the database.
- Sends email notifications using `gomail`.

## 🚀 Tech Stack
- **Backend:** Golang with Fiber
- **Message Broker:** Apache Kafka
- **Database:** PostgreSQL
- **Docker:** Containerized microservices
- **Environment Management:** `.env` files

## 📦 Installation & Setup

### 1️⃣ **Clone the repository**
```sh
git clone https://github.com/telman03/go-microservices.git
cd go-microservices
```

### 2️⃣ **Set up environment variables**
Create a `.env` file in both `user-service` and `notification-service` directories:

#### Example `.env` for User Service
```
PORT=8080
DATABASE_URL=postgres://user:password@db:5432/userdb
```

#### Example `.env` for Notification Service
```
KAFKA_BROKER=kafka:9092
KAFKA_TOPIC=order-events
KAFKA_GROUP=notification-group
DATABASE_URL=postgres://user:password@db:5432/notifications
EMAIL_HOST=smtp.example.com
EMAIL_PORT=587
EMAIL_USER=your_email@example.com
EMAIL_PASSWORD=your_password
```

### 3️⃣ **Run the services with Docker**
```sh
docker-compose up --build
```

### 4️⃣ **Run services locally**

#### Start User Service
```sh
cd user-service
go run main.go
```

#### Start Notification Service
```sh
cd notification-service
go run main.go
```

## 🛠️ API Endpoints

### **User Service API**
| Method | Endpoint       | Description         |
|--------|--------------|---------------------|
| GET    | `/users/:id` | Get user by ID |

### **Notification Service** (Consumes Kafka Events)
- Listens for `order-events` topic and sends email notifications.

## 📬 Kafka Event Format (Order Event)
```json
{
  "user_id": 1,
  "amount": 99.99
}
```

## 📩 Email Notification Example
```
Subject: Order Confirmation
Message: Your order of $99.99 has been placed successfully!
```

## 🏆 Contributing
Feel free to contribute by opening an issue or submitting a pull request!

## 📄 License
This project is open-source and available under the MIT License.

