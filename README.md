## Examples of concurrency patterns with golang

This project contains several examples you can use in production.
Feel free to clone, fork, join, copy and reuse in a DRY spirit.
each pattern has a description given as a comment on top of the package declaration.

You can execute any example. This example runner is powered by cobra & viper. see below:

### To build:
```
cd goconpa && go build
```
### To run:

```
goconpa run 
```
to get available options. For example to run the Cond broadcast example:

```
goconpa run broadcast
```