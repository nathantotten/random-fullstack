# Full Stack App (Vue + TS + Vite Frontend, Go + PostgreSQL Backend)

This project is a simple full-stack application built for random tinkering, experimentation, and learning purposes. It consists of two main parts:

1. **Frontend**: A Vue 3 application powered by TypeScript and Vite for fast development.
2. **Backend**: A Go REST service using net/http with PostgreSQL as the database.

---

## Table of Contents

- [Full Stack App (Vue + TS + Vite Frontend, Go + PostgreSQL Backend)](#full-stack-app-vue--ts--vite-frontend-go--postgresql-backend)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Frontend Setup](#frontend-setup)
    - [Steps to Set Up the Frontend](#steps-to-set-up-the-frontend)
  - [Backend Setup](#backend-setup)
    - [Steps to Set Up the Backend](#steps-to-set-up-the-backend)
  - [Running the App](#running-the-app)
  - [Technologies Used](#technologies-used)
  - [License](#license)

---

## Overview

This project serves as a playground for experimenting with new technologies and features. It combines a modern frontend with Vue 3, TypeScript, and Vite, alongside a backend powered by Go and PostgreSQL.

You can use this repository to explore the following:

* Vue 3 with TypeScript for building dynamic and reactive user interfaces.
* Go and net/http for building RESTful APIs and handling business logic.
* PostgreSQL for managing relational data.

---

## Frontend Setup

The frontend is built using Vue 3 and TypeScript. The project is scaffolded with Vite for a fast and efficient development experience.

### Steps to Set Up the Frontend

1. **Install Dependencies**:
   Make sure you have [pnpm](https://pnpm.io/) installed, then install the dependencies:

   ```bash
   pnpm install
   ```

2. **Run the Development Server**:
   To start the frontend development server, run:

   ```bash
   pnpm dev
   ```

   This will start the Vite dev server, and you can access the frontend at [http://localhost:3000](http://localhost:3000).

---

## Backend Setup

The backend is a Go application that communicates with a PostgreSQL database.

### Steps to Set Up the Backend

1. **Install Go**:
   Make sure you have [Go](https://golang.org/dl/) installed.

2. **Configure PostgreSQL** (optional, if database features are added):
   Make sure PostgreSQL is installed and running. You will need to configure your database credentials in the application code or environment variables.

3. **Run the Go Application**:
   To start the backend, run the following command in the `backend/` directory:

   ```bash
   go run cmd/api/main.go
   ```

   The backend will be available at [http://localhost:8080](http://localhost:8080).

---

## Running the App

With both the frontend and backend running, you can visit the app in your browser.

1. Start the backend by running the Go application (`go run cmd/api/main.go`).
2. Start the frontend development server (`pnpm dev`).

---

## Technologies Used

* **Frontend**:

  * Vue 3
  * TypeScript
  * Vite (for fast development)
* **Backend**:

  * Go
  * PostgreSQL

---

## License

This project is licensed under the MIT License.
