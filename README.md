# go-socket.io-sample

Sample application of go-socket.io. This app can run on Heroku

## Prequisities

This app use [gom](https://github.com/mattn/gom) as package manager, so install gom first.

```
$ go get github.com/mattn/gom
```

## How to run on Local

```
$ git clone git@github.com:hakobera/go-socket.io-sample.git
$ cd go-socket.io-sample
$ gom install
$ PORT=3333 gom run main.go
```

Then open http://localhost:3333 on your browser.

## How to run on Heroku

Use [hakobera/heroku-buildpack-go](https://github.com/hakobera/heroku-buildpack-go), customized heroku-buildpack-go for gom.  

```
$ heroku create -b https://github.com/hakobera/heroku-buildpack-go
$ git push heroku master
```

After deply success, run `heroku open` like this.

```
$ heroku open
```

