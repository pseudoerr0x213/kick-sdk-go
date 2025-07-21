# Kick-SDK-Go 

## SDK Go-Client for Kick API 

## Key features 

- Efficient and fast  
- Easy to integrate with Go web applications 
- Concurrent-safe -- compile to binary and use with goroutines

## Done: 

- DTO models 
- Project layout 
- HTTP layer + auth with options (app / user flows as defined in "https://docs.kick.com/getting-started/generating-tokens-oauth2-flow") 
- Wire up the SDK logic in entrypoint 

## ToDo:  
- Define retry logic as a func option for Client 
- Add request/response logging
- Add rate limiting support
- Add custom error types 
- Structured Logging with custom errors 
- Full concurrency support