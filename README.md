# Overview
Tapir is a tool that can efficiently enable and disable for AWS CloudWatchAlarm Actions.


## demo
![tapir](https://user-images.githubusercontent.com/6120093/111757203-6d221700-88de-11eb-9ac2-367f8cbe7aed.gif)


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
  status      Displays the 'ActinsEnabled' and 'State' status of the CloudWatchAlarms

Flags:
      --config string   config file (default is $HOME/.tapir.yaml)
  -h, --help            help for tapir

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

# Licence
MIT

# Author
tomozo6