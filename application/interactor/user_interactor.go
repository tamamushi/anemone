/* vim: set ts=4 sw=4: */

package interactor

import (
	"fmt"
)

type userInteractor struct {
}

func NewUserInteractor() *userInteractor {
	return &userInteractor{}
}

func (u *userInteractor) Create() error {
	fmt.Println("create !")
	return nil
}
func (u *userInteractor) Remove(id string) error {
	fmt.Println("remove !")
	return nil
}
func (u *userInteractor) Update() error {
	fmt.Println("update !")
	return nil
}
func (u *userInteractor) FindById(id string) error {
	return nil
}
func (u *userInteractor) Finds() error {
	return nil
}

/*
import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"io"
	"log"
	"os"
)

type Users struct {
	Users *[]UserData
}

type UserData struct {
	Username string   `json:"Username"`
	Email    string   `json:"Email"`
	RoleList []string `json:RoleList"`
}

func Handler() (interface{}, error) {

	sess := session.Must(session.NewSession())
	cognitoClient := cognitoidentityprovider.New(
		sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	results, err := cognitoClient.ListUsers(
		&cognitoidentityprovider.ListUsersInput{
			UserPoolId: aws.String("ap-northeast-1_QRNVIquNW"),
		})

	if err != nil {
		fmt.Println("Got error listing users")
		fmt.Println(err)
		os.Exit(1)
	}

	logFile, err := os.OpenFile("/tmp/out.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Printf("Cannot open logFile: %v", err)
	}

	users := []UserData{}

	for _, user := range results.Users {
		attributes := user.Attributes
		userdata := &UserData{}
		userdata.Username = *user.Username

		for _, a := range attributes {
			if *a.Name == "email" {
				userdata.Email = *a.Value
			} else if *a.Name == "custom:groups" {

				roles := []byte(*a.Value)
				error := json.Unmarshal(roles, &userdata.RoleList)

				if error != nil {
					fmt.Printf("Error while parsing data: %s", error)
				}
			}
		}
		users = append(users, *userdata)
	}

	defer logFile.Close()
	log.SetOutput(io.MultiWriter(logFile, os.Stderr))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	res := users
	log.Printf("return %+v", res)
	return res, nil
}

func main() {
	lambda.Start(Handler)
}
*/
