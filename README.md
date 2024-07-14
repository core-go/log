# log
- wrapper of [zap](https://pkg.go.dev/go.uber.org/zap)
- wrapper of [logrus](https://github.com/sirupsen/logrus)
- middleware tracing log for [Echo](https://github.com/labstack/echo), [Gin](https://github.com/gin-gonic/gin), or net/http ([Gorilla mux](https://github.com/gorilla/mux), [Go-chi](https://github.com/go-chi/chi))

### A typical micro service
- When you zoom one micro service, the flow is as below, and you can see "log (zap, logrus)" and "middleware" in the full picture:
  ![A typical micro service](https://cdn-images-1.medium.com/max/800/1*d9kyekAbQYBxH-C6w38XZQ.png)

## Content for logging
### Request
#### Features
- <b>Log Request Method and URL</b>: Log the HTTP method (GET, POST, etc.) and the requested URL.
- <b>Log Request Headers</b>: Option to log request headers for debugging purposes.
- <b>Log Request Body</b>: Option to log the request body (with configurable size limits to avoid logging large payloads).
#### Benefits
- <b>Debugging</b>: Helps in tracing and debugging issues by providing complete information about incoming requests.
- <b>Monitoring</b>: Provides visibility into the types of requests being received.

### Response
#### Features
- <b>Log Response Status Code</b>: Log the HTTP status code of the response.
- <b>Log Response Headers</b>: Option to log response headers.
- <b>Log Response Body</b>: Option to log the response body (with configurable size limits to avoid logging large payloads).
#### Benefits
- <b>Debugging</b>: Assists in diagnosing issues by providing complete information about the responses sent by the server.
- <b>Auditing</b>: Helps in auditing and reviewing server responses for compliance and monitoring purposes.

### Response Time
#### Features
- <b>Log Response Time</b>: Calculate and log the time taken to process each request.
#### Benefits
- <b>Performance Monitoring</b>: Helps in identifying slow requests and performance bottlenecks.
- <b>Optimization</b>: Provides data to optimize and improve server response times.

### Response Size
#### Features
- <b>Log Response Size</b>: Log the size of the response payload in bytes.
#### Benefits
- <b>Bandwidth Monitoring</b>: Helps in monitoring and managing bandwidth usage.
- <b>Optimization</b>: Provides insights into the response sizes to optimize payloads and improve performance.

## Features
### Middleware Integration
#### Features
- <b>Middleware Function</b>: Designed to integrate seamlessly with existing Go libraries: [Echo](https://github.com/labstack/echo), [Gin](https://github.com/gin-gonic/gin), or net/http ([Gorilla mux](https://github.com/gorilla/mux), [Go-chi](https://github.com/go-chi/chi)).
  - Sample for [Echo](https://github.com/labstack/echo) is at [go-sql-echo-sample](https://github.com/go-tutorials/go-sql-echo-sample)
  - Sample for [Gin](https://github.com/gin-gonic/gin) is at [go-sql-gin-sample](https://github.com/go-tutorials/go-sql-gin-sample)
  - Sample for [Gorilla mux](https://github.com/gorilla/mux) is at [go-sql-sample](https://github.com/go-tutorials/go-sql-sample)
- <b>Context Handling</b>: Pass context to handle request-specific data throughout the middleware chain.
#### Benefits
- <b>Ease of Use</b>: Simplifies the integration of logging into existing web applications.
- <b>Consistency</b>: Ensures consistent logging across different parts of the application.

### Logging Libraries Integration
- Do not depend on any logging libraries.
- Already supported to integrate with [zap](https://pkg.go.dev/go.uber.org/zap), [logrus](https://github.com/sirupsen/logrus)
- Can be integrated with any logging library.

### Enable/Disable Logging
#### Features
- <b>Enable/Disable Logging</b>: Allow users to turn on or off logging for requests, responses, headers, and bodies independently.
- <b>Logging Levels</b>: Support different logging levels (e.g., INFO, DEBUG, ERROR) to control the verbosity of logs.
#### Benefits
- <b>Flexibility</b>: Provides users with the flexibility to configure logging based on their needs and environment.
- <b>Efficiency</b>: Reduces overhead by allowing selective logging, especially in production environments.

### Asynchronous Logging
#### Features
- <b>Non-Blocking Logs</b>: Implement asynchronous logging to ensure that logging does not block request processing.
- <b>Log Buffering</b>: Use buffering to improve logging performance and reduce latency.
#### Benefits:
- <b>Performance</b>: Improves the overall performance of the application by reducing logging overhead.
- <b>Scalability</b>: Allows the application to handle high-throughput logging without performance degradation.

### Sensitive Data Encryption
#### Features
- Mask/Encrypt sensitive data in the request and response bodies.
  - Sample for [Echo](https://github.com/labstack/echo) is at [go-sql-echo-sample](https://github.com/go-tutorials/go-sql-echo-sample)
  - Sample for [Gin](https://github.com/gin-gonic/gin) is at [go-sql-gin-sample](https://github.com/go-tutorials/go-sql-gin-sample)
  - Sample for [Gorilla mux](https://github.com/gorilla/mux) is at [go-sql-sample](https://github.com/go-tutorials/go-sql-sample)
- Sensitive Data Identification: identify and encrypt specific fields in JSON payloads.

#### Benefits:
- <b>Security</b>: Protects sensitive information from being exposed in logs.
- <b>Compliance</b>: Helps meet security and compliance requirements by safeguarding sensitive data.
- <b>Ease of Use</b>: Simplifies the integration of encryption/masking into any existing applications.
- <b>Consistency</b>: Ensures that sensitive data is consistently encrypted or masked across all logged requests and responses

## Use Cases of sensitive data masking/encrypting
### Financial Transactions
- <b>Benefit</b>: Encrypting sensitive financial data, such as credit card numbers and transaction details, helps comply with PCI-DSS standards and secures financial transactions from exposure in logs.
### Healthcare
- <b>Benefit</b>: Encrypting patient data such as medical records and health information in logs ensures compliance with HIPAA regulations and protects patient privacy.
### E-commerce
- <b>Benefit</b>: Protecting customer information, such as addresses and payment details, enhances customer trust and protects the e-commerce platform from potential data breaches.

## Benefits to Developers
#### Enhanced Debugging
- Provides detailed logs for requests and responses, aiding in troubleshooting and debugging issues.
#### Performance Monitoring
- Logs response times and sizes, allowing developers to monitor and optimize application performance.
#### Flexibility
- Configurable logging settings enable developers to tailor the logging behavior to their needs.
#### Ease of Integration
- Middleware function can be easily integrated into existing web frameworks, simplifying the setup process.
#### Improved Maintainability
- Centralized logging logic ensures consistency and makes the codebase easier to maintain.
#### Security:
- Encrypt or mask sensitive data in logs, reducing the risk of data exposure and meeting compliance requirements.

## Conclusion
By implementing these features, you provide a comprehensive logging solution that enhances the visibility, performance, and maintainability of any GO applications.

## Installation
Please make sure to initialize a Go module before installing core-go/log:

```shell
go get -u github.com/core-go/log
```

Import:
```go
import "github.com/core-go/log"
```

## middleware log tracing
middleware log for
- http
- echo
- gin

### Features
#### log tracing at middleware
##### Support to turn on, turn off
- request
- response
- duration
- http response status code
- response content length
##### Support to mask or encrypt fields
- support to mask or encrypt fields, such as mobileNumber, creditCardNumber

## Appendix
### Microservice Architect
![Microservice Architect](https://cdn-images-1.medium.com/max/800/1*vKeePO_UC73i7tfymSmYNA.png)

### A typical micro service
- When you zoom one micro service, the flow is as below, and you can see "middleware" in the full picture:
  ![A typical micro service](https://cdn-images-1.medium.com/max/800/1*d9kyekAbQYBxH-C6w38XZQ.png)

### Cross-cutting concerns
- "middleware" in the full picture of cross-cutting concerns
  ![cross-cutting concerns](https://cdn-images-1.medium.com/max/800/1*y088T4NoJNrL9sqrKeSyqw.png)
