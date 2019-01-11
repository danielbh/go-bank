## Go Banking App

This is not meant to be a real banking app. It is to showcase go.

- ['realize' for task runner](https://github.com/oxequa/realize) It must be installed
`go get github.com/oxequa/realize`

#### Commands:

- **make install**: install dependencies
- **make dev**: start development
- **make build**: build go binary

#### Docker

`docker build . -t container-name`
`docker run -d -p 8080:8080 container-name`

#### Future Work
- [ ] Make better validation messages with exact strings of keys
- [ ] send error on negative deposit
- [ ] https://github.com/gorilla/csrf
- [ ] setup race detector test
- [ ] add map
- [ ] add interface
- [ ] make folders domain driven
- [ ] create auth folder with apikey checking
- [ ] live reload webpack page when go server changes
- [ ] document How to setup realize
- [ ] use godocs
- [ ] testing
- [ ] data store
- [ ] add authentication
- [ ] make multiple accounts
- [ ] add linting
- [ ] Make templates part of binary: https://github.com/gin-gonic/gin#build-a-single-binary-with-template
