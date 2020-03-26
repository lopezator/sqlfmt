# sqlfmt

sqlfmt formats CockroachDB SQL statements from stdin to line length n.

PostgreSQL and maybe other SQL should work good also.

## Installation

sqlfmt uses internal codebase of [cockroachdb](https://github.com/cockroachdb/cockroach) but decouples from the root 
project so you can use it in a handy way on command line and CI without having to install/build the entire cockroachdb 
project.

You can just 
[download the precompiled binary for your platform](https://github.com/lopezator/sqlfmt/releases), or build it by your own by executing `make build`.

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

## Sqlfmt (bash kung-fu) usage examples for dev & CI

The Makefile shipped with the project includes two targets useful for dev & CI time.

Dependencies or this targets are [bash](https://www.gnu.org/software/bash/) & [moreutils](https://joeyh.name/code/moreutils/) for your platform of preference, instalable via apt, brew, apk, yum or whatever.

1. [Target](Makefile#72) to format all `*.sql` files on local project (excluding vendor folder):

    ```
    $> make sql-fmt
    ```

    Will fail, because of error on `examples/sql_incorrect.sql`

    ```
    syntax error at or near "crete" sql-check failed, run "make sql-fmt" and try again
    make: *** [sql-check] Error 1
    ```

    Correcting the typo on `examples/sql_incorrect.sql` and running `make sql-fmt` again will format `examples/*.sql` contents.

    This target will check syntax and **will overwrite** all `*.sql` files (excluding vendor folder) applying format, so is **ideal for dev usage** before committing.

2. [Target](Makefile#78) to check if all `*.sql` files on local project have the correct syntax and format (excluding vendor folder):

    ```
    $> make sql-check
    ```

    This target will check syntax and format correctness but **will not change anything**, so is **ideal for a CI check**.
    Introducing any format changes (add or remove line breaks or spaces, for example) into any of the well formatted (after point 1) `examples/*.sql` will result into a format error:

    ```
    sql-check failed, run "make sql-fmt" and try again
    make: *** [sql-check] Error 1
    ```

    Error is self-descriptive, running `make sql-fmt` again will make `make sql-check` pass.
