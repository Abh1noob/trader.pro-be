# 🚀 Trader.pro Backend

> **High-performance trading simulation engine with real-time data processing**

A robust, scalable backend service built with Go Fiber that powers the Trader.pro paper trading platform. Features real-time market data integration, advanced trade execution engine, and comprehensive portfolio management.

## 🏗️ Architecture

### Core Components
- **🎯 Trade Engine**: Real-time order processing and execution
- **📊 Market Data Service**: ICICI Breeze API integration with caching
- **💼 Portfolio Manager**: Position tracking and P&L calculations
- **🔐 Auth Service**: Firebase JWT validation and user management
- **⏰ Time Travel Engine**: Historical data replay with configurable speed
- **📈 Analytics Engine**: Performance metrics and risk calculations

### Technical Stack
- **Framework**: Go Fiber (Fast, Express-inspired web framework)
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: Firebase Admin SDK
- **Market Data**: ICICI Breeze API
- **Caching**: Redis for session and market data
- **Deployment**: AWS EC2 with Docker
- **CI/CD**: GitHub Actions

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 14+
- Redis 6+
- ICICI Breeze API credentials
- Firebase Admin SDK key

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Abh1noob/trader.pro-be.git
   cd trader.pro-be
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Environment Setup**
   ```bash
   cp .env.example .env
   ```
   
   Configure your environment variables:
   
5. **Database Setup**
   ```bash
   # Run database migrations
   go run cmd/migrate/main.go
   
   # Seed initial data (optional)
   go run cmd/seed/main.go
   ```

6. **Run the server**
   ```bash
   # Development mode
   go run main.go
   
   # Or with live reload
   air
   ```

## 📁 Project Structure

```
trader.pro-be/
├── cmd/                    # Application entry points
│   ├── migrate/           # Database migrations
│   ├── seed/             # Database seeding
│   └── server/           # Main server
├── internal/              # Private application code
│   ├── auth/             # Authentication middleware
│   ├── config/           # Configuration management
│   ├── database/         # Database connection and models
│   ├── handlers/         # HTTP request handlers
│   ├── middleware/       # Custom middleware
│   ├── models/           # Data models and schemas
│   ├── repositories/     # Data access layer
│   ├── services/         # Business logic
│   └── utils/            # Utility functions
├── pkg/                   # Public library code
│   ├── breeze/           # ICICI Breeze API client
│   ├── cache/            # Redis caching utilities
│   └── validator/        # Input validation
├── docs/                  # API documentation
├── scripts/              # Deployment and utility scripts
└── tests/                # Test files
```

## 🔄 API Endpoints

### Authentication
```
POST   /api/v1/auth/login          # User login
POST   /api/v1/auth/refresh        # Refresh token
GET    /api/v1/auth/profile        # Get user profile
PUT    /api/v1/auth/profile        # Update profile
```

### Trading
```
POST   /api/v1/trades              # Place new trade
GET    /api/v1/trades              # Get trade history
GET    /api/v1/trades/:id          # Get specific trade
DELETE /api/v1/trades/:id          # Cancel pending trade
```

### Portfolio
```
GET    /api/v1/portfolio           # Get portfolio summary
GET    /api/v1/positions           # Get current positions
GET    /api/v1/positions/:symbol   # Get position for symbol
POST   /api/v1/positions/close     # Close position
```

### Market Data
```
GET    /api/v1/market/symbols      # Get available symbols
GET    /api/v1/market/quote/:symbol # Get real-time quote
GET    /api/v1/market/history/:symbol # Get historical data
GET    /api/v1/market/search       # Search symbols
```

### Simulation
```
POST   /api/v1/simulation/start    # Start simulation
POST   /api/v1/simulation/pause    # Pause simulation
POST   /api/v1/simulation/reset    # Reset simulation
GET    /api/v1/simulation/status   # Get simulation status
```

## 🎯 Core Features

### Real-time Market Data
- **Live Quotes**: Real-time price updates via WebSocket
- **Historical Data**: OHLCV data with configurable timeframes
- **Data Caching**: Redis-based caching for improved performance
- **Rate Limiting**: Intelligent API rate limiting and request batching

### Portfolio Management
- **Position Tracking**: Real-time position updates
- **P&L Calculations**: Unrealized and realized profit/loss
- **Risk Metrics**: Portfolio risk analysis and exposure calculations
- **Performance Analytics**: Historical performance tracking

## 🔗 Related Projects

- **Frontend**: [trader.pro-fe](https://github.com/Abh1noob/trader.pro-fe)
