# esctl

This project aims to provide an alternative for managing Elasticsearch clusters, offering a Command Line Interface (CLI) for manipulating resources through its REST API.

# Getting Started

TODO

# Commands

## get

```bash
 $ esctl -n local get indices
 $ esctl -n local get aliases
 $ esctl -n local get nodes
 $ esctl -n local get shards
 $ esctl -n local get thread-pool
 $ esctl -n local get pending-tasks
 $ esctl -n local get tasks
 $ esctl -n local get health
 $ esctl -n local get repositories
 $ esctl -n local get snapshots
 $ esctl -n local get config
```
## change

```bash
 $ esctl -n local change alias --body '{}'
 $ esctl -n local change mapping --body '{}'
 $ esctl -n local change security ???
 $ esctl -n local change config 'CLUSTER_NAME'
```

## describe

```bash
$ esctl -n 'local' describe index 
$ esctl -n 'local' describe index doc --id '1234567' --fields 'name,age' 'INDEX_NAME'
$ esctl -n 'local' describe index alias 'INDEX_NAME_1,INDEX_NAME_2'
$ esctl -n 'local' describe index stats 'INDEX_NAME_1,INDEX_NAME_2'
$ esctl -n 'local' describe index mapping 'INDEX_NAME_1,INDEX_NAME_2'
$ esctl -n 'local' describe index settings 'INDEX_NAME_1,INDEX_NAME_2'
$ esctl -n 'local' describe task --id '1q2w3e4r'
$ esctl -n 'local' describe count 'INDEX_NAME'
$ esctl -n 'local' describe security user -n 'USERNAME'
```

## create

```bash
$ esctl -n 'local' create index --body '{}' 'INDEX_NAME' 
$ esctl -n 'local' create index doc --id '1q2w3e4r' --body '{}' 'INDEX_NAME'
```

## delete

```bash
$ esctl -n 'local' delete index 'INDEX_NAME'
$ esctl -n 'local' delete index alias  --name 'ALIAS_NAME' 'INDEX_PATTERN'
$ esctl -n 'local' delete security user 'USERNAME'
```

## task

```bash
$ esctl -n 'local' task list 
$ esctl -n 'local' task 
$ esctl -n 'local' task cancel
```

## search

```bash
$ esctl -n 'local' search --query '{}'
```

## apply

```bash
$ esctl -n 'local' apply -f 'CONFIG_FILE'
```
