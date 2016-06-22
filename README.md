## Stan Coding Challenge

My solution to https://challengeaccepted.streamco.com.au/.

## Setup

The app is a standalone binary which can be compiled and run via (assuming you're in the project directory):

```sh
$ make test
$ bin/stan-challenge
```

This will start an HTTP server listening on `localhost:80`.

## Heroku Deployment

Everything that is required to get this app deployed to [Heroku](https://www.heroku.com/) is already in-place. 
Refer to [Deploying Go Apps on Heroku](https://devcenter.heroku.com/articles/deploying-go for the) for the details.

A sample app has been deployed to http://stan-challenge-solution.herokuapp.com

## Running Tests

Tests are written with standard go tooling.

```sh
$ make test
```

Generating coverage reports:

```sh
$ make cover
```

## Author(s)

* Erwin Ticzon

## License: MIT

MIT License. &copy; 2016 Erwin Ticzon