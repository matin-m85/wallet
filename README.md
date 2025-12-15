# 🚀 Microservices Wallet & Gift Code Platform – API Gateway

A modern, clean, and scalable microservices ecosystem built in **Go**.  
This system simulates a real-world digital wallet & gift-code workflow using separate services for the **Wallet**, **Gift Code**, and a centralized **API Gateway**.  
Designed for simplicity, modularity, and production-like behavior. ⚡🔥

---

## 🎯 Project Goal

This project demonstrates how multiple independent microservices can coordinate financial operations:

- Users have personal wallets 💰  
- Admins create **gift-code groups** containing multiple redeemable codes 🎁🔑  
- Users redeem codes → wallet balance increases instantly  
- API Gateway manages communication between services  
- Each service runs independently and is fully isolated  

The architecture showcases distributed service communication through clean, lightweight HTTP interactions.

---

## 🧱 Architecture Overview

| Service | Responsibility | Tech |
|--------|----------------|------|
| 🏦 **Wallet Service** | Wallet management, transactions, balance updates | Go, Fiber, PostgreSQL (optional), in-memory fallback |
| 🎁 **Gift Code Service** | Create gift-code groups, generate unique codes, track usage | Go, Fiber, PostgreSQL (optional), in-memory storage |
| 🌐 **API Gateway** | Unified API layer, request routing, service orchestration | Go, Fiber, Docker-ready |

---

## ⚙️ How the Flow Works

When a user redeems a **gift code**:

1. 🌐 API Gateway → sends code + phone number to Gift Code Service  
2. 🎁 Gift Code Service → validates code & returns the redeem amount  
3. 🏦 Gateway → updates wallet balance using Wallet Service  
4. 📦 Gateway → returns merged response containing wallet + code usage info  

This provides a smooth, real-world multi-service transaction and demonstrates service-to-service communication.

---

## 🏦 Wallet Service (Summary)

- Create new user wallets 🪪  
- Add balance to a wallet ➕💰  
- List all wallets  
- Track wallet transactions 📜  
- Uses clean struct-based models  
- Stores data either **in-memory** or via **PostgreSQL** (configurable)  
- Acts as the financial engine of the system  

---

## 🎁 Gift Code Service (Summary)

- Create **gift-code groups**  
- Auto-generate unique codes 🔑  
- Track **used / unused** codes  
- View statistics for each gift-code group 📊  
- See which users redeemed codes  
- Fetch full details of each code  
- Supports in-memory storage with easy switch to PostgreSQL  
- Ensures each code is redeemable only once  

---

## 🌐 API Gateway (Summary)

- Single point of external access 🌍  
- Routes & validates incoming requests  
- Communicates with Wallet & Gift-Code services  
- Combines multi-service responses in one clean JSON output  
- Built with **Fiber** for maximum performance ⚡  
- Fully Docker-ready 🐳  

---

## 🛠 Tech Stack

### 💻 Languages & Frameworks
- **Go 1.21+**  
- **Fiber Web Framework**  
- **Go Modules**  

### 🗄 Storage Options
- **PostgreSQL** (planned or optional)  
- **In-Memory Storage** (default)  

### 🐳 DevOps & Tooling
- **Docker & Docker Compose**  
- **RESTful Service Design**  
- **Clean Architecture (Modular Services)**  

---

## ▶️ Running the Services

| Service | Port |
|--------|------|
| 🏦 Wallet | `8081` |
| 🎁 Gift Code | `8082` |
| 🌐 Gateway | `8080` |

Services can be run individually for development,  
or all together using **Docker Compose (recommended)**.

---

## 📦 Project Highlights

- 100% isolated microservices  
- No shared DB — each service owns its own domain  
- Fast, asynchronous-friendly HTTP communication between services  
- Easy to extend: add new services without breaking old ones  
- Perfect for learning microservices with Go  

---
## 📁 Project Structure

The project follows a clean and easy-to-navigate microservices architecture:

app/
├─ wallet-service/
│  ├─ cmd/
│  │  └─ main.go
│  ├─ config/
│  │  └─ config.go
│  └─ internal/
│     ├─ models/
│     │  ├─ wallet.go
│     │  └─ transaction.go
│     ├─ repository/
│     │  ├─ db.go
│     │  ├─ wallet_repository.go
│     │  └─ transaction_repository.go
│     ├─ service/
│     │  └─ wallet_service.go
│     └─ api/
│        ├─ controllers/
│        │  └─ wallet_controller.go
│        ├─ routes/
│        │  └─ routes.go
│        └─ server.go
│
├─ gift-service/
│  ├─ cmd/
│  │  └─ main.go
│  ├─ config/
│  │  └─ config.go
│  └─ internal/
│     ├─ models/
│     │  |─ gift_group.go
|     |  └─ giftـcard.go
│     ├─ repository/
│     │  ├─ db.go
│     │  |─ gift_card_repo.go
│     │  └─ gift_group_repo.go
│     ├─ service/
│     │  └─ gift_service.go
│     └─ api/
│        ├─ controllers/
│        │  └─ gift_controller.go
│        ├─ routes/
│        │  └─ routes.go
│        └─ server.go
│
└─ api-gateway/
   ├─ cmd/
   │  └─ main.go
   ├─ config/
   │  └─ config.go
   └─ internal/
      ├─ client/
      │  ├─ wallet_client.go
      │  └─ gift_client.go
      ├─ service/
      │  └─ gateway_service.go
      └─ api/
         ├─ controllers/
         │  └─ gateway_controller.go
         ├─ routes/
         │  └─ routes.go
         └─ server.go

----

## 🔌 Wallet API Endpoints

### 🧱 1) Create New Wallet
curl -X POST http://localhost:8080/api/wallet \
  -H "Content-Type: application/json" \
  -d '{"phone":"09123456789"}'
  

### ➕ 2) Add Balance
curl -X POST http://localhost:8080/api/wallet/add \
  -H "Content-Type: application/json" \
  -d '{"phone":"09123456789","amount":1000000,"reference":"test"}'


### 🔎 3) Get Wallet Info
curl -X GET http://localhost:8080/api/wallet/09123456789


### 📜 4) Get Wallet Transactions
curl -X GET http://localhost:8080/api/wallet/09123456789/transactions


### 📋 5) List All Wallets
curl -X GET http://localhost:8080/api/wallet/list


## 🔌 Gift Code API Endpoints

### 🎁 1) Create a Gift Code Group (e.g., Yalda)
curl -X POST http://localhost:8080/api/group/create \
  -H "Content-Type: application/json" \
  -d '{"name":"yalda","amount_due":50000,"count":10}'


### 🎁 2) Use a Gift Code
curl -X POST http://localhost:8080/api/use \
  -H "Content-Type: application/json" \
  -d '{"code":"YOUR_GIFT_CODE_HERE","phone":"09123456789"}'


### 🧾 3) Group Statistics Report
curl -X GET http://localhost:8080/api/group/YOUR_GROUP_ID_HERE/stats


### 📜 4) List of Users Who Redeemed Gift Codes
curl -X GET http://localhost:8080/api/group/YOUR_GROUP_ID_HERE/users


### 🔍 5) Get Gift Code Information
curl -X GET http://localhost:8080/api/card/YOUR_GIFT_CODE_HERE


### 📋 6) List All Groups with Statistics
curl -X GET http://localhost:8080/api/group/list


### 🧾 7) List All Codes in a Group
curl -X GET http://localhost:8080/api/group/YOUR_GROUP_ID_HERE/codes

---

# ▶️ How to Run the Project (Docker-Based)

This project is fully **Dockerized** 🐳 and designed to run all services together using **Docker Compose**.

✔ Each microservice has its **own Dockerfile**  
✔ A single **docker-compose.yml** orchestrates the entire system  
✔ One command → all services up and running  

---

## 🧩 Prerequisites

Make sure you have the following installed:

- Docker  
- Docker Compose  

Verify installation:

docker --version
docker compose version

---

## 1️⃣ Clone the Repository

git clone https://github.com/matin-m85/wallet.git

---

## 2️⃣ Build & Run All Services

From the **project root**, simply run:

docker compose up --build

---


Docker Compose will automatically:

- Build all service images  
- Start **Wallet Service**, **Gift Code Service**, and **API Gateway**  
- Create an isolated Docker network  
- Connect services using internal service names  

No manual configuration needed ✅

---

## 3️⃣ Service Ports

| Service | Port |
|-------|------|
| 🏦 Wallet Service | 8081 |
| 🎁 Gift Code Service | 8082 |
| 🌐 API Gateway | 8080 |

After startup, access the system via:

http://localhost:8080


All external requests must go through the **API Gateway** 🌐

---

## 4️⃣ Stopping the Project

To stop all services:

docker compose down



