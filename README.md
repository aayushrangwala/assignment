## Setup
- Go installation: 1.12.7
- Editor/IDE for Go: Goland
- golanci-lint
- Postman (To test the endpoints)

## General Instructions
- Do not start a new task before you finished the one before.
- Put this project in a Git repository.
- Each step should be at least one separate commit.
- Do not commit binaries.
- Go code should be formatted.
- When you finished the tasks, push all commits to the gitlab repo created for you.

## Tasks
1. Build a go program.
   * Create a `go.mod` file containing `module discovergy`
   * Create a `main.go` which outputs "Hello World" https://gobyexample.com/hello-world
   * Compile your program with `go build`
   
2. Build a Go service with an HTTP endpoint which returns the requested path.
   * `curl http://127.0.0.1:3333/xxx` should return `xxx` (https://golang.org/pkg/net/http/)
   * `xxx` should be replaceable by any random characters. Any URL should work!
   
3. Write a test which verifies that the response is correct.
   * Create file named `main_test.go` and put the function `func TestSimplePath(t *testing.T) {...}` in there
   * Write a test with  which checks if the returned value is correct
   * The path should be a generated random string
   * Run the tests with `go test`
   
4. Write a `Makefile` with the following content:
   ```
   export CGO_ENABLED=0
   export GO111MODULE=on
   
   build:
	     go build -o www
   ```
   * What does each line do?
   * What do you expect to happen if you type `make`?
   * Extend the `Makefile` with the targets:
     - `test` should run `go test`
     - `run`  should compile `main.go` to a file named `www` and run it
   
5. Add a new HTTP endpoint `http://127.0.0.1:3333/encode/xxx`
   * The endpoint should return the encrypted text according to [Caeser cipher](https://en.wikipedia.org/wiki/Caesar_cipher) with right shift of 32 letters:
     * Alphabet: `abcdefghijklmnopqrstuvwxyz`
   * Use a primitive algorithm to encrypt the text
     * `http://localhost:3333/decode/discovergy` => joyiubkxme
   * Write a test case for this.
   
6. Add another HTTP endpoint `http://127.0.0.1:3333/decode/xxx`

   - This endpoint should do the decode the text encoded in 5.
     - `http://localhost:3333/decode/joyiubkxme => discovergy

7. Build a Docker container which runs your Go service
   * Create a `Dockerfile` from `alpine:latest`
   * Build the `Dockerfile` with `docker build -t www .` and run it
   * Extend the `Makefile` with the targets:
     - `docker-build`
     - `docker-run` should run in foreground
   * let it run

8. BONUS:

   - How can you improve the encryption/decryption algorithm?
   - How can you reduce the Docker image size?
   - Add Gitlab CI file to run the tests automatically