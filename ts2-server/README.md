TS2 Server
==========

This is a clone of the [TS2 server](https://github.com/ts2/ts2-sim-server), the core of the ts2 simulator, that is to be extended as part of the BahnTS2 project.

Starting the server
-------------------
Assuming that you have the Go language distribution (https://golang.org/dl/) on your local machine, run the following 
command in a teminal:

```bash
go run main.go /path/to/simulation-file.json
```

The server is accessible accessed via websocket at `ws://localhost:22222/ws`

> Note that the server only accepts JSON simulation files. 
> If you have a `.ts2` file, you must unzip it first, extract the `simulation.json` file inside and start the server with it.

Web UI
------
The server ships with a minimal Web UI to interact with the webservice.

Start the server and head to `http://localhost:22222`.
