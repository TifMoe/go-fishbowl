# go-fishbowl
## check out the game here: [fishbowl.rocks](https://fishbowl.rocks/)

A Go + React web app to play FishBowl remotely

## Getting started

This project requires Go, npm and Docker to be installed. On OS X with Homebrew you can install dependencies with:
| Dependency | Homebrew Install |
|---|---|
| Go | `brew install go` |
| npm | `brew install node` |
| Docker | `brew install docker` | 


## Local Dev
For development, you might find it easier to run the backend and frontend separately (outside of the docker container) so you can watch frontend changes. 

Starting the go server:
```console
$ docker-compose up db
$ go run cmd/go-fishbowl/main.go
```
You can hit the api directly to start games and add cards `curl POST localhost:8080/v1/api/game/`

And in another terminal starting the frontend:
```console
$ cd ui/ && npm start
```
You can view the frontend by navigating to http://localhost:3000

To deploy the web app locally should be as simple as:
```console
$ make local
```
Then navigate to http://localhost:8080 to check out a preview of the dockerized app as it will be deployed to production

### Testing
Run the backend unit tests with:
```console
$ make test
```
