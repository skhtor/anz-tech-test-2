## ANZ Technical Test 2

### Local Development

#### Requirements

- Docker

#### Building the image

`docker-compose build go-anz`

#### Running the container

`docker run -p 8000:8000 go-anz`

Open `localhost:8000` in your browser

#### Running the tests

`docker-compose run tests`
