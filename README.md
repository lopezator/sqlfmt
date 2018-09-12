# sqlfmt

sqlfmt formats CockroachDB SQL statements from stdin to line length n.

PostgreSQL and maybe other SQL should work good also.

## Installation

sqlfmt uses internal codebase of [cockroachdb 2.1 release](https://github.com/cockroachdb/cockroach/tree/release-2.1) but decouples from the root project so you can use it in a handy way on command line and CI without having to install/build the entire cockroachdb project.

Just execute the following command and you are good to go:

```sh
$ go get github.com/lopezator/sqlfmt/...
```

## Usage

```
Usage:
  sqlfmt [flags]

Flags:
      --align           Align the output.
  -h, --help            help for sqlfmt
      --len int         The line length where sqlfmt will try to wrap. (default 60)
      --no-simplify     Don't simplify output.
      --tab-width int   Number of spaces per indentation level. (default 4)
      --use-spaces      Indent with spaces instead of tabs.
```

## Examples

1. SQL Input from file, output to stdout:

    ```
    sqlfmt < path/to/my/sql/input.sql
    ```

2. SQL Input from file, output to file:

    ```
    sqlfmt < path/to/my/sql/input.sql > path/to/my/sql/output.sql
    ```

3. SQL Input from stdin string, output to file:

    ```
    sqlfmt <<< "select * from    mytable where id=7 group by name" > output.sql
    ```

4. SQL Input from stdin string, output to stdout:

    ```
    sqlfmt <<< "select * from    mytable where id=7 group by name"
    ```

    prints:

    ```
    SELECT * FROM mytable WHERE id = 7 GROUP BY name
    ```

5. sqlfmt also validates input validity, so you can use as a checker of your .sql files in CI:

    Mispell of from (fron):

    ```
    sqlfmt <<< "select * fron    mytable where id=7 group by name"
    ```

    Would print:

    ```
    Error: syntax error at or near "fron"
    ```

## Build cross-platform binaries from source

Each release comes with pre-compiled binaries for several platforms:

https://github.com/lopezator/sqlfmt/releases

Anyway, if you want to cross-compile from source, you can do it using docker in a handy way:

1. Build container:
    ```
    $> docker run --rm -t -d --name="sqlfmt-builder" -v="$PWD:/go/src/github.com/lopezator/sqlfmt" --workdir="/go/src/github.com/lopezator/sqlfmt" --entrypoint="cat" cockroachdb/builder:20180813-101406
    ```

2. Build binaries inside the container:
    ```
    $> docker exec -it sqlfmt-builder make build
    ```

3. Stop container:
    ```
    $> docker container stop sqlfmt-builder
    ```