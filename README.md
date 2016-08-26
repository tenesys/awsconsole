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
- profile-name - use profile from credentials file instead of environment variables

**Do not change positions of parameters, `-v` MUST BE before the profile name***
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

