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

game steps
1. server create room and wait for conn
2. client 1 connect to server, join the room and receive his team (X/O) >> TO1
3. client 2 connect to server, join the room and receive his team (X/O) >> TX2
4. client 1 play                                                        >> SO0
5. client 1 and 2 receive game changed event from server                >> SO0
6. client 1 and 2 receive game over event                               >> WO1
