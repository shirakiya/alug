# alug
![](https://github.com/shirakiya/alug/workflows/Go/badge.svg?branch=master) ![](https://github.com/shirakiya/alug/workflows/Release/badge.svg)  

AWS console Login URL Generator "alug".  
It creates and shows URL attached federation token to login AWS console.


## Usage
alug requires `profile` value. This `profile` is **the name in your AWS config file you want to login as**.  
You can pass profile value in below ways.

- Command line option
- Environment variable

### Command line option
```
$ alug -p profile_you_want_to_login
https://signin.aws.amazon.com/federation?Action=login&Destination=https%3A%2F%2Fconsole.aws.amazon.com%2F&Issuer=IssuedByAlug&SigninToken=XXXXXXXXXX...
```

### Environment variable
```
$ export ALUG_PROFILE=profile_you_want_to_login
$ alug
https://signin.aws.amazon.com/federation?Action=login&Destination=https%3A%2F%2Fconsole.aws.amazon.com%2F&Issuer=IssuedByAlug&SigninToken=XXXXXXXXXX...
```


If you are required to input MFA Code for your IAM role, alug also requires its token code.
(This behavior is changed by configure.)

```
$ alug
Input MFA Code: XXXXXX
https://signin.aws.amazon.com/federation?Action=login&Destination=https%3A%2F%2Fconsole.aws.amazon.com%2F&Issuer=IssuedByAlug&SigninToken=XXXXXXXXXX...
```


## Description
It is easy to switch role in AWS console in your browser, but session duration time is 1 hour
and we cannot customize it. I was frustrated with behavior.  
However [console login URL with Federation Token](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html)
enables to expand this session duration time. This is the reason why I made alug.


## Installation
### macOS
You can use Homebrew.

```
$ brew install shirakiya/alug/alug
```

Or built binaries available. See Linux section.

### Linux
#### Binary
Download built binary from [GitHub release](https://github.com/shirakiya/alug/releases) and locate it to a directory in your path.

```
$ wget https://github.com/shirakiya/alug/releases/download/0.0.1/alug_0.0.1_darwin_amd64.tar.gz
$ tar xvzf alug_0.0.1_darwin_amd64.tar.gz
$ mv alug /to/our/path  # (i.e. /usr/local/bin)
```

#### go get
You can also install by go get.

```
$ go get github.com/shirakiya/alug/...
```


## Configuration
alug requires some variable in `~/.aws/config`.  
Show the sample config below included variables related to alug.
(Full settings in `~/.aws/config` is described in the [official article](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) )

```
[profile foo]
role_arn = arn:aws:iam::000000000000:role/role_name
duration_seconds = 43200
mfa_serial = arn:aws:iam::123456789012:mfa/user
role_session_name = bar
source_profile = default
```

### role_arn (required)
The ARN of an IAM role that you want to switch to.

### duration_seconds
(optional, default: `3600`)  

The maximum duration time of the role session.
Specifies the maximum duration of the role session, in seconds. The value can range
from 900 seconds (15 minutes) up to the maximum session duration setting for the role
(which can be a maximum of 43200).

### mfa_serial
(optional)  

The identification number of an MFA device to use when assuming a role. **This is mandatory
only if the trust policy of the role being assumed includes a condition that requires MFA
authentication**. The value can be either a serial number for a hardware device (such as `GAHT12345678`)
or an Amazon Resource Name (ARN) for a virtual MFA device (such as `arn:aws:iam::123456789012:mfa/user`).

### role_session_name
(optional, default: `assume-by-alug-{timestamp}`)  

The name to attach to the role session.

### source_profile
(optional, default: `default`)  

A profile name which is used for getting the Federation token from AWS. alug lookups
`~/.aws/credentials`.  
Or environment variables `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` are available instead.
