//
// Package
//
package main


//
// Import
//
import (
  "fmt"
  "os"
  "github.com/fatih/color"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/sts"
) // import



//
// Custom Types
//





//
// Global Variables
//







//
// Functions
//

func getEnvVar(varname1 string, varname2 string, vardef string) string {

  dots := "....................................."

  val0 := "xxxx"
  val1 := ""
  val2 := ""

  val1 = os.Getenv(varname1)
  fmt.Println(color.WhiteString("  +") + 
              color.GreenString(" v1- ") + 
              color.YellowString(varname1) + 
              color.WhiteString( dots[0:(30 - len(varname1) )] ) +
              color.WhiteString(": ") +
              color.CyanString(val1))
  if (("x"+val1) != "x") {
    val0 = val1
  } // if

  if (("x"+varname2) != "x") {
    val2 = os.Getenv(varname2)
    fmt.Println(color.WhiteString("  +") + 
                color.GreenString(" v2- ") + 
                color.YellowString(varname2) + 
                color.WhiteString( dots[0:(30 - len(varname2) )] ) +
                color.WhiteString(": ") + 
                color.CyanString(val2))
    if (("x"+val2) != "x") {
      val0 = val2
    } // if
  } // if


  if (val0 == "xxxx") {
    val0 = vardef
    fmt.Println(color.WhiteString("  +") + 
                color.GreenString(" vd- ") + 
                color.YellowString(varname1) + 
                color.WhiteString( dots[0:(30 - len(varname1) )] ) +
                color.WhiteString(": ") + 
                color.CyanString(vardef))
  } // if


  if (val0 == "") {
    fmt.Println("\n" + color.WhiteString("  !!!! ") + color.RedString("Missing ") + color.YellowString(varname1) + color.RedString(" environment variable") + color.WhiteString(" !!!! ") + "\n")
    os.Exit(1)
  } // if

  return val0

} // func getEnvVar



func main() {

  fmt.Println(color.BlueString("-------------------------------------------------------------------------------"))

  awsRegion          := getEnvVar("AWS_DEFAULT_REGION",    "AWS_REGION", "us-east-1")
  awsAccessKey       := getEnvVar("AWS_ACCESS_KEY_ID",     "",           "")
  awsSecretAccessKey := getEnvVar("AWS_SECRET_ACCESS_KEY", "",           "")
  awsRoleArn         := getEnvVar("AWS_ROLE_ARN",          "",           "arn:aws:iam:::role/codedeploy-service")
  awsRoleSessionName := getEnvVar("AWS_ROLE_SESSIONNAME",  "",           "codedeploysess")
  awsExternalId      := getEnvVar("AWS_EXTERNAL_ID",       "",           "get_sts_creds")

  stsCredentialsFilename := getEnvVar("AWS_STS_CREDENTIALS_FILE", "", "/etc/codedeploy-agent/conf/awsStsCredentials.ini")

  fmt.Println(color.BlueString("-------------------------------------------------------------------------------"))
  fmt.Println(color.WhiteString("  +") + color.YellowString(" AWS Region") + color.WhiteString(".......................: ") + color.CyanString(awsRegion))
  fmt.Println(color.WhiteString("  +") + color.YellowString(" AWS Access Key") + color.WhiteString("...................: ") + color.CyanString(awsAccessKey))
  fmt.Println(color.WhiteString("  +") + color.YellowString(" AWS Secret Access Key") + color.WhiteString("............: ") + color.CyanString(awsSecretAccessKey))
  fmt.Println(color.BlueString("-------------------------------------------------------------------------------"))


  creds := credentials.NewStaticCredentials(
    awsAccessKey,
    awsSecretAccessKey,
    "",
  )

  sess, _ := session.NewSession(&aws.Config{
    Region:      aws.String(awsRegion),
    Credentials: creds,
  })

  stsService := sts.New(sess)

  stsAssumeRoleParams := &sts.AssumeRoleInput{
    RoleArn:         aws.String(awsRoleArn),
    RoleSessionName: aws.String(awsRoleSessionName),
    ExternalId:      aws.String(awsExternalId),
  } // AssumeRoleInput

  resp, err := stsService.AssumeRole(stsAssumeRoleParams)
  if (err != nil) {
    fmt.Println(err)
  } // if err




  stsAccessKeyId     := *resp.Credentials.AccessKeyId
  stsSecretAccessKey := *resp.Credentials.SecretAccessKey
  stsSessionToken    := *resp.Credentials.SessionToken
  stsExpiration      := resp.Credentials.Expiration.String()

  fmt.Println(color.WhiteString("  +") + color.YellowString(" STS Access Key ID") + color.WhiteString("................: ") + color.CyanString(stsAccessKeyId))
  fmt.Println(color.WhiteString("  +") + color.YellowString(" STS Secret Access Key") + color.WhiteString("............: ") + color.CyanString(stsSecretAccessKey))
  fmt.Println(color.WhiteString("  +") + color.YellowString(" STS Session Token") + color.WhiteString("................: ") + color.CyanString(stsSessionToken))
  fmt.Println(color.WhiteString("  +") + color.YellowString(" STS Expiration") + color.WhiteString("...................: ") + color.CyanString(stsExpiration))
  fmt.Println(color.BlueString("-------------------------------------------------------------------------------"))



  file, err := os.Create(stsCredentialsFilename)
  if err != nil {
    fmt.Println(color.WhiteString("  !!!! ") + color.RedString("Cannot create credentials file file") + color.WhiteString(" !!!! "), err)
  } // if
  defer file.Close()

  fmt.Fprintf(file, "[default]" + "\n")
  fmt.Fprintf(file, "aws_access_key_id     = "    + stsAccessKeyId     + "\n")
  fmt.Fprintf(file, "aws_secret_access_key = "    + stsSecretAccessKey + "\n")
  fmt.Fprintf(file, "aws_session_token     = "    + stsSessionToken    + "\n")
  fmt.Fprintf(file, "region                = "    + awsRegion          + "\n")
  fmt.Fprintf(file, "output                = "    + "json"             + "\n")
  fmt.Fprintf(file, "\n")
  fmt.Fprintf(file, "# these credentials expire " + stsExpiration      + "\n")
  fmt.Fprintf(file, "\n")

  fmt.Println(color.WhiteString("  + ") +
              color.YellowString("AWS STS ") + 
              color.GreenString("Credentials written to ") + 
              color.WhiteString(stsCredentialsFilename))
  fmt.Println(color.BlueString("-------------------------------------------------------------------------------"))



} // func main



