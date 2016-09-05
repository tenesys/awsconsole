package main

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"
    "net/url"
    "io/ioutil"
    "flag"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sts"
    "github.com/aws/aws-sdk-go/service/iam"
    "github.com/pkg/browser"
)

const (
    policy = `{"Version": "2012-10-17", "Statement": [{"Effect": "Allow", "Action": ["*"], "Resource": ["*"]}]}`
    duration = int64(3600)
    awsfed = "https://signin.aws.amazon.com/federation"
    console = "https://console.aws.amazon.com/console/home?region="
)

var (
    version = "0.1.0"
)

func ParseArgs() (bool, bool, string) {
    verbose := flag.Bool("v", false, "print URL instead of opening browser")
    printVersion := flag.Bool("V", false, "print version")
    flag.Parse()
    profile := ""
    if len(flag.Args()) > 0 {
        profile = flag.Arg(0)
    }

    return *verbose, *printVersion, profile
}

func GetSession(profile string) *session.Session {
    if profile == "" {
        profile = os.Getenv("AWS_PROFILE")
    }
    userSession, err := session.NewSessionWithOptions(session.Options{
        Profile: profile,
        SharedConfigState: session.SharedConfigEnable,
    })
    if err != nil {
        fmt.Println("Could not get profile", profile)
        fmt.Println(err)
        os.Exit(7)
    }
    return userSession
}

func PrepareBrowser() {
    browser.Stderr = ioutil.Discard
    browser.Stdout = ioutil.Discard
}

func main() {
    verbose, printVersion, profile := ParseArgs()

    if printVersion {
        fmt.Printf("awsconsole v%v\n", version)
        fmt.Println("https://github.com/tenesys/awsconsole")
        os.Exit(0)
    }

    userSession := GetSession(profile)
    region := *userSession.Config.Region
    stsSvc := sts.New(userSession)

    iamSvc := iam.New(userSession)
    user, err := iamSvc.GetUser(nil)

    if err != nil {
        fmt.Println("Could not get user information")
        fmt.Println(err)
        os.Exit(1)
    }

    tokenOutput, err := stsSvc.GetFederationToken(&sts.GetFederationTokenInput{
        Name: aws.String(fmt.Sprintf("%s-awsconsole", *user.User.UserName)),
        DurationSeconds: aws.Int64(duration),
        Policy: aws.String(policy),
    })

    if err != nil {
        fmt.Println("Could not connect to STS Service:")
        fmt.Println(err)
        os.Exit(2)
    }

    jsonSignin, err := json.Marshal(struct {
        SessionID string `json:"sessionId"`
        SessionKey string `json:"sessionKey"`
        SessionToken string `json:"sessionToken"`
    }{
        SessionID: *tokenOutput.Credentials.AccessKeyId,
        SessionKey: *tokenOutput.Credentials.SecretAccessKey,
        SessionToken: *tokenOutput.Credentials.SessionToken,
    })

    if err != nil {
        fmt.Println("Could not generate token input")
        os.Exit(3)
    }

    response, err := http.Get(awsfed + "?Action=getSigninToken&Session=" +
        url.QueryEscape(string(jsonSignin)))
    if err != nil {
        fmt.Println("Could not get service response")
        os.Exit(4)
    }

    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Could not read service response")
        os.Exit(5)
    }
    response.Body.Close()

    var data map[string]interface{}
    err = json.Unmarshal(responseBody, &data)
    if err != nil {
        fmt.Println("Could not parse json")
        os.Exit(6)
    }

    url := (awsfed + "?Action=login&Destination=" +
        url.QueryEscape(console + region) + "&SigninToken=" +
        data["SigninToken"].(string))

    if verbose == true {
        fmt.Println(url)
    } else {
        PrepareBrowser()
        browser.OpenURL(url)
    }
}


/* vim: set ts=8 sw=4 tw=100 et :*/
