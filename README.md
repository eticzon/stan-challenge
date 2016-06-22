## Stan Coding Challenge

My solution to https://challengeaccepted.streamco.com.au/.

## Setup

To setup dev dependencies, `cd` to the project directory and run:

```sh
make setup
```

The app is a standalone binary and doesn't require/vendor in any third-party dependencies.
All things `stdlib` bru! :)

To build and run:

```sh
$ make test
$ make build
$ bin/stan-challenge
```

That will start HTTP server which will bind/listen on `0.0.0.0:8080` (default port).

The default port can also be overridden with an environment variable `PORT`.

```sh
$ PORT=5000 bin/stan-challenge
```

## Heroku Deployment

Everything that is required to get this app deployed to [Heroku](https://www.heroku.com/) is already in-place. 
Refer to [Deploying Go Apps on Heroku](https://devcenter.heroku.com/articles/deploying-go for the) for details.

## Running Tests

Tests are written with standard [Go](https://golang.org/) tooling.

```sh
$ make test
```

## Author(s)

* Erwin Ticzon

## License: MIT

MIT License. &copy; 2016 Erwin Ticzon
