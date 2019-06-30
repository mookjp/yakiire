🔥🔥🔥 yakiire 🔥🔥🔥
================================================================================

[![CircleCI](https://circleci.com/gh/mookjp/yakiire/tree/master.svg?style=svg)](https://circleci.com/gh/mookjp/yakiire/tree/master)

`yakiire` (yaki-ire; 焼入れ) is a CLI to manage and operate data on GCP [Firestore](https://firebase.google.com/docs/firestore).

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Contents

- [Configuration](#configuration)
  - [`gcp`](#gcp)
- [Usage](#usage)
  - [Get](#get)
- [TODOs](#todos)
  - [Set](#set)
  - [Query](#query)
- [For development](#for-development)
  - [Run tests](#run-tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Configuration

`yakiire` needs configuration file.
It sees `${HOME}`/.yakiire.yaml by default.

You can set the file path by `--config` flag.

The yaml file is like this:

```yaml
gcp:
  credentials: "credentials/cred.json"
```

### `gcp`

| key | value |
|-----|-------|
| credentials | GCP's credential file path |


## Usage

### Get

```bash
yakiire get -c <collection name> <document ID>
```

e.g.

```bash
$ yakiire --config .yakiire.yaml get -c products 002VQIDE4D

# it shows a doc in JSON format
```

It is handy to use [jq](https://firebase.google.com/docs/firestore) to check the result from the command.

```bash
$ yakiire --config .yakiire.yaml get -c products 002VQIDE4D | tail -n 1 | jq .
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
