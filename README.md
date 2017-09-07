# awsconsole

**awsconsole** is a simple application which allows you to generate a sign-in URL for the AWS Console with your IAM Access Keys and STS service. It uses `~/.aws/credentials` and environment variables as a source for the access keys. The default duration of the sign-in is 3600 seconds.

![Showtime](https://media.giphy.com/media/3o6Ztafy3u2XyXeYOQ/giphy.gif)

## Installation
### Using go tools
`$ go get -u github.com/tenesys/awsconsole`
### Using binary packages
Just download a tarball suitable for you architecture from [gobuilder.me](https://gobuilder.me/github.com/tenesys/awsconsole)

## Usage
`$ awsconsole [-v] [profile-name]`

- -v - print URL instead of opening variable
- -d - session duration (eg. 8h, 30m)
- profile-name - use profile from credentials file instead of environment variables


**Do not change positions of parameters, `-v` MUST BE before the profile name**
### Environment variables
```
$ env | grep AWS_PROFILE
AWS_PROFILE=tenesys
$ awsconsole
*opens browser*
$ awsconsole -v
https://aws.amazon.com/...
```

### Profile name
It will use profiles defined in `~/.aws/credentials`  
```
$ awsconsole -v tenesys
```

## Configuration
### Session duration
You can change your default session duration by exporting an environment variable
`AWSCONSOLE_DURATION`. It takes the same value as `-d` flag.

## Required IAM permissions
To use awsconsole, you should have permission to call some IAM and STS actions. If you are encountering `Could not get user information` or `Could not connect to STS service` errors, make sure you have required permissions to call `IAM:GetUser` and `STS:GetFederationToken`.

### Policy
You can add the policy below to your IAM users or groups.
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Stmt1472645401000",
            "Effect": "Allow",
            "Action": [
                "iam:GetUser"
            ],
            "Resource": [
                "arn:aws:iam::ACCOUNT-ID-WITHOUT-HYPHNES:user/${aws:username}"
            ]
        },
        {
            "Sid": "Stmt1472645461000",
            "Effect": "Allow",
            "Action": [
                "sts:GetFederationToken"
            ],
            "Resource": [
                "arn:aws:sts::ACCOUNT-ID-WITHOUT-HYPHNES:federated-user/${aws:username}"
            ]
        }
    ]
}
```
Please note that you have to replace **ACCOUNT-ID-WITHOUT-HYPHENS** with your Account Id.

## License
GPLV3

## Maintainers
- Jakub Woźniak \<j.wozniak@tenesys.pl\>

