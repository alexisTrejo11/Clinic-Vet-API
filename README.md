# 🏥 Veterinary Clinic Management System API

A comprehensive RESTful API system built with Go for managing veterinary clinic operations, including customer management, pet records, appointments, medical history, payments, and staff administration.

## 📋 Table of Contents

- [Overview](#-overview)
- [Tech Stack](#-tech-stack)
- [System Architecture](#-system-architecture)
- [Business Modules](#-business-modules)
- [Prerequisites](#-prerequisites)
- [Environment Setup](#-environment-setup)
- [Docker Installation & Setup](#-docker-installation--setup)
- [API Documentation](#-api-documentation)
- [Development Workflow](#-development-workflow)
- [Testing](#-testing)
- [Contributing](#-contributing)

## 🎯 Overview

This veterinary clinic management system provides a complete backend solution for modern veterinary practices. It handles everything from customer registration and pet profiles to appointment scheduling, medical records, payment processing, and staff management.

### Key Features

- 👥 **Customer Management**: Complete customer profiles with contact information and history
- 🐕 **Pet Management**: Detailed pet records with medical history, vaccinations, and care notes
- 📅 **Appointment Scheduling**: Smart appointment booking with veterinarian assignments
- 🏥 **Medical Records**: Comprehensive medical history tracking and treatment documentation
- 💳 **Payment Processing**: Multi-currency payment handling with various payment methods
- 👨‍⚕️ **Staff Management**: Employee profiles, schedules, and specializations
- 🔐 **Authentication & Authorization**: JWT-based security with role-based access control
- 📧 **Notifications**: Email and SMS notifications via Twilio integration
- 📊 **Audit Logging**: Complete audit trail for all system operations

## 🛠 Tech Stack

### Programming Language
- **Go 1.23.0** - Primary backend language with high performance and concurrency

### Web Framework
- **Gin** - High-performance HTTP web framework for Go
- **Gin-CORS** - Cross-Origin Resource Sharing middleware

### Databases
- **PostgreSQL 12** - Primary relational database for structured data
- **MongoDB 7.0** - Document database for flexible data storage and logging
- **Redis 7** - In-memory data structure store for caching and session management

### Authentication & Security
- **JWT (golang-jwt/jwt/v5)** - JSON Web Token implementation
- **bcrypt** - Password hashing and validation
- **CORS middleware** - Cross-origin request handling

### Database Tools
- **pgx/v5** - PostgreSQL driver and toolkit
- **SQLC** - Type-safe SQL code generation
- **golang-migrate** - Database migration tool

### External Services
- **Twilio** - SMS notifications and communication
- **SMTP** - Email notifications

### Development & Testing
- **Testify** - Testing toolkit with assertions and mocks
- **GoMock** - Mock generation for testing
- **Swagger** - API documentation generation
- **Zap** - Structured logging
- **Lumberjack** - Log rotation

### Containerization
- **Docker** - Application containerization
- **Docker Compose** - Multi-container orchestration

## 🏗 System Architecture

The application follows **Clean Architecture** principles with clear separation of concerns:

```
├── app/
│   ├── config/          # Configuration management
│   ├── core/
│   │   ├── domain/      # Business entities and rules
│   │   ├── repository/  # Data access interfaces
│   │   └── service/     # Business logic services
│   ├── middleware/      # HTTP middleware (auth, CORS, rate limiting)
│   ├── modules/         # Feature-based modules
│   └── shared/          # Shared utilities and components
├── cmd/                 # Application entry points
├── db/                  # Database migrations and queries
└── docs/                # API documentation
```

### Architecture Patterns
- **CQRS (Command Query Responsibility Segregation)** - Separate read and write operations
- **Repository Pattern** - Data access abstraction
- **Dependency Injection** - Loose coupling between components
- **Middleware Pattern** - Cross-cutting concerns handling

## 📦 Business Modules

### 1. **Authentication Module** (`/auth`)
- User registration and login
- JWT token management
- Password reset functionality
- Role-based access control

### 2. **Customer Module** (`/customer`)
- Customer profile management
- Contact information handling
- Customer relationship tracking

### 3. **Pet Module** (`/pets`)
- Pet registration and profiles
- Medical information tracking
- Vaccination and treatment history
- Microchip and insurance management

### 4. **Employee Module** (`/employee`)
- Staff profile management
- Veterinarian specializations
- Schedule management
- License tracking

### 5. **Appointment Module** (`/appointment`)
- Appointment scheduling and management
- Service type categorization
- Status tracking (pending, confirmed, completed)
- Veterinarian assignment

### 6. **Medical Module** (`/medical`)
- Medical history documentation
- Diagnosis and treatment tracking
- Prescription management
- Emergency case handling

### 7. **Payment Module** (`/payments`)
- Multi-currency payment processing
- Payment method support (cash, cards, online)
- Invoice generation and management
- Payment status tracking and refunds

### 8. **Notifications Module** (`/notifications`)
- Email notification system
- SMS alerts via Twilio
- Appointment reminders
- System notifications

### 9. **User Management Module** (`/users`)
- User account administration
- Profile management
- Role and permission handling

## 📋 Prerequisites

Before setting up the project, ensure you have the following installed:

- **Docker** (v20.10 or higher)
- **Docker Compose** (v2.0 or higher)
- **Git** (for cloning the repository)

### Optional (for local development)
- **Go** (v1.23.0 or higher)
- **PostgreSQL** (v12 or higher)
- **MongoDB** (v7.0 or higher)
- **Redis** (v7.0 or higher)

## ⚙️ Environment Setup

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/Clinic-Vet-API.git
cd Clinic-Vet-API
```

### 2. Create Environment File

Create a `.env` file in the root directory with the following variables:

```bash
# PostgreSQL Configuration
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_secure_password
POSTGRES_DB=vet_database
TEST_POSTGRES_DB=vet_test_database

# MongoDB Configuration
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=mongopassword
MONGO_INITDB_DATABASE=vet_clinic

# Redis Configuration
REDIS_PASSWORD=redis_secure_password

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here

# Twilio Configuration (for SMS)
TWILIO_ACCOUNT_SID=your_twilio_account_sid
TWILIO_AUTH_TOKEN=your_twilio_auth_token
TWILIO_PHONE_NUMBER=your_twilio_phone_number

# SMTP Configuration (for Email)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password
FROM_EMAIL=noreply@vetclinic.com
FROM_NAME=Vet Clinic System

# Application Configuration
PROJECT_NAME=Veterinary Clinic Management System
LOGO_URL=https://your-domain.com/logo.png
```

**⚠️ Security Note**: Never commit the `.env` file to version control. Add it to your `.gitignore` file.

## 🐳 Docker Installation & Setup

### Step 1: Install Docker

#### On Ubuntu/Debian:
```bash
# Update package index
sudo apt update

# Install Docker
sudo apt install docker.io docker-compose

# Start and enable Docker
sudo systemctl start docker
sudo systemctl enable docker

# Add your user to docker group (optional)
sudo usermod -aG docker $USER
# Log out and back in for this to take effect
```

#### On macOS:
```bash
# Install Docker Desktop from https://docker.com/products/docker-desktop
# Or using Homebrew:
brew install --cask docker
```

#### On Windows:
Download and install Docker Desktop from [https://docker.com/products/docker-desktop](https://docker.com/products/docker-desktop)

### Step 2: Verify Docker Installation

```bash
# Check Docker version
docker --version
docker-compose --version

# Test Docker installation
docker run hello-world
```

### Step 3: Build and Run the Application

#### Quick Start (Recommended)

```bash
# Build and start all services
docker-compose up --build

# Or run in detached mode (background)
docker-compose up --build -d
```

#### Step-by-Step Process

1. **Build the services:**
   ```bash
   docker-compose build
   ```

2. **Start the databases first:**
   ```bash
   docker-compose up postgres12 mongodb redis -d
   ```

3. **Wait for databases to be ready (about 30-60 seconds), then start the API:**
   ```bash
   docker-compose up api
   ```

4. **Verify services are running:**
   ```bash
   docker-compose ps
   ```

### Step 4: Access the Application

Once the containers are running:

- **API Server**: [http://localhost:8000](http://localhost:8000)
- **API Documentation (Swagger)**: [http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)
- **PostgreSQL**: `localhost:5431` (main), `localhost:5433` (test)
- **MongoDB**: `localhost:27017`
- **Redis**: `localhost:6379`

### Step 5: Verify Setup

Test the API with a simple health check:

```bash
curl http://localhost:8000/health
```

Or visit the Swagger documentation to explore all available endpoints.

### Container Management Commands

```bash
# View running containers
docker-compose ps

# View logs
docker-compose logs api
docker-compose logs postgres12

# Stop all services
docker-compose down

# Stop and remove volumes (⚠️ This will delete all data)
docker-compose down -v

# Restart a specific service
docker-compose restart api

# Execute commands inside containers
docker-compose exec api sh
docker-compose exec postgres12 psql -U postgres -d vet_database
```

## 📚 API Documentation

### Swagger Documentation

The API documentation is automatically generated using Swagger and is available at:
- **Local**: [http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)

### API Endpoints Overview

```
Authentication:
POST   /api/v1/auth/register     - Register new user
POST   /api/v1/auth/login        - User login
POST   /api/v1/auth/refresh      - Refresh JWT token

Customers:
GET    /api/v1/customers         - List customers
POST   /api/v1/customers         - Create customer
GET    /api/v1/customers/:id     - Get customer details
PUT    /api/v1/customers/:id     - Update customer
DELETE /api/v1/customers/:id     - Delete customer

Pets:
GET    /api/v1/pets              - List pets
POST   /api/v1/pets              - Register new pet
GET    /api/v1/pets/:id          - Get pet details
PUT    /api/v1/pets/:id          - Update pet information
DELETE /api/v1/pets/:id          - Delete pet

Appointments:
GET    /api/v1/appointments      - List appointments
POST   /api/v1/appointments      - Schedule appointment
GET    /api/v1/appointments/:id  - Get appointment details
PUT    /api/v1/appointments/:id  - Update appointment
DELETE /api/v1/appointments/:id  - Cancel appointment

Medical Records:
GET    /api/v1/medical/:pet_id   - Get pet medical history
POST   /api/v1/medical           - Create medical record
PUT    /api/v1/medical/:id       - Update medical record

Payments:
GET    /api/v1/payments          - List payments
POST   /api/v1/payments          - Process payment
GET    /api/v1/payments/:id      - Get payment details
PUT    /api/v1/payments/:id      - Update payment status
```

## 🔧 Development Workflow

### Running Locally (Without Docker)

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Set up databases locally:**
   ```bash
   # PostgreSQL, MongoDB, and Redis must be installed and running
   ```

3. **Run migrations:**
   ```bash
   migrate -path db/migrations -database "postgresql://user:password@localhost/vet_database?sslmode=disable" up
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

### Database Operations

```bash
# Create new migration
migrate create -ext sql -dir db/migrations -seq add_new_table

# Run migrations
docker-compose exec api migrate -path ./db/migrations -database "postgresql://postgres:password@postgres12:5432/vet_database?sslmode=disable" up

# Rollback migrations
docker-compose exec api migrate -path ./db/migrations -database "postgresql://postgres:password@postgres12:5432/vet_database?sslmode=disable" down 1
```

### Code Generation

```bash
# Generate SQL code with SQLC
sqlc generate

# Generate Swagger documentation
swag init

# Generate mocks for testing
mockgen -source=app/core/repository/user_repository.go -destination=app/test/mock/user_repository_mock.go
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test package
go test ./app/modules/auth/application

# Run tests in Docker
docker-compose exec api go test ./...
```

### Test Database

The system uses a separate test database (`postgres12_test`) for running tests without affecting development data.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Write unit tests for new features
- Update documentation for API changes
- Use meaningful commit messages
- Ensure all tests pass before submitting PR

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

For support and questions:
- Create an issue in the GitHub repository
- Contact the development team
- Check the [API documentation](http://localhost:8000/swagger/index.html) for endpoint details

## 🚀 Deployment

For production deployment, consider:
- Using environment-specific configurations
- Implementing proper logging and monitoring
- Setting up CI/CD pipelines
- Configuring load balancers and scaling
- Implementing backup strategies for databases
- Setting up SSL/TLS certificates

---

**Built with ❤️ using Go and modern development practices**