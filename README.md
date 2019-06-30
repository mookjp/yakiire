🔥🔥🔥 yakiire 🔥🔥🔥
================================================================================

[![CircleCI](https://circleci.com/gh/mookjp/yakiire.svg?style=svg&circle-token=07d36e051e436463f6dac42c402f664e4be4db3a)](https://circleci.com/gh/mookjp/yakiire)

`yakiire` (yaki-ire; 焼入れ) is a CLI to manage and operate data on GCP [Firestore](https://firebase.google.com/docs/firestore).

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Get](#get)
- [TODOs](#todos)
  - [Set](#set)
  - [Query](#query)
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

If `YAKIIRE_GOOGLE_APPLICATION_CREDENTIALS` was not set, `yakiire` uses `YAKIIRE_GOOGLE_APPLICATION_CREDENTIALS` to access to Firestore.


## Usage

### Get

```bash
yakiire get -c <collection name> <document ID>
```

e.g.

```bash
$ yakiire --config .yakiire.yaml get -c products 002VQIDE4D

# it shows a doc in JSON format

{"Attributes":{"color":"red","size":"100"},"CategoryIDs":["1","2","3"],"ID":"002VQIDE4D","Name":"Test Product"}
```

It is handy to use [jq](https://firebase.google.com/docs/firestore) to check the result from the command.

```bash
$ yakiire --config .yakiire.yaml get -c products 002VQIDE4D | tail -n 1 | jq .

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

## TODOs

### Set

```bash
yakiire set -c <collection name> <query>
```

### Query

```bash
yakiire query -c <collection name> <query>
```


## For development

### Run tests

```bash
make test
```

Test needs running Firestore emulator and it can be run with `docker-compose`.
