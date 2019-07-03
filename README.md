🔥🔥🔥 yakiire 🔥🔥🔥
================================================================================

[![CircleCI](https://circleci.com/gh/mookjp/yakiire.svg?style=svg&circle-token=07d36e051e436463f6dac42c402f664e4be4db3a)](https://circleci.com/gh/mookjp/yakiire)

`yakiire` (yaki-ire; 焼入れ) is a CLI to manage and operate data on GCP [Firestore](https://firebase.google.com/docs/firestore).

**THIS IS THE ALPHA VERSION !!**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Get](#get)
  - [Query](#query)
- [TODOs](#todos)
  - [Set](#set)
- [For development](#for-development)
  - [Run tests](#run-tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation

```bash
go get github.com/mookjp/yakiire
```

## Configuration

`yakiire` needs environment variables:

| ENV | value | required |
|-----|-------|----------|
| YAKIIRE_FIRESTORE_PROJECT_ID | Firestore project ID | Yes |
| YAKIIRE_GOOGLE_APPLICATION_CREDENTIALS | GCP's credential file path | No |

If `YAKIIRE_GOOGLE_APPLICATION_CREDENTIALS` was not set, `yakiire` uses `GOOGLE_APPLICATION_CREDENTIALS` to access to Firestore.


## Usage

### Get

```bash
yakiire get -c <collection name> <document ID>
```

e.g.

```bash
$ yakiire get -c products 002VQIDE4D

# it shows a doc in JSON format

{"Attributes":{"color":"red","size":"100"},"CategoryIDs":["1","2","3"],"ID":"002VQIDE4D","Name":"Test Product"}
```

It is handy to use [jq](https://firebase.google.com/docs/firestore) to check the result from the command.

```bash
$ yakiire get -c products 002VQIDE4D | tail -n 1 | jq .

{
  "Attributes": {
    "color": "red",
    "size": "100"
  },
  "CategoryIDs": [
    "1",
    "2",
    "3"
  ],
  "ID": "002VQIDE4D",
  "Name": "Test Product"
}
```

### Query

```bash
yakiire query --collection products \
    --where '{"Path": "Attributes.size", "Op": ">", "Value": 0}' \
    --where '{"Path": "Attributes.color", "Op": "==", "Value": "red"}' \
    --limit 1
```

e.g.

```bash
yakiire query --collection products \
    --where '{"Path": "Attributes.size", "Op": ">", "Value": 0}'

# it shows docs in line-delimited JSON format

{"Attributes":{"color":"red","size":100},"CategoryIDs":["1","2","3"],"ID":"1","Name":"Test Product"}
{"Attributes":{"color":"red","size":200},"CategoryIDs":["1","2","3"],"ID":"2","Name":"Another Test Product"}

yakiire query --collection products \
    --where '{"Path": "Attributes.size", "Op": ">", "Value": 0}' \
    --limit 1

# limit to 1 result
# default number of limit is 20

{"Attributes":{"color":"red","size":100},"CategoryIDs":["1","2","3"],"ID":"1","Name":"Test Product"}

yakiire query --collection products \
    --where '{"Path": "Attributes.size", "Op": ">", "Value": 0}' \
    --where '{"Path": "CategoryIDs", "Op": "array-contains", "Value": "1"}' \
    --limit 1

# multiple where conditions

{"Attributes":{"color":"red","size":100},"CategoryIDs":["1","2","3"],"ID":"1","Name":"Test Product"}
```

## TODOs

### Set

```bash
yakiire set -c <collection name> <query>
```


## For development

### Run tests

```bash
FIRESTORE_EMULATOR_HOST=localhost:8080 make test
```

Test needs running Firestore emulator and it can be run with `docker-compose`.
in `Makefile`, `test` will start firestore emulator container before it starts tests.
