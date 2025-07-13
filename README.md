# Full Stack App (Vue + TS + Vite Frontend, Spring + PostgreSQL Backend)

This project is a simple full-stack application built for random tinkering, experimentation, and learning purposes. It consists of two main parts:

1. **Frontend**: A Vue 3 application powered by TypeScript and Vite for fast development.
2. **Backend**: A Spring Boot application with PostgreSQL as the database.

---

## Table of Contents

- [Full Stack App (Vue + TS + Vite Frontend, Spring + PostgreSQL Backend)](#full-stack-app-vue--ts--vite-frontend-spring--postgresql-backend)
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

This project serves as a playground for experimenting with new technologies and features. It combines a modern frontend with Vue 3, TypeScript, and Vite, alongside a robust backend powered by Spring Boot and PostgreSQL.

You can use this repository to explore the following:

* Vue 3 with TypeScript for building dynamic and reactive user interfaces.
* Spring Boot for building RESTful APIs and handling business logic.
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

The backend is a Spring Boot application that communicates with a PostgreSQL database.

### Steps to Set Up the Backend

1. **Install Java and Maven**:
   Make sure you have [Java](https://adoptopenjdk.net/) and [Maven](https://maven.apache.org/) installed.

2. **Configure PostgreSQL**:
   Make sure PostgreSQL is installed and running. You will need to configure your database credentials in the `application.properties` file (located in `backend/src/main/resources/`):

   ```properties
   spring.datasource.url=jdbc:postgresql://localhost:5432/mydb
   spring.datasource.username=your-username
   spring.datasource.password=your-password
   ```

3. **Run the Spring Boot Application**:
   To start the backend, run the following command in the `backend/` directory:

   ```bash
   ./mvnw spring-boot:run
   ```

   The backend will be available at [http://localhost:8080](http://localhost:8080).

---

## Running the App

With both the frontend and backend running, you can visit the app in your browser.

1. Start the backend by running the Spring Boot application (`./mvnw spring-boot:run`).
2. Start the frontend development server (`pnpm dev`).

---

## Technologies Used

* **Frontend**:

  * Vue 3
  * TypeScript
  * Vite (for fast development)
* **Backend**:

  * Spring Boot
  * PostgreSQL

---

## License

This project is licensed under the MIT License.
