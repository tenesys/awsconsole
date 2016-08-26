# awsconsole

**awsconsole** is a simple application which allows you to generate URL for AWS Console by using your IAM Access Keys. It uses `~/.aws/credentials` and environment variables as source for access keys.

## Installation
### Using go tools
`$ go get -u github.com/tenesys/awsconsole`
### Using binary package
Just download tarball from *Downloads* section and copy it to `/usr/bin/`.

## Usage
`$ awsconsole [-v] [profile-name]`

- -v - print URL instead of opening variable
- profile-name - use profile from credentials file instead of environment variables

**Do not change positions of parameters, `-v` MUST BE before profile name***
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

## License
GPLV3

## Maintainers
- Jakub Woźniak \<j.wozniak@tenesys.pl\>

