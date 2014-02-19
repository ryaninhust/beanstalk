# Beanstalk
Go client for [beanstalkd](http://kr.github.com/beanstalkd/).

 * Add beanstalkd connection pool base on origin from KR
 * Lazy connection for avoiding connection error occurring 

## Install

    $ go get github.com/ryaninhust/beanstalk

## Use

Produce jobs:

    pool := beanstalk.NewPool(["localhost:11300", "remotehost:11300"])
    id, err := pool.Put([]byte("hello"), 1, 0, 120*time.Second)

Consume jobs:
    TODO

