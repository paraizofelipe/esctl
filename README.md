# esctl

## Introduction

esctl is a Command Line Interface (CLI) tool designed to simplify the management of Elasticsearch clusters. By leveraging Elasticsearch's REST API, esctl provides a comprehensive set of commands to manipulate various resources effectively. Whether you're handling indices, shards, or security settings, esctl streamlines your interactions with Elasticsearch clusters.
This project aims to provide an alternative for managing Elasticsearch clusters, offering a Command Line Interface (CLI) for manipulating resources through its REST API.

## Getting Started

### Prerequisites

- Ensure you have Go installed, version 1.21 or higher. You can download it from [Go's official website](https://golang.org/dl/).
- Elasticsearch cluster up and running
- Basic understanding of Elasticsearch concepts and CLI operations

### Installation

#### Go get

1. Install `esctl` using Go: `go get github.com/paraizofelipe/esctl`
2. This command will automatically download and install `esctl` in your Go bin directory.
3. Ensure your Go bin directory is in your system's PATH to run `esctl` from any location.

#### Tarball

1. Download the Tarball: Visit the Releases section of the esctl GitHub repository and download the latest tarball for your operating system.

2. Extract the Tarball: Use a tool like tar to extract the downloaded file:

```bash
    tar -xzf esctl_<version>_<os>.tar.gz
```
Replace <version> and <os> with the corresponding version number and operating system.

3. Move the Binary: Create a directory at ~/.config/esctl/bin and move the esctl binary there:

```bash
    mkdir -p ~/.config/esctl/bin
    mv esctl ~/.config/esctl/bin
```

4. Create a Symbolic Link: Link the esctl binary to ~/.local/bin for easy execution:

```bash
    ln -s ~/.config/esctl/bin/esctl ~/.local/bin/esctl
```

Make sure ~/.local/bin is in your PATH. If not, add it to your .bashrc or .zshrc file:

```bash
    export PATH=$PATH:~/.local/bin
```

5. Verify Installation: Check if esctl is installed correctly:

```bash
    esctl --version
```

### Alternative Installation: Using get_esctl.sh Script

You can install esctl using a Bash script, which automates the download and installation process. Execute the following command:

```bash
    curl -sSL http://github.com/paraizofelipe/esctl/scripts/get_esctl.sh | bash
```

This command does the following:


1. Fetches the Script: Uses curl to download `get_esctl.sh` script from the specified URL.
2. Executes the Script: Pipes the downloaded script to bash for execution.

The script `get_esctl.sh` should be designed to:

- Download the latest esctl tarball from the GitHub releases.
- Extract the tarball to a temporary directory.
- Move the esctl binary to ~/.config/esctl/bin.
- Create a symbolic link in ~/.local/bin.
- Clean up any temporary files created during the process.

Please ensure you have the required permissions to execute scripts from the internet and that you trust the source of the script.



## Commands

### get

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
### change

```bash
 $ esctl -n local change alias --body '{}'
 $ esctl -n local change mapping --body '{}'
 $ esctl -n local change security ???
 $ esctl -n local change config 'CLUSTER_NAME'
```

### describe

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

### create

```bash
$ esctl -n 'local' create index --body '{}' 'INDEX_NAME' 
$ esctl -n 'local' create index doc --id '1q2w3e4r' --body '{}' 'INDEX_NAME'
```

### delete

```bash
$ esctl -n 'local' delete index 'INDEX_NAME'
$ esctl -n 'local' delete index alias  --name 'ALIAS_NAME' 'INDEX_PATTERN'
$ esctl -n 'local' delete security user 'USERNAME'
```

### task

```bash
$ esctl -n 'local' task list 
$ esctl -n 'local' task 
$ esctl -n 'local' task cancel
```

### search

```bash
$ esctl -n 'local' search --query '{}'
```

### apply

Apply cluster reroute:

```bash
$ esctl -n 'local' apply -f reroute.json
```

File `users.json`:

```json
{
  "kind": "ClusterReroute",
  "body": {
    "commands": [
      {
        "allocate_replica": {
          "index": "index1",
          "shard": 0,
          "node": "es-node-01"
        }
      }
    ]
  }
}

```

Apply new users:

```bash
$ esctl -n 'local' apply -f users.json
```

File `users.json`:

```json
{
  "kind": "SecurityUser",
  "body": [
    {
      "username": "test",
      "full_name": "User Test",
      "email": "es.test@esctl.com",
      "password": "icecream123",
      "roles": [
        "admin"
      ],
      "metadata": {}
    }
  ]
}
```

Apply actions to index aliases:

```bash
$ esctl -n 'local' apply -f alias.json
```

File `alias.json`

```json
{
  "kind": "IndexAlias",
  "body": {
    "actions": [
      {
        "add": {
          "index": "index1",
          "alias": "alias1"
        }
      },
      {
        "add": {
          "index": "index2",
          "alias": "alias1"
        }
      }
    ]
  }
}
```

Apply mapping in index:
```bash
$ esctl -n 'local' apply -f mapping.json
```

File `mapping.json`:

```json
{
  "kind": "IndexMapping",
  "index": [
    "index1"
  ],
  "body": {
    "properties": {
      "name": {
        "type": "keyword"
      },
      "description": {
        "type": "text"
      },
      "age": {
        "type": "integer"
      }
    }
  }
}
```
