# Go Skeleton Project
![Coverage](https://img.shields.io/badge/Coverage-26.7%25-red)

This is a skeleton project for kickstarting Go applications. It provides a basic directory structure and configuration setup to help you get started with your Go projects quickly.

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)


## Project Structure

The project structure follows the standard Go project layout:

```
├── build
├── cmd
├── internal
│   ├── config
│   ├── constants
│   ├── container
│   ├── delivery
│   │   ├── queue
│   │   └── rest
│   │       ├── handler
│   │       │   └── healthcheck
│   │       ├── middleware
│   │       └── model
│   │           ├── request
│   │           └── response
│   ├── entity
│   ├── infra
│   │   ├── cache
│   │   └── database
│   ├── mocks
│   │   ├── delivery
│   │   │   └── rest
│   │   │       ├── handler
│   │   │       │   └── healthcheck
│   │   │       ├── middleware
│   │   │       └── model
│   │   │           └── response
│   │   ├── infra
│   │   │   ├── cache
│   │   │   └── database
│   │   ├── repository
│   │   │   └── cache
│   │   └── usecase
│   │       └── healthcheck
│   ├── repository
│   │   └── cache
│   └── usecase
│       └── healthcheck
└── resource
```

- build: This directory is typically used for storing build artifacts generated during the build process, such as compiled binaries or object files.

- cmd: The cmd directory contains the main application entry points or command-line interfaces (CLIs). Each Go file within this directory represents a separate application or command.

- internal: The internal directory houses the core application logic and functionalities. It's divided into several subdirectories:
  - config: Contains files related to application configuration, such as reading configuration files or environment variables.
  - constants: Holds constant values used throughout the application, such as error codes or status constants.
  - container: Contains files related to dependency injection or inversion of control (IoC) container setup.
  - delivery: This directory handles incoming requests and routes them to appropriate handlers or controllers. It's further divided into:
    - queue: Contains files related to handling message queues (if applicable).
    - rest: Handles REST API requests. It's divided into:
      - handler: Contains HTTP request handlers or controllers. Each subdirectory within handler may represent a different endpoint or resource.
      - middleware: Houses middleware functions that intercept and preprocess incoming HTTP requests or responses.
      - model: Holds request and response models used by the REST API.
  - entity: Contains domain entities or data models used within the application.
  - infra: Houses files related to external integrations or infrastructure components. It's divided into:
    - cache: Contains files related to caching mechanisms.
    - database: Contains files related to database connections, migrations, or ORM configurations.
  - mocks: Holds mock implementations used for testing. It mirrors the internal package structure and contains mock implementations for different layers of the application.
  - repository: Contains repository interfaces and implementations responsible for data access operations.
  - usecase: Contains use case interfaces and implementations representing business logic or application workflows. It's further divided into specific use case domains, such as healthcheck.
- resource: This directory typically stores static resources or configuration files used by the application, such as YAML or JSON configuration files.

Each directory in the project structure serves a specific purpose and helps organize the codebase into manageable components, promoting modularity and maintainability.

## Installation
1. Clone the repository to your local machine:
```
git clone https://github.com/nofendian17/go_starter_kit.git
```
2. Navigate to the project directory:
```
cd go_starter_kit
```
3. configure application & database
```
cp resource/example.config.yaml resource/config.yaml
```
4. running your application
```
go run main.go
```