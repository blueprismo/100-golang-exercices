# Exercise 97: Fiber Web Framework Basics

## Introduction to Fiber

Fiber is a modern Go web framework inspired by Express.js, designed for building fast and scalable web applications and APIs.

## Key Fiber Features

- **Zero Memory Allocation**: Efficient memory usage
- **Fast HTTP Server**: Built on top of Fasthttp
- **Express-like API**: Familiar routing and middleware patterns
- **Built-in Middleware**: CORS, Logger, Recovery, and more
- **Template Engines**: Support for various template engines

## Basic Fiber Application

```go
app := fiber.New()

app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})

app.Listen(":3000")
```

## Context Object

The `*fiber.Ctx` context provides:
- **Request Data**: Headers, parameters, body
- **Response Methods**: Send JSON, HTML, files
- **Middleware**: Next() for middleware chaining
- **Utilities**: Parsing, validation, caching

## Route Definition

Fiber supports all HTTP methods:
- **GET**: Retrieve data
- **POST**: Create resources
- **PUT**: Update resources
- **DELETE**: Remove resources
- **PATCH**: Partial updates

## Response Methods

Common response methods:
- **SendString()**: Plain text response
- **JSON()**: JSON response
- **HTML()**: HTML response
- **SendFile()**: File download
- **Redirect()**: HTTP redirects

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a new Fiber application instance
2. Implementing the helloWorld handler function
3. Setting up basic routes with setupRoutes function
4. Understanding Fiber's application structure

This exercise introduces the Fiber web framework and establishes the foundation for building web APIs.