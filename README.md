# log
- wrapper of zap
- wrapper of logrus
- middleware tracing log

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

### Microservice Architect
![Microservice Architect](https://cdn-images-1.medium.com/max/800/1*vKeePO_UC73i7tfymSmYNA.png)

### A typical micro service
- When you zoom one micro service, the flow is as below, and you can see "middleware" in the full picture:
  ![A typical micro service](https://cdn-images-1.medium.com/max/800/1*d9kyekAbQYBxH-C6w38XZQ.png)

### Cross-cutting concerns
- "middleware" in the full picture of cross-cutting concerns
  ![cross-cutting concerns](https://cdn-images-1.medium.com/max/800/1*y088T4NoJNrL9sqrKeSyqw.png)
