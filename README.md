### Goals

See `Subject.txt`
### Usage & tests

First, make sure you have docker, curl, bash, and go installed

execute `./startServices.sh`

if you don't have permission
execute `cmod +x startServices.sh`

then execute
`bash startServices.sh`

When all the services are running, you can launch the tester by executing
`bash test.sh`

## Roadmap

- [x] Implement creation of file for ever user added in DB
- [x] Implement creation of file for ever user added in DB
- [x] Implement POST addUsers with json input
- [x] Implement DELETE deleteUserById
- [x] Implement GET getUserById
- [x] Implement GET getUserList
- [x] Implement PATCH user
- [x] Implement POST Login with Jwt
- [x] Implement concurrence for every user added (go routines)
- [x] Implement creation of file for ever user added in DB
- [x] Update server userFile when user Document is updated in DB
- [x] Dockerize MongoDB
- [x] Hash user password in DB
- [ ] Implement refresh token
- [ ] Implement Mongo-Express
- [ ] Implement secrets and protect MongoDB
- [ ] Improve flexibility of the Router configuration
- [ ] Restrict size and complexity of credentials (regex)
- [ ] Reorganize the authentication middleware(gin way)
- [ ] Split login function
- [ ] Add pepper to protect passwords in DB (protect from rainbow tables)
- [ ] Improve tests in golang instead of bash
- [ ] improve granularity of updates
- [ ] Improve management of panic
- [ ] Improve data validation (Go Validator Package)
- [ ] Improve serialization of inputs
- [ ] Improve verbose
- [ ] Match RFC standards for requests responses
- [ ] Optimize some functions and make them more flexible (especially for data inputs)
- [ ] Implement an nginx load balancer with SSL
- [ ] Dockerize all services (API / Mongo-Express ) with persistent volume