# Overview
Tapir is a tool that can efficiently enable and disable for AWS CloudWatchAlarm Actions.


## demo
![tapir](https://user-images.githubusercontent.com/6120093/111776182-09efaf00-88f5-11eb-98d9-760e2b5a802c.gif)

# Install
## Homebrew (macOS and Linux)
```
$ brew install tomozo6/tap/tapir
```

## Binary packagesk
[Releases](https://github.com/tomozo6/tapir/releases)

# Usage
```
Tapir can check the status of CloudWatchAlarm and can enable and disable alarm actions.

Usage:
  tapir [command]

Available Commands:
  disable     Disables CloudWatchAlarm actions
  enable      Enables CloudWatchAlarm actions
  help        Help about any command
  showtags    show tags for CloudWatchAlarm
  status      Displays the 'ActinsEnabled' and 'State' status of the CloudWatchAlarms

Flags:
      --config string         config file (default is $HOME/.tapir.yaml)
  -h, --help                  help for tapir
  -p, --prefix string         Alam Name Preix
  -r, --region string         Target AWS Region (default "us-east-1")
  -t, --tags stringToString   for filter. ex) project=test,env=dev (default [])

Use "tapir [command] --help" for more information about a command.
```
# Quick Start
tapir can easily manage the enable/disable for your existing CloudWatchAlarm actions.


# Configuration file
YAML format.

The default path for the configuration file is `$HOME/.tapir.yaml`.

The path can be changed with --config Flag.


```yaml
region: us-east-1
prefix: "HTTP monitoring"
tags: {project: tomozo6, env: prod}
```

# Example of execute
## status
### simple
If you don't give flags. Tapir display the status of all CloudWatchAlarms in the us-east-1 region.

```bash
$ tapir status
+-----------------------------------------------------------------------+----------------+-------------+
|                               ALARMNAME                               | ACTIONSENABLED | ALARMSSTATE |
+-----------------------------------------------------------------------+----------------+-------------+
| httpMonitoring_dev_tomozo6-ecSiteTopPage(https://dev.tomozo6.com)     | enable         | OK          |
| httpMonitoring_prod_tomozo6-ecSiteTopPage(https://www.tomozo6.com)    | enable         | OK          |
| httpMonitoring_prod_tomozo6-ecSiteTopPageOld(https://old.tomozo6.com) | enable         | ALARM       |
| httpMonitoring_stg_tomozo6-ecSiteTopPage(https://stg.tomozo6.com)     | enable         | OK          |
| resourceMonitoring_dev_tomozo6-ec2-cpuUtilizationIsHigh               | enable         | OK          |
| resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh              | enable         | OK          |
| resourceMonitoring_stg_tomozo6-ec2-cpuUtilizationIsHigh               | enable         | OK          |
+-----------------------------------------------------------------------+----------------+-------------+
```

### Use tags flag
```bash
$ tapir status -t project=tomozo6,env=prod
+-----------------------------------------------------------------------+----------------+-------------+
|                               ALARMNAME                               | ACTIONSENABLED | ALARMSSTATE |
+-----------------------------------------------------------------------+----------------+-------------+
| httpMonitoring_prod_tomozo6-ecSiteTopPageOld(https://old.tomozo6.com) | enable         | ALARM       |
| resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh              | enable         | OK          |
| httpMonitoring_prod_tomozo6-ecSiteTopPage(https://www.tomozo6.com)    | enable         | OK          |
+-----------------------------------------------------------------------+----------------+-------------+
```

### Use prefix flag
```bash
$ tapir status -p resourceMonitoring_
+----------------------------------------------------------+----------------+-------------+
|                        ALARMNAME                         | ACTIONSENABLED | ALARMSSTATE |
+----------------------------------------------------------+----------------+-------------+
| resourceMonitoring_dev_tomozo6-ec2-cpuUtilizationIsHigh  | enable         | OK          |
| resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh | enable         | OK          |
| resourceMonitoring_stg_tomozo6-ec2-cpuUtilizationIsHigh  | enable         | OK          |
+----------------------------------------------------------+----------------+-------------+
```

It's faster to be able to filter by prefix than by tag.

## disable
```bash
$ tapir status -t project=tomozo6,env=prod
+-----------------------------------------------------------------------+----------------+-------------+
|                               ALARMNAME                               | ACTIONSENABLED | ALARMSSTATE |
+-----------------------------------------------------------------------+----------------+-------------+
| httpMonitoring_prod_tomozo6-ecSiteTopPageOld(https://old.tomozo6.com) | enable         | ALARM       |
| httpMonitoring_prod_tomozo6-ecSiteTopPage(https://www.tomozo6.com)    | enable         | OK          |
| resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh              | enable         | OK          |
+-----------------------------------------------------------------------+----------------+-------------+
$ tapir disable -t project=tomozo6,env=prod
Success
$ tapir status -t project=tomozo6,env=prod
+-----------------------------------------------------------------------+----------------+-------------+
|                               ALARMNAME                               | ACTIONSENABLED | ALARMSSTATE |
+-----------------------------------------------------------------------+----------------+-------------+
| httpMonitoring_prod_tomozo6-ecSiteTopPageOld(https://old.tomozo6.com) | disable        | ALARM       |
| httpMonitoring_prod_tomozo6-ecSiteTopPage(https://www.tomozo6.com)    | disable        | OK          |
| resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh              | disable        | OK          |
+-----------------------------------------------------------------------+----------------+-------------+
```
## enable
```bash
$ tapir status -t project=tomozo6,env=prod
+-----------------------------------------------------------------------+----------------+-------------+
|                               ALARMNAME                               | ACTIONSENABLED | ALARMSSTATE |
+-----------------------------------------------------------------------+----------------+-------------+
| httpMonitoring_prod_tomozo6-ecSiteTopPageOld(https://old.tomozo6.com) | disable        | ALARM       |
| httpMonitoring_prod_tomozo6-ecSiteTopPage(https://www.tomozo6.com)    | disable        | OK          |
| resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh              | disable        | OK          |
+-----------------------------------------------------------------------+----------------+-------------+
$ tapir enable -t project=tomozo6,env=prod
Success
$ tapir status -t project=tomozo6,env=prod
+-----------------------------------------------------------------------+----------------+-------------+
|                               ALARMNAME                               | ACTIONSENABLED | ALARMSSTATE |
+-----------------------------------------------------------------------+----------------+-------------+
| httpMonitoring_prod_tomozo6-ecSiteTopPage(https://www.tomozo6.com)    | enable         | OK          |
| httpMonitoring_prod_tomozo6-ecSiteTopPageOld(https://old.tomozo6.com) | enable         | ALARM       |
| resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh              | enable         | OK          |
+-----------------------------------------------------------------------+----------------+-------------+
```

## showtags
```
$ tapir showtags
Using config file: /Users/tomohiro.b.sasaki/.tapir.yaml
Use the arrow keys to navigate: ↓ ↑ → ←
Which alarms tags do you want to show?
  👉 httpMonitoring_dev_tomozo6-ecSiteTopPage(https://dev.tomozo6.com)
    httpMonitoring_prod_tomozo6-ecSiteTopPage(https://www.tomozo6.com)
    httpMonitoring_prod_tomozo6-ecSiteTopPageOld(https://old.tomozo6.com)
    httpMonitoring_stg_tomozo6-ecSiteTopPage(https://stg.tomozo6.com)
    resourceMonitoring_dev_tomozo6-ec2-cpuUtilizationIsHigh
    resourceMonitoring_prod_tomozo6-ec2-cpuUtilizationIsHigh
    resourceMonitoring_stg_tomozo6-ec2-cpuUtilizationIsHigh

--------- Tags ----------

{created_by: terraform}

{env: dev}

{project: tomozo6}

{type: http}
```
# Licence
MIT

# Author
tomozo6