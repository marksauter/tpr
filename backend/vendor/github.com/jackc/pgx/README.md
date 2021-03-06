[![](https://godoc.org/github.com/jackc/pgx?status.svg)](https://godoc.org/github.com/jackc/pgx)

# Pgx

## Version 3 Beta Branch

This is the `v3` branch which is currently in beta. General release is planned
for July. `v2` is the current release branch. `v3` is considered to be stable in
the sense of lack of known bugs, but the API is not considered stable until
general release. No further changes are planned, but the beta process may
surface desirable changes. If possible API changes are acceptable, then `v3` is
the recommented branch for new development.

Pgx is a pure Go database connection library designed specifically for
PostgreSQL. Pgx is different from other drivers such as
[pq](http://godoc.org/github.com/lib/pq) because, while it can operate as a
database/sql compatible driver, pgx is primarily intended to be used directly.
It offers a native interface similar to database/sql that offers better
performance and more features.

## Features

Pgx supports many additional features beyond what is available through database/sql.

* Listen / notify
* Transaction isolation level control
* Full TLS connection control
* Binary format support for custom types (can be much faster)
* Copy protocol support for faster bulk data loads
* Logging support
* Configurable connection pool with after connect hooks to do arbitrary connection setup
* PostgreSQL array to Go slice mapping for integers, floats, and strings
* Hstore support
* JSON and JSONB support
* Maps inet and cidr PostgreSQL types to net.IPNet and net.IP
* Large object support
* Null mapping to Null* struct or pointer to pointer.
* Supports database/sql.Scanner and database/sql/driver.Valuer interfaces for custom types
* Logical replication connections, including receiving WAL and sending standby status updates
* Notice response handling (this is different than listen / notify)
* Batch queries
* Single-round trip query mode
* pgtype package includes support for approximately 60 different PostgreSQL types - these are usable in pgx native and any database/sql PostgreSQL adapter

## Performance

Pgx performs roughly equivalent to [pq](http://godoc.org/github.com/lib/pq) and
[go-pg](https://github.com/go-pg/pg) for selecting a single column from a single
row, but it is substantially faster when selecting multiple entire rows (6893
queries/sec for pgx vs. 3968 queries/sec for pq -- 73% faster).

See this [gist](https://gist.github.com/jackc/d282f39e088b495fba3e) for the
underlying benchmark results or checkout
[go_db_bench](https://github.com/jackc/go_db_bench) to run tests for yourself.

## database/sql

Import the ```github.com/jackc/pgx/stdlib``` package to use pgx as a driver for
database/sql. It is possible to retrieve a pgx connection from database/sql on
demand. This allows using the database/sql interface in most places, but using
pgx directly when more performance or PostgreSQL specific features are needed.

## Documentation

pgx includes extensive documentation in the godoc format. It is viewable online at [godoc.org](https://godoc.org/github.com/jackc/pgx).

## Testing

pgx supports multiple connection and authentication types. Setting up a test
environment that can test all of them can be cumbersome. In particular,
Windows cannot test Unix domain socket connections. Because of this pgx will
skip tests for connection types that are not configured.

### Normal Test Environment

To setup the normal test environment, first install these dependencies:

    go get github.com/jackc/fake
    go get github.com/shopspring/decimal
    go get gopkg.in/inconshreveable/log15.v2

Then run the following SQL:

    create user pgx_md5 password 'secret';
    create user " tricky, ' } "" \ test user " password 'secret';
    create database pgx_test;
    create user pgx_replication with replication password 'secret';

Connect to database pgx_test and run:

    create extension hstore;

Next open conn_config_test.go.example and make a copy without the
.example. If your PostgreSQL server is accepting connections on 127.0.0.1,
then you are done.

### Connection and Authentication Test Environment

Complete the normal test environment setup and also do the following.

Run the following SQL:

    create user pgx_none;
    create user pgx_pw password 'secret';

Add the following to your pg_hba.conf:

If you are developing on Unix with domain socket connections:

    local  pgx_test  pgx_none  trust
    local  pgx_test  pgx_pw    password
    local  pgx_test  pgx_md5   md5

If you are developing on Windows with TCP connections:

    host  pgx_test  pgx_none  127.0.0.1/32 trust
    host  pgx_test  pgx_pw    127.0.0.1/32 password
    host  pgx_test  pgx_md5   127.0.0.1/32 md5

### Replication Test Environment

Add a replication user:

    create user pgx_replication with replication password 'secret';

Add a replication line to your pg_hba.conf:

    host replication pgx_replication 127.0.0.1/32 md5

Change the following settings in your postgresql.conf:

    wal_level=logical
    max_wal_senders=5
    max_replication_slots=5

Set `replicationConnConfig` appropriately in `conn_config_test.go`.

## Version Policy

pgx follows semantic versioning for the documented public API on stable releases. Branch `v2` is the latest stable release. `master` can contain new features or behavior that will change or be removed before being merged to the stable `v2` branch (in practice, this occurs very rarely).
