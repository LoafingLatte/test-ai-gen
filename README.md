# LBK Points API

A comprehensive REST API for managing users and point transfers in the LBK Points system. Built with Go, Fiber web framework, SQLite database, and GORM ORM.

## 🎯 Features

### User Management (CRUD)
- ✅ Create users with membership information
- ✅ Read/retrieve users by ID or list all users
- ✅ Update user information and point balances
- ✅ Delete users from the system
- ✅ Track membership levels (Gold, Silver, Platinum, Bronze)

### Point Transfer System
- ✅ Atomic point transfers between users
- ✅ Automatic idempotency key generation for tracking
- ✅ Validates sender has sufficient points
- ✅ Prevents self-transfers
- ✅ Maximum transfer limit: 2.00 (200 cents)
- ✅ Automatic point balance updates
- ✅ Transaction audit trail with point ledger

### Data Management
- ✅ SQLite database with GORM ORM
- ✅ Automatic database migrations
- ✅ Point ledger (append-only) for transaction history
- ✅ Foreign key constraints and indexes
- ✅ Transaction support for atomicity

## 🏗️ Architecture

```
workshop3/
├── main.go                          # Application entry point
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
│
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   │
│   ├── database/
│   │   └── init.go                 # Database initialization & migrations
│   │
│   ├── models/
│   │   ├── user.go                 # User model
│   │   └── transfer.go             # Transfer & PointLedger models
│   │
│   ├── handler/
│   │   ├── user.go                 # User CRUD handlers
│   │   └── transfer.go             # Transfer operation handlers
│   │
│   └── routes/
│       └── routes.go               # Route registration
│
├── pkg/
│   └── response/
│       └── response.go             # Response formatting utilities
│
├── swagger.yml                      # OpenAPI 3.0.0 specification
├── database.md                      # Database documentation & ER diagram
├── README.md                        # This file
├── test.sh                          # User CRUD test script
└── test_transfers.sh                # Transfer API test script
```

## 🚀 Getting Started

### Prerequisites
- Go 1.19 or higher
- SQLite (included)
- curl (for testing)

### Installation

1. **Clone the repository**
```bash
git clone <repository-url>
cd workshop3
```

2. **Install dependencies**
```bash
go mod download
```

3. **Build the application**
```bash
go build -o workshop3
```

4. **Run the application**
```bash
./workshop3
```

The server will start on `http://localhost:3000`

### Docker Support (Optional)

```bash
# Build Docker image
docker build -t lbk-points-api .

# Run in Docker
docker run -p 3000:3000 -v ./data:/app/data lbk-points-api
```

## 📚 API Documentation

### Base URL
```
http://localhost:3000
```

### Health Check
```bash
GET /health
```

### User Endpoints

#### Create User
```bash
POST /api/v1/users
Content-Type: application/json

{
  "membership_id": "LBK001234",
  "first_name": "สมชาย",
  "last_name": "ใจดี",
  "phone": "081-234-5678",
  "email": "somchai@example.com",
  "membership_date": "15/6/2566",
  "membership_level": "Gold",
  "points": 15420
}
```

**Response (201 Created):**
```json
{
  "status": 201,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "membership_id": "LBK001234",
    "first_name": "สมชาย",
    "last_name": "ใจดี",
    "phone": "081-234-5678",
    "email": "somchai@example.com",
    "membership_date": "15/6/2566",
    "membership_level": "Gold",
    "points": 15420,
    "created_at": 1697548992,
    "updated_at": 1697548992
  }
}
```

#### Get All Users
```bash
GET /api/v1/users
```

#### Get User by ID
```bash
GET /api/v1/users/{id}
```

#### Update User
```bash
PUT /api/v1/users/{id}
Content-Type: application/json

{
  "membership_level": "Platinum",
  "points": 25000
}
```

#### Delete User
```bash
DELETE /api/v1/users/{id}
```

### Transfer Endpoints

#### Create Transfer
```bash
POST /api/v1/transfers
Content-Type: application/json

{
  "fromUserId": 1,
  "toUserId": 2,
  "amount": 100,
  "note": "ขอบคุณสำหรับช่วยงาน"
}
```

**Response (201 Created):**
```json
{
  "status": 201,
  "message": "Transfer created successfully",
  "data": {
    "transfer": {
      "idemKey": "a8b4f2e0556264f1c9b622a2f2f4c9b1",
      "transferId": 1,
      "fromUserId": 1,
      "toUserId": 2,
      "amount": 100,
      "status": "completed",
      "note": "ขอบคุณสำหรับช่วยงาน",
      "createdAt": "2025-10-17T14:03:12Z",
      "updatedAt": "2025-10-17T14:03:12Z",
      "completedAt": "2025-10-17T14:03:12Z"
    }
  }
}
```

**Response Header:**
```
Idempotency-Key: a8b4f2e0556264f1c9b622a2f2f4c9b1
```

#### List Transfers for User
```bash
GET /api/v1/transfers?userId=1&page=1&pageSize=20
```

**Query Parameters:**
- `userId` (required): User ID to filter transfers
- `page` (optional): Page number (default: 1)
- `pageSize` (optional): Items per page, max 200 (default: 20)

**Response:**
```json
{
  "status": 200,
  "message": "Transfers retrieved successfully",
  "data": {
    "data": [
      {
        "idemKey": "a8b4f2e0556264f1c9b622a2f2f4c9b1",
        "transferId": 1,
        "fromUserId": 1,
        "toUserId": 2,
        "amount": 100,
        "status": "completed",
        "createdAt": "2025-10-17T14:03:12Z",
        "updatedAt": "2025-10-17T14:03:12Z",
        "completedAt": "2025-10-17T14:03:12Z"
      }
    ],
    "page": 1,
    "pageSize": 20,
    "total": 42
  }
}
```

#### Get Transfer by Idempotency Key
```bash
GET /api/v1/transfers/{idempotency_key}
```

Example:
```bash
GET /api/v1/transfers/a8b4f2e0556264f1c9b622a2f2f4c9b1
```

## ⚠️ Validation Rules

### User Validation
- First name and last name required (non-empty)
- Email must be unique
- Membership ID must be unique
- Phone number format must be valid

### Transfer Validation
- **Amount Constraints**:
  - Minimum: 1 cent (0.01)
  - Maximum: 200 cents (2.00)
  - Must be positive integer

- **User Constraints**:
  - Both users must exist
  - Sender and receiver must be different
  - Sender must have sufficient points

- **Transfer Uniqueness**:
  - Idempotency key is unique
  - Prevents duplicate transfers from retried requests

### Error Responses

**400 Bad Request:**
```json
{
  "status": 400,
  "message": "Invalid request parameters",
  "error": "fromUserId, toUserId, and amount must be valid"
}
```

**404 Not Found:**
```json
{
  "status": 404,
  "message": "User not found",
  "error": ""
}
```

**409 Conflict:**
```json
{
  "status": 409,
  "message": "Insufficient points",
  "error": "User does not have enough points to transfer"
}
```

**422 Unprocessable Entity:**
```json
{
  "status": 422,
  "message": "Cannot transfer to yourself",
  "error": "fromUserId and toUserId must be different"
}
```

## 🧪 Testing

### Manual Testing with curl

```bash
# Create user 1
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"membership_id":"LBK001","first_name":"User","last_name":"One","phone":"081-111-1111","email":"user1@example.com","membership_date":"15/6/2566","membership_level":"Gold","points":500}'

# Create user 2
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"membership_id":"LBK002","first_name":"User","last_name":"Two","phone":"081-222-2222","email":"user2@example.com","membership_date":"15/6/2566","membership_level":"Gold","points":100}'

# Create transfer
curl -X POST http://localhost:3000/api/v1/transfers \
  -H "Content-Type: application/json" \
  -d '{"fromUserId":1,"toUserId":2,"amount":100,"note":"Test transfer"}'
```

### Automated Testing

```bash
# Run user CRUD tests
./test.sh

# Run transfer API tests
./test_transfers.sh
```

## 📊 Database

### Schema
See `database.md` for comprehensive database documentation including:
- Entity Relationship Diagram (Mermaid format)
- Table structures and constraints
- Indexes for query optimization
- Access patterns and common queries

### Database File Location
```
./app.db (SQLite)
```

### Migrations
Migrations are automatic on startup using GORM AutoMigrate:
- `users` table
- `transfers` table
- `point_ledger` table

## 🔐 Security Considerations

1. **Input Validation**: All inputs are validated before database operations
2. **SQL Injection Prevention**: Using parameterized queries via GORM
3. **Atomic Transactions**: All transfers are atomic (all-or-nothing)
4. **Idempotency**: Unique idempotency keys prevent duplicate operations
5. **Audit Trail**: Point ledger maintains complete transaction history

### Future Enhancements
- [ ] Authentication and authorization (JWT)
- [ ] Rate limiting
- [ ] Request logging and monitoring
- [ ] Encryption at rest for sensitive data
- [ ] API key management

## 🚦 Deployment

### Environment Variables
```bash
# Database
DB_PATH=./app.db          # SQLite database file location

# Server
SERVER_PORT=:3000         # Server port
```

### Production Checklist
- [ ] Set up monitoring and logging
- [ ] Configure database backups
- [ ] Enable HTTPS/TLS
- [ ] Implement rate limiting
- [ ] Set up authentication
- [ ] Configure CORS policies
- [ ] Use environment-specific configs

## 📝 Dependencies

```
github.com/gofiber/fiber/v2    # Web framework
gorm.io/gorm                     # ORM
gorm.io/driver/sqlite            # SQLite driver
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 API Documentation

Complete API documentation is available in:
- **Swagger/OpenAPI**: See `swagger.yml`
- **Database Schema**: See `database.md`
- **Examples**: See test scripts (`test.sh`, `test_transfers.sh`)

### View Swagger UI

To view the API documentation interactively, use Swagger UI:

```bash
# Using Docker
docker run -p 8080:8080 -e SWAGGER_JSON=/swagger.yml -v $(pwd)/swagger.yml:/swagger.yml swaggerapi/swagger-ui

# Then visit: http://localhost:8080
```

Or use online Swagger Editor: https://editor.swagger.io/

## 📞 Support

For issues and questions:
1. Check existing documentation in `swagger.yml` and `database.md`
2. Review test scripts for usage examples
3. Check database schema for data structure details

## 📄 License

This project is licensed under the MIT License - see LICENSE file for details.

## 🎓 Learning Resources

- [Go Documentation](https://golang.org/doc)
- [Fiber Framework Docs](https://docs.gofiber.io)
- [GORM Documentation](https://gorm.io)
- [OpenAPI Specification](https://spec.openapis.org)

---

**Version**: 1.0.0  
**Last Updated**: October 17, 2025  
**Built with**: Go, Fiber, SQLite, GORM
