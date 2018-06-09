# Dapper

A collection of [Docker](https://docker.com) contained application

# Usage

Dapper can start any of the supported applications that are listed in the usage information found by running dapper with no parameters:

```bash
 $ ./dapper
```


You can then start any one of these apps and pass any switches you would like to that app like so:

```bash
 $ ./dapper oni --help
```

# Upgrading

Docker images are built on demand from Dockerfile's, if there has been an update to any of them you may need to manually rebuild the containers by executing the following command inside the dapper repository:

```bash
 $ docker-compose build
```

This is also done for you when you run dapper without any parameters.

```bash
 $ ./dapper
```
