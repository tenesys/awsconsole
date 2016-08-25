package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "net/url"
    "io/ioutil"
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
    console = "https://console.aws.amazon.com/"
)

func main() {
    userSession := session.New()
    stsSvc := sts.New(userSession)

    iamSvc := iam.New(userSession)
    user, _ := iamSvc.GetUser(nil)

    tokenOutput, _ := stsSvc.GetFederationToken(&sts.GetFederationTokenInput{
        Name: aws.String(fmt.Sprintf("%s-awsconsole", *user.User.UserName)),
        DurationSeconds: aws.Int64(duration),
        Policy: aws.String(policy),
    })

    jsonSignin, _ := json.Marshal(struct {
        SessionID string `json:"sessionId"`
        SessionKey string `json:"sessionKey"`
        SessionToken string `json:"sessionToken"`
    }{
        SessionID: *tokenOutput.Credentials.AccessKeyId,
        SessionKey: *tokenOutput.Credentials.SecretAccessKey,
        SessionToken: *tokenOutput.Credentials.SessionToken,
    })


    response, _ := http.Get(awsfed + "?Action=getSigninToken&Session=" +
        url.QueryEscape(string(jsonSignin)))
    responseBody, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()

    var data map[string]interface{}
    json.Unmarshal(responseBody, &data)

    url := (awsfed + "?Action=login&Destination=" +
        url.QueryEscape(console) + "&SigninToken=" +
        data["SigninToken"].(string))

    browser.OpenURL(url)
}


/* vim: set ts=8 sw=4 tw=100 et :*/
