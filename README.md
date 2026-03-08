# MindBox

MindBox is a GraphQL backend for structured productivity and technical learning management.

Instead of simple to-do lists, MindBox models productivity as a hierarchy of **Roadmaps → Categories → Tasks**, allowing long-term planning while also supporting **daily habits and knowledge tracking**.

The project focuses on:

- clear domain modeling
- clean separation between API, domain and persistence
- a reproducible development environment
- a strongly typed GraphQL API

---

# Overview

MindBox organizes work in three levels:

Roadmap -> Category -> Task

**Roadmaps** represent long-term goals.

**Categories** group tasks within a roadmap and define time constraints.

**Tasks** represent actionable items and may optionally behave as recurring daily habits.

This structure allows combining long-term planning with short-term execution.

---

# Features

### Roadmap Management

Create structured learning paths or project plans.

Examples:

- Backend Engineering
- System Design
- Kubernetes

Each roadmap contains multiple categories.

---

### Categories with Deadlines

Categories represent phases inside a roadmap.

Each category includes a **lifetime** field representing its expiration time.

If the deadline is reached and tasks remain incomplete, the category becomes **expired**.

---

### Task Tracking

Tasks are the smallest unit of work and include:

- completion status
- optional expiration
- optional daily recurrence

Examples:

- Implement GraphQL API
- Study database indexing
- Read system design literature

---

### Daily Tasks

Daily tasks support habit tracking.

Examples:

- Code for one hour
- Study system design
- Write "3 things I learned today"

Daily tasks are designed to reset every day.

---

### Layer Responsibilities

**GraphQL Layer**

Handles the API surface.

- schema definitions
- query and mutation resolvers
- GraphQL model types

**Domain Layer**

Defines core business entities.

- roadmap
- category
- task

**Repository Layer**

Encapsulates persistence logic.

- database queries
- entity hydration
- domain persistence rules

---

# Tech Stack

Backend

- Go
- gqlgen
- PostgreSQL

Infrastructure

- Docker
- Docker Compose
- golang-migrate

API

- GraphQL

---

# GraphQL Schema

## Roadmap

type Roadmap {
id: ID!
name: String!
description: String
categories: [Category!]!
}

---

## Category

type Category {
id: ID!
name: String!
description: String
lifetime: String!
roadmap: Roadmap!
tasks: [Task!]!
}

---

## Task

type Task {
id: ID!
name: String!
description: String
done: Boolean!
isDaily: Boolean!
isExpired: Boolean!
lifetime: String
category: Category!
}

---

# API Usage

## Queries

### Retrieve a roadmap

```graphql
query {
  roadmap(id: "roadmap-id") {
    id
    name
    categories {
      id
      name
    }
  }
}
List tasks
query {
  tasks(limit: 10, offset: 0) {
    id
    name
    done
  }
}
Daily tasks
query {
  dailyTasks(limit: 10, offset: 0) {
    id
    name
  }
}
Expired tasks
query {
  expiredTasks(limit: 10, offset: 0) {
    id
    name
  }
}
Mutations
Create roadmap
mutation {
  createRoadmap(
    input: {
      name: "Backend Engineering"
      description: "Learning backend architecture"
    }
  ) {
    id
    name
  }
}
Create category
mutation {
  createCategory(
    input: {
      name: "GraphQL"
      lifetime: "30d"
      roadmapId: "roadmap-id"
    }
  ) {
    id
  }
}
Create task
mutation {
  createTask(
    input: {
      name: "Build GraphQL API"
      isDaily: false
      categoryId: "category-id"
    }
  ) {
    id
  }
}
Toggle task completion
mutation {
  toggleTaskDone(id: "task-id") {
    id
    done
  }
}
```

Running the Project
Requirements

Docker

Docker Compose

Environment Variables

Create a .env file in the project root:

PORT=8080

DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=mindbox
Start the Services
docker-compose up --build

This will start:

PostgreSQL database

migration service

GraphQL API server

GraphQL Playground

After startup the playground is available at:

http://localhost:PORT/

Design Considerations
GraphQL

GraphQL was chosen to allow flexible querying of nested structures such as:

roadmaps with categories

categories with tasks

This avoids multiple REST endpoints for hierarchical data.
