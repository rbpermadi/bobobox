# bobobox

## Description

Bobobox is a repository for Bobobox Test Assignment. Its a web-service app using to search hotel room availability. It's using **Go** as its programming language.

## Onboarding and Development Guide

### Prequisites

* [**Git**](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* [**Go (1.9.7 or later)**](https://golang.org/doc/install)
* [**Go Dep 0.5 or later**](https://golang.github.io/dep/docs/installation.html)
* [**MySQL**](https://www.mysql.com/downloads/)

### Setup

- Please install/clone the [Prequisites](#prequisites) and make sure it works perfectly on your local machine.

- Clone this repo in your local at $GOPATH/src/github.com/rbpermadi If you have not set your GOPATH, set it using [this](https://golang.org/doc/code.html#GOPATH) guide. If you don't have directory src, github.com, or rbpermadi in your GOPATH, please make them.

    ```
    cd $GOPATH/src/github.com/rbpermadi
    git clone git@github.com:rbpermadi/bobobox.git
    ```

- Go to Bobobox directory and sync the vendor file

    ```
    cd $GOPATH/src/github.com/rbpermadi/bobobox
    make dep
    ```

- Copy and edit(optional) `.env.sample`

    ```
    cp env.sample .env
    ```

- To prepare database, you can use your own mysql_client to create new database, after that you import db/bobobox_db_schema.sql to your database.

- (optional) after import db schema, you can also import bobobox_db_data_sample.sql to your database.

Finally, run **Bobobox** in your local machines.

```
> make run
```

To kill the server you just need to hold `Ctrl + C`

### Contributing

1. Make new branch with descriptive name about your change(s) and checkout to that branch

   ````
   git checkout -b branch_name
   ````


2. Commit and push your change to upstream

   ````
   git commit -m "message"
   git push [remote_name] [branch_name]
   ````

3. Open pull request in `Github`

4. Ask someone to review your code.

5. If your code is approved, the pull request can be merged.

### Test

You can test the project using test/Bobobox.postman_collection.json in your postman

## FAQ

> Not available yet
