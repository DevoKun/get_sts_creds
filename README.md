get_sts_creds
=============

Generate AWS STS credentials using **sts::assumeRole** and write those credentials to a config file for use with **Amazon CodeDeploy**



## Purpose

* The purpose of **get_sts_creds** is almost identical to that of Amazon's own [aws-codedeploy-session-helper](https://github.com/awslabs/aws-codedeploy-samples/tree/master/utilities/aws-codedeploy-session-helper) tool.
* The predominent motivation for the creation of this **golang** based app is portability, while **aws-codedeploy-session-helper** requires that Ruby and the AWS-SDK be installed.
* While I [Devon] am a fan of Ruby, the versioning setup for the AWS-SDK means that when you may have coded something for v2 of the SDK, if v3 gets installed, everything will be broken.




## Usage

* All configuration options are taken via environment variables.
* The use of environment variables is not always ideal, but it can simplify deployments in certain environments.



### Environment Variables

#### AWS_DEFAULT_REGION

- **name:** AWS_DEFAULT_REGION

- **alternate name:** AWS_REGION
- **default:** us-east-1



#### AWS_ACCESS_KEY_ID

- **name:** AWS_ACCESS_KEY_ID



#### AWS_SECRET_ACCESS_KEY

- **name:** AWS_SECRET_ACCESS_KEY



#### AWS_ROLE_ARN

- **name:** AWS_ROLE_ARN
- **example:** arn:aws:iam::xxxxxxxxxx:role/codedeploy-service



#### AWS_ROLE_SESSIONNAME

- **name:** AWS_ROLE_SESSIONNAME
- **default:** codedeploysess



#### AWS_EXTERNAL_ID

- **name:** AWS_EXTERNAL_ID
- **default:** get_sts_creds



#### AWS_STS_CREDENTIALS_FILE

- **name:** AWS_STS_CREDENTIALS_FILE
- **default:** /etc/codedeploy-agent/conf/awsStsCredentials.ini





### Call via Cron

* The credentials don't expire for 3 hours, so you do not need to regenerate the credentials too rapidly.
* The easiest method of acquiring the credentials is using a cron job

```shell
0 * * * * COMMAND (AWS_DEFAULT_REGION="us-east-1" AWS_REGION="us-east-1" AWS_ACCESS_KEY_ID="AKIxxxxxxxxxxx" AWS_SECRET_ACCESS_KEY="xxxxxxxxxx" AWS_ROLE_ARN="arn:aws:iam::xxxxxxxxxx:role/codedeploy-service" AWS_ROLE_SESSIONNAME="codedeploysess" AWS_EXTERNAL_ID="getStsCreds" AWS_STS_CREDENTIALS_FILE="awsStsCredentials.ini" get_sts_creds 1>/dev/null 2>&1)
```





### Examples

#### Example 1

```shell

export AWS_DEFAULT_REGION="us-east-1"
export AWS_REGION="us-east-1"
export AWS_ACCESS_KEY_ID="AKIxxxxxxxxxxx"
export AWS_SECRET_ACCESS_KEY="xxxxxxxxxx"
export AWS_ROLE_ARN="arn:aws:iam::xxxxxxxxxx:role/codedeploy-service"
export AWS_ROLE_SESSIONNAME="codedeploysess"
export AWS_EXTERNAL_ID="getStsCreds"
export AWS_STS_CREDENTIALS_FILE="awsStsCredentials.ini"

get_sts_creds

```



#### Example 2

```shell

AWS_DEFAULT_REGION="us-east-1" \
AWS_REGION="us-east-1" \
AWS_ACCESS_KEY_ID="AKIxxxxxxxxxxx" \
AWS_SECRET_ACCESS_KEY="xxxxxxxxxx" \
AWS_ROLE_ARN="arn:aws:iam::xxxxxxxxxx:role/codedeploy-service" \
AWS_ROLE_SESSIONNAME="codedeploysess" \
AWS_EXTERNAL_ID="getStsCreds" \
AWS_STS_CREDENTIALS_FILE="awsStsCredentials.ini" \
get_sts_creds

```



#### Example 3

```shell
AWS_DEFAULT_REGION="us-east-1" AWS_REGION="us-east-1" AWS_ACCESS_KEY_ID="AKIxxxxxxxxxxx" AWS_SECRET_ACCESS_KEY="xxxxxxxxxx" AWS_ROLE_ARN="arn:aws:iam::xxxxxxxxxx:role/codedeploy-service" AWS_ROLE_SESSIONNAME="codedeploysess" AWS_EXTERNAL_ID="getStsCreds" AWS_STS_CREDENTIALS_FILE="awsStsCredentials.ini" get_sts_creds
```



#### Example of the Geneated Credential File

```ini
[default]
aws_access_key_id     = ASIxxxxxxxxxxxxxx
aws_secret_access_key = xxxxxxxxxxxxxxxxx
aws_session_token     = 2RKPTFF4...HUBQ==
region                = us-east-1
output                = json

# these credentials expire 2018-02-23 21:37:02 +0000 UTC

```







