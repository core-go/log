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
![Microservice Architect](https://camo.githubusercontent.com/cf46a1780520d3612f1d81b219b56a14428fc24bb4ae9f4eede169aa9c58bee8/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a764b6565504f5f5543373369377466796d536d594e412e706e67)

### A typical micro service
- When you zoom one micro service, the flow is as below, and you can see "middleware" in the full picture:
  ![A typical micro service](https://camo.githubusercontent.com/581033268b9152e7ea8881904f533a51a29eeb3a63e8d6478540668c6e422ce3/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a64396b79656b416251594278482d4336773338585a512e706e67)

### Cross-cutting concerns
- "middleware" in the full picture of cross-cutting concerns
  ![cross-cutting concerns](https://camo.githubusercontent.com/0416e6d9aa090b3b42901b4dd22b19c8962abe6c589988b1e97dea97b63a278d/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a7930383854344e6f4a4e724c397371724b65537971772e706e67)
