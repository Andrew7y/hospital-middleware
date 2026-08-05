# Hospital Middleware API

Hospital Middleware API เป็น RESTful API ที่ทำหน้าที่เป็น Middleware ระหว่างระบบ Hospital Information System (HIS) กับผู้ใช้งาน (Staff)

ระบบรองรับการ

- Staff Registration
- Staff Login (JWT Authentication)
- Patient Search
- Pagination
- Sorting
- Filtering
- Reverse Proxy ด้วย Nginx

---

# Tech Stack

- Go 1.25.12
- Gin
- GORM
- PostgreSQL
- JWT Authentication
- Docker
- Docker Compose
- Nginx

---

# Architecture

```
                 Client
                    │
                    │ HTTP
                    ▼
             +--------------+
             |    Nginx     |
             +--------------+
                    │
                    ▼
             +--------------+
             | Gin REST API |
             +--------------+
                    │
             +------+------+
             |             |
             ▼             ▼
      PostgreSQL      JWT Authentication
```

Hospital Middleware API นี้จะดึงข้อมูลจากฐานข้อมูลของตนเองเพื่อรองรับการค้นหาผู้ป่วยแบบกรองหลายเงื่อนไข โดยในโปรเจกต์นี้จำลองข้อมูลจาก HIS ผ่าน SQL Seed

---

# Project Structure

```
hospital-middleware
├── cmd/
│   └── main/
│       └── main.go
│
├── docs/
│   ├── development-planning.docx
│
├── internal/
│   ├── config/
│   ├── dto/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── router/
│   └── service/
│
├── pkg/
│   ├── normalize/
│   ├── password/
│   ├── query/
│   ├── token/
│
├── nginx/
│   └── nginx.conf
│
├── test/
│   └── patient_hanler_test.go
│
├── seed/
│   └── seed.sql
│
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

# Database Schema

ประกอบด้วย 4 ตาราง

## hospitals

| Column | Description |
|---------|-------------|
| id | Primary Key |
| name | Hospital Name |
| api_url | HIS API URL |

---

## patients

| Column | Description |
|---------|-------------|
| id | Primary Key |
| national_id | National ID |
| date_of_birth | Date of Birth |
| gender | Gender |

---

## patient_hospitals

| Column | Description |
|---|---|
| patient_id | Primary Key |
| hospital_id | Primary Key |
| first_name_th | First Name in Thai |
| last_name_th | Last Name in Thai |
| middle_name_th | Middle Name in Thai|
| first_name_en | First Name in English|
| middle_name_en | Middle Name in English |
| last_name_en | Last Name in English|
| passport_id | Passport ID |
| phone_number | Phone Number|
| email | Email Address |
| patient_hn | Patient Hospital Number (HN)|
---

## staffs

| Column      | Description |
|-------------|-------------|
| id          | Primary Key |
| username    | Username    |
| password    | Password    |
| hospital_id | Hospital ID |

---

# Features

- Staff Registration
- Staff Login
- JWT Authentication
- Search Patient
- Filter
- Pagination
- Sorting

---

# API

## Create Staff

```
POST http://localhost/api/staff/create
```

Request

```json
{
    "username":"admin",
    "password":"123456",
    "hospital_id":1
}
```

---

## Login

```
POST http://localhost/api/staff/login
```

Request

```json
{
    "username":"admin",
    "password":"123456"
}
```

Response

```json
{
    "message":"Login successful",
    "data":{
        "access_token":"...",
        "expires_at":"..."
    }
}
```

---

## Search Patient

```
GET http://localhost/api/patient/search
```

Header

```
Authorization: Bearer <JWT Token>
```

Example

```
GET http://localhost/api/patient/search?page=1&page_size=10
```

Example

```
GET http://localhost/api/patient/search?first_name_th=สม
```

Example

```
GET http://localhost/api/patient/search?patient_hn=BK000001
```

Example

```
GET http://localhost/api/patient/search?sort_by=first_name_en&sort_order=asc
```

---

# Running Project

## 1 Clone

```
git clone https://github.com/Andrew7y/hospital-middleware.git
```

```
cd hospital-middleware
```

---

## 2 Create Environment File

สร้างไฟล์

```
.env
```

ตัวอย่างข้อมูล

```
DB_USER = root
DB_PASSWORD = root
DB_NAME = hospital_db
DB_PORT = 5432

JWT_SECRET = 8fK2m9LpX1vQz7rT0wYe6aNu3iBc5dEfG8hJiJkLmNo

```

---

## 3 Build

```
docker compose up -d --build
```

---

## 4 Seed Database

Powershell
```
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
Get-Content -Encoding utf8 .\seed\seed.sql | docker compose exec -T postgres psql -U root -d hospital_db
```
---

## 5 Create Staff

```
POST http://localhost/api/staff/create
```

---

## 6 Login

```
POST hhttp://localhost/api/staff/login
```

นำ JWT Token ที่ได้ไปใช้งาน

---

## 7 Search Patient

```
GET http://localhost/api/patient/search
```

Header

```
Authorization: Bearer <token>
```

---

# Unit Test

รันทั้งหมด

```
go test ./...
```

หรือ

```
go test -cover ./...
```

---

# Test Cases

| No | Test Case |
|----|-----------|
| 1 | Create Staff Success |
| 2 | Cannot Create Duplicate Username |
| 3 | Cannot Create Empty Username or Password |
| 4 | Login Success |
| 5 | Login Failed (Incorrect Password) |
| 6 | Login Failed (Username Not Found) |
| 7 | Search Patient Success |
| 8 | Search Patient Without Login |

---

# Reverse Proxy

Nginx ทำหน้าที่เป็น Reverse Proxy

```
Client
    │
    ▼
 Nginx :80
    │
    ▼
 Go API :8080
    │
    ▼
 PostgreSQL
```

---

# Future Improvements

- HIS API Integration
- Refresh Token
- Role-based Authorization
- Swagger / OpenAPI
- Redis Cache
- Audit Log
- Docker Health Check
- CI/CD Pipeline

---

# Author

Kamphaengphet Singkhon