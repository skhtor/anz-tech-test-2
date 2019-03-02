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

### Notes

I decided to challenge myself with this test and write it using a language I've never used before, Golang.

I found certain things really interesting and easy, particularly how you can compile the app into a static binary, and have it run on a scratch container, which contains nothing but the binary itself.

Other things, however, I found quite difficult. Initially I had the `BUILD_VERSION` and `COMMIT_SHA` injected into a JSON file, which the app would then read and `Marshal` into a struct. However I found while going down that path, I made it especially difficult to write unit tests.

I not only had to figure out a way to test HTTP calls, I also had to figure out how to stub out JSON file reads and how to handle a global config object in the test case in a language I've never written in before.

I decided to change my implementation and have my application read those variables from environment variables, which were made available by the Dockerfile. This turned out to me a much simpler implementation, although I would've preferred to have had proper config management.

While frustrating at times having so much trouble doing things I would've otherwise done so easily in a more familiar language (like Python), I am so glad I decided to do this in Go, as I have learned so much about this language that I otherwise wouldn't had the opportunity to learn.
