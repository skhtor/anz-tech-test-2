## ANZ Technical Test 2

This repo contains a simple Golang web app, which has a single endpoint ("/healthcheck") that returns some metadata about the app itself

### Local Development

#### Requirements

- docker
- docker-compose

#### Building the image

Export environment variables

```
export BUILD_VERSION=<my_version>
export COMMIT_SHA=<my_commit_sha>
```

Then build the container using docker-compose

`docker-compose build go-anz`

#### Running the container

`docker run -p 8000:8000 go-anz`

Open `localhost:8000` in your browser

#### Running the tests
Run linting and unit tests using the following command:

`docker-compose run tests`

### CI

This app uses Travis CI as its CI engine.
Please see the `.travis.yml` file for build steps
