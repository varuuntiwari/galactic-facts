# galactic-facts
An API to give random facts about space

# Install
Download the repository by cloning it locally using command: `git clone https://github.com/varuuntiwari/galactic-facts`

# Usage
The API can be run by building the binary with `go build api.go` or running it on a Docker container.
It runs on the port specified by the env variable `PORT`. On requesting the endpoint, it returns a JSON object of the form:
```
{
	data: [array of facts as strings],
	count: number of jokes as int,
	status: status of response returned as int,
}
```
Get a response from the API from the `/api/fact` endpoint, which also accepts a `count` parameter for getting more than one fact at a time. For example: `/api/fact?count=5` will return an array of 5 facts in the `data` parameter of response.

# Adding facts
For adding more facts, go to the `data/facts.go` file and add them to the `Facts` slice.

# Docker
Run the API as a Docker container by running the command `docker build .`. Also, the container can be pulled from Docker Hub by running `docker pull varuuntiwari/galactic-facts`. Make sure that Docker is installed on the system and the container is run with a PORT variable assigned, failing which the app will panic with an error message.