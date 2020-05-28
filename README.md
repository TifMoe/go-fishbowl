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
$ make local-backend
```
You should keep this terminal open to see logs from the go service

In another terminal start the frontend with:
```console
$ make local-frontend
```
*Note: If this is your first time running the app, you will have to run `npm install` before `npm start`.*

You can view the frontend by navigating to http://localhost:3000

### Testing
Run the backend unit tests with:
```console
$ make test
```

### Contributing
To contribute to this project please create a feature branch and open a PR against the **staging** branch.
