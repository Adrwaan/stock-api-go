# 📈 Stock API REST in Golang

A RESTful API service for stock market data management built with Go. This API provides endpoints for retrieving, managing, and processing stock market information.

## 🔍 Overview

This project implements a stock market API using Go's standard library and routing capabilities. The application follows a modular architecture with separated concerns for routing, business logic, and data handling.

## 📁 Project Structure

```
stock-api-go/
├── main.go          # Application entry point
├── routes/          # HTTP route handlers and middleware
└── README.md        # Project documentation
```

## 🚀 Getting Started

### 📋 Prerequisites

- Go 1.19 or higher
- Git

### 💾 Installation

1. Clone the repository:

```bash
git clone https://github.com/Adrwaan/stock-api-go.git
cd stock-api-go
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the application:

```bash
go run main.go
```

## 🏗️ Architecture

The application is structured with the following components:

- **main.go**: Entry point that initializes the application by calling the routes bootstrap function
- **routes/**: Contains HTTP route definitions, handlers, and middleware configuration
- **models/**: Contains data models and business logic for product management, including CRUD operations
- **config/**: Handles application configuration and logging setup

## 🔗 API Endpoints

[Documentation for specific endpoints would be added here based on the routes implementation]

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

MIT License

Copyright (c) 2025 Adrean Danilo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
