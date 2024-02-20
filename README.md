# xo-go
xo-go is a centralized multi player tic tac toe game

the server will create, accept players and control the game state

the client will render the new game state for the players

server will listen connections on port 8888
```shell
make runserver
```

clients will connect to server on port 8888
```shell
make runclient
```
run client twice to connect players 1 and 2

to run tests or build just use the commands bellow:
```shell
# for build apps
make build

# for running tests
make test
```

# how it works
after stebelishing connetions, server and client will exchange messages using a fixed binary string of length 3 content as following:

- first character is for the changed state
- second is for the team
- third is for the square

we group this 3 informations into a type that we named statement

## Changed State
- 'T': for team selection
- 'B': for board changes
- 'O': for game over

## Team
- 'X': for X team
- 'O': for O team

## Square
we map the tic tac toe board using an array of length 9, so each index represent an square in the board as following:

```
0 | 1 | 2
---------
3 | 4 | 5
---------
6 | 7 | 8
```

## Exchange state between server / client
1. server create room and wait for conn
2. client 1 connect to server, join the room and receive his team (X/O) i.e.: TO0
3. client 2 connect to server, join the room and receive his team (X/O) i.e.: TX0
4. client 1 play, it will send to the server an board change            i.e.: BO0
5. server accept de board change and send to both clients board change  i.e.: BO0
6. client 1 and 2 receive game over state*                              i.e.: GO8

*game over state works similar to the board change, except they will inform the winner