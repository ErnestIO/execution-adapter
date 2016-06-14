# Execution Adapter

execution Adapter gets input nats messages in order to create, update or delete executions, and it converts to its proper adapter messages.

## Build status

* master:  [![CircleCI](https://circleci.com/gh/ErnestIO/execution-builder/tree/master.svg?style=svg)](https://circleci.com/gh/ErnestIO/execution-builder/tree/master)
* develop: [![CircleCI](https://circleci.com/gh/ErnestIO/execution-builder/tree/develop.svg?style=svg)](https://circleci.com/gh/ErnestIO/execution-builder/tree/develop)

## Installation

```
make deps
make install
```

## Running Tests

```
make deps
go test
```

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).

