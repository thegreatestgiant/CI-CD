![code coverage badge](https://github.com/thegreatestgiant/CI-CD/actions/workflows/ci.yml/badge.svg)

## Local Development

Make sure you're on Go version 1.20+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8000"
```

Run the server:

```bash
go build -o notely && ./notely
```

_This starts the server in non-database mode._ It will serve a simple webpage at `http://localhost:8000`.

You do _not_ need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!
