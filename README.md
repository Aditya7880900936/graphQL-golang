# Job Details CRUD API

A **GraphQL CRUD API** for managing job listings, built with **Go**, **MongoDB**, and **GraphQL**.

---

## 🚀 Features

- **Full CRUD Operations:** Easily create, read, update, and delete job listings.
- **GraphQL API:** Flexible and efficient data queries and mutations.
- **MongoDB Integration:** Persistent and scalable NoSQL storage.
- **Clean & Modular Go Structure:** Organized, maintainable, and extensible backend.
- **Authentication & Basic Error Handling:** Secure API access and informative error feedback.
- **GraphQL Playground:** Interactive UI for testing and exploring API queries.
- **Auto-generated Code:** Rapid development with [gqlgen](https://github.com/99designs/gqlgen).

---

## 🗂️ Project Structure

```
.
├── database
│   └── database.go             # MongoDB connection and CRUD operations
├── graph
│   ├── generated               # Auto-generated GraphQL code
│   ├── model                   # GraphQL model definitions
│   ├── resolver.go             # GraphQL resolvers
│   ├── schema.graphqls         # GraphQL schema
│   └── schema.resolvers.go     # Resolver definitions
├── tools
│   └── tools.go                # Utility functions (if any)
├── go.mod                      # Go modules
├── go.sum
├── gqlgen.yml                  # gqlgen config
├── main.go                     # Application entry point
└── server.go                   # HTTP server setup and GraphQL handler
```

---

## ⚡️ Tech Stack

- **Go:** High-performance backend language.
- **GraphQL:** Modern API query language.
- **MongoDB:** NoSQL database for scalable storage.
- **gqlgen:** Go library for generating GraphQL servers.
- **GraphQL Playground:** In-browser IDE for API testing.

---

## 📦 Installation

1. **Clone the repository:**
    ```sh
    git clone <your-repo-url>
    cd graphql-golang
    ```

2. **Install Go dependencies:**
    ```sh
    go mod tidy
    ```

3. **Start MongoDB:**

    Make sure MongoDB is running locally at `mongodb://localhost:27017`.  
    The backend will use the database **graphql-golang**.

---

## 🏃 Getting Started

1. **Run the server:**
    ```sh
    go run server.go
    ```
    or (if your entry point is `main.go`):
    ```sh
    go run main.go
    ```

2. **Open the Playground:**

    Visit [http://localhost:8080/](http://localhost:8080/) in your browser to interact with the API using GraphQL Playground.

---

## 📚 GraphQL Schema at a Glance

```graphql
type JobListing {
  _id: ID!
  title: String!
  description: String!
  company: String!
  url: String!
}

type Query {
  jobs: [JobListing!]!
  job(id: ID!): JobListing!
}

type Mutation {
  createJobListing(input: CreateJobListingInput!): JobListing!
  updateJobListing(id: ID!, input: UpdateJobListingInput!): JobListing!
  deleteJobListing(id: ID!): DeleteJobResponse!
}

input CreateJobListingInput {
  title: String!
  description: String!
  company: String!
  url: String!
}

input UpdateJobListingInput {
  title: String
  description: String
  url: String
  company: String
}

type DeleteJobResponse {
  deleteJobId: String!
}
```

---

## 🔧 Usage Examples

### Query All Jobs

```graphql
query {
  jobs {
    _id
    title
    company
    url
  }
}
```

### Create a Job Listing

```graphql
mutation {
  createJobListing(input: {
    title: "Frontend Developer"
    description: "React.js & Next.js"
    company: "Tech Co"
    url: "https://techco.com/jobs/123"
  }) {
    _id
    title
    company
  }
}
```

### Update a Job Listing

```graphql
mutation {
  updateJobListing(id: "JOB_ID", input: {
    title: "Senior Frontend Developer"
  }) {
    _id
    title
  }
}
```

### Delete a Job Listing

```graphql
mutation {
  deleteJobListing(id: "JOB_ID") {
    deleteJobId
  }
}
```

---

## 💡 Next Steps & Suggestions

- Add **role-based access control (RBAC)** for enhanced security.
- Implement **pagination** and **filtering** for large datasets.
- Deploy the API using **Docker** and cloud providers (e.g., Heroku, AWS, GCP).
- Improve error handling and add custom validations.
- Add support for environment variables (e.g., using `dotenv` package).
- Write unit and integration tests.
- Add CI/CD for automated testing and deployment.

---

## 📷 Screenshots

<!-- Add UI/API screenshots here (if available) -->

---

## 📜 License

This project is open-source and available under the **MIT License**.
