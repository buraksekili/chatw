chatw provides a client and server to imitate basic `chat` application.

## Installation

```shell script
git clone https://github.com/buraksekili/chatw.git
```

## Setup server 

The port 8080 is used by default. 
A custom port can be configured with `-p` flag as; 
`-p <PORT_NUMBER>`

For example, to run server on port 3000;
```shell script
go run server/main.go -p 3000
```

## Setup client 

The client dials a server specified through command-line argument.
By default, the client dials `:8080`

For example, to dial server running on `localhost:8080`;
```shell script
go run client/main.go localhost:8080
```
