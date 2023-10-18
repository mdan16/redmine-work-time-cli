# redmine-work-time-cli

## Usage

```yaml:mywork.yaml
apiHost: http://localhost:8080
redmineApiKey: YOUR_REDMINE_API_KEY
spentOn: 2023-10-18
timeEntries:
  - issueId: 1
    hours: 1
    activityId: 1
    comments: comment1
  - issueId: 2
    hours: 1
    activityId: 1
    comments: comment2
```

```
./redmine-work-time-cli -f mywork.yaml
```

## Help

```
NAME:
   redmine-work-time-cli - A new cli application

USAGE:
   redmine-work-time-cli [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value
   --help, -h              show help
```
