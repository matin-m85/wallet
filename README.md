# wallet

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

Start each service individually, or use Docker Compose (coming soon).  
After all services are running → Gateway automatically connects to them.

---

## 📦 Project Highlights

- 100% isolated microservices  
- No shared DB — each service owns its own domain  
- Fast, asynchronous-friendly HTTP communication  
- Easy to extend: add new services without breaking old ones  
- Perfect for learning microservices with Go  

---
