# async

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

