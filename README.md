# async

Simple golang worker sync pool that matches your best number of goroutines with the available CPUs. This optimizes the number of go routines so that you always have the best performance.

This lib does:
1. create go routines according to your number of available CPUs;
2. sync the work to be done by them;
3. stop work with the given context;

## Why

1. Managing worker loads is repetitive work;
2. Worker code always leads to managing go routines;
3. This provides a standard way of doing that in a efficient manner;
4. Less code for you to manage;
5. No external dependencies;

## How to use

1. Start your app
2. Define your work function
3. Setup the `async` lib
4. Send data to be worked on your function
5. Manage your errors


### Full example

```go

import (
    "github.com/arxdsilva/async"
)

func main() {
    ... // 1. start your app
    workFn := myFunction // 2.define your work function

    dataChan, errChan := async.New(
        ctx, workFn, 
        async.WithChanSizeData(1000),
        async.WithChanSizeErr(2)) // 3. setup the async lib

    mw := myWorker{
        DC: dataChan, // 4. Setup your worker to send data to the async execution
        EC: errChan, // 5. Setup your worker to manage generated errors
    }

    mw.Start()
}
```