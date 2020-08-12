# Reload server
```
reflex -d none -s -R vendor. -r \.go$ make
```
 
## Problems
- Need a way to consider the distributed transaction problem, using rabbitMQ....
- Need to return a proper error still...validation error can be 200, but internal errors has to be 500
- how to get the transaction working????