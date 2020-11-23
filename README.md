# Quick start

This gin middleware will output the bodies of each request-response pair.   

The middleware is expensive due to extra buffer copies, so it should be only used for trace level.   
Logrus logger is required.   

```go
// Set to trace and use BodyLogger
log.SetLevel(log.TraceLevel)
if log.GetLevel() == log.TraceLevel {
  router.Use(gintrace.BodyLogger)
}
```

