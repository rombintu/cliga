### CLI_LIGA
Output of the --help command
```bash
NAME:
   cliga - CLI for LIGA agent

USAGE:
   cliga [global options] command [command options]

COMMANDS:
   get      Get [something]
   check    Check [number sprint]
   ping     Simple ping server
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug     (default: false)
   --help, -h  show help
```

### Usage examples
Ping a private server that checks sprints. Checking availability
```bash
$ cliga ping
[Agent]: PING!
[Agent]: Waiting...
[Server]: PONG!
```

Get the conditions and the result of the tasks of a certain sprint
```bash
$ cliga get sprint 99
[Agent]: Sprint [99] not found
[Agent]: Exit
```
Check the completion of tasks. If everything is TRUE, then send the result to the server
```bash
$ cliga check --user iivanov 1
[Agent]: Username: iivanov [123456789]
[Agent]: All the tasks in the Sprint [1] are solved!
[Agent]: Waiting...
[Server]: Updated!
[Agent]: Exit
```
Where the digit (1) is the sprint number and the first letter is the first letter of your first name, the rest is your last name  
**Example**: Ivan Ivanov - iivanov  
**Attention**: remember it, it will be your login