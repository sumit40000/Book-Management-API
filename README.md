📚 Book Management API (Go + MongoDB)

A beginner-friendly RESTful API built with Golang and MongoDB that allows users to manage books — including creating, reading, updating, and deleting book records.
This project was created to learn the fundamentals of Go backend development, routing, and database connections.

🚀 Features

Create a new book 📖

Fetch all books

Fetch a single book by ID

Update a book (mark as read ✅)

Delete a book 🗑️

MongoDB Atlas connection for cloud database

Environment variable configuration using .env

Clean folder structure (cmd, internal/config, internal/controllers, internal/routes)

🧱 Project Structure
```text
book-management-api/
│
├── cmd/
│   └── main.go                  # Entry point of the project
│
├── internal/
│   ├── config/                  # MongoDB connection setup
│   │   └── config.go
│   │
│   ├── controllers/             # Controller logic (CRUD)
│   │   └── controllers.go
│   │
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   │
│   └── models/                  # Book model
│       └── models.go
│
├── go.mod                       # Go module file
├── go.sum
├── .env                         # Environment variables
└── README.md
```

⚙️ Tech Stack

Language: Go (Golang)

Database: MongoDB Atlas

Routing: Gorilla Mux

Environment Variables: godotenv

Driver: mongo-driver

🔧 Installation & Setup

Clone the repository
```text
git clone https://github.com/<your-username>/book-management-api.git
cd book-management-api
```

Install dependencies
```text
go mod tidy
```

Create a .env file
```text
MONGO_URI=mongodb+srv://<username>:<password>@cluster0.mongodb.net/?appName=Cluster0
PORT=:8000
DB_NAME = bookdb
```

Run the server
```text
go run ./cmd/main.go
```

Build the project (optional)
```text
go build -o ./bin/book-management-api ./cmd/main.go
```
📡 API Endpoints
```text
Method	Endpoint	Description
GET	/books	Get all books
GET	/books/{id}	Get a single book
POST	/books	Create a new book
PUT	/books/{id}	Update a book (mark as read)
DELETE	/books/{id}	Delete a book
```
🧠 Sample JSON (for POST request)
```text
{
  "title": "The Go Programming Language",
  "author": "Alan Donovan",
  "price": 599.99,
  "rating": 4.8,
  "read": false
}
```

🪄 Example Request (Using cURL)
```text
curl -X POST http://localhost:8000/books \
-H "Content-Type: application/json" \
-d '{
  "title": "Atomic Habits",
  "author": "James Clear",
  "price": 450.00,
  "rating": 4.9,
  "read": false
}'
```
🧑‍💻 Author

Sumit Vishwakarma
Beginner backend developer exploring Go, MongoDB, and scalable API design.

⭐ Support

If you like this project, don’t forget to star 🌟 the repo on GitHub — it helps a lot!
