# xo-go
xo-go is a remote multi player tic tac toe game

the server will create, accept players and control the game state

the client will render the new game state for the player

server will listen connections on port 8889
```shell
make runserver
```

clientes will connect to server on port 8889
```shell
make runclient
```

to run tests or build just use the commands bellow:
```shell
# for build apps
make build

# for running tests
make tests
```