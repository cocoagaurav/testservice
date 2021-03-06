package firebase

import (
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

var (
	app    *firebase.App
	client *auth.Client
	//FireAuthStr string
)

const fireAuthStr = `{"type": "service_account","project_id": "testproject-fa267","private_key_id": "341dae30dee88fcffaa2febac93b9a6d4911bdbc","private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCHYLLQGld5IiBN\ncRkaL28b5wW6s3Avp0kLNGnLrdWsWI4FrK06tDXul41MNHRwuWRZoHgKbtmvJhKG\n0EyM/2UsRDMg4kgwMpF+ljePJQxkaxX5oipnCvEp+QuLNhPnea6d9jSlKUm+qGSz\n+5Ei9wBGVeHUwcPMl9pMGFEqKTEbslv85iBiFPBOoCaBqRDFgz32nil3Vw07rebs\nXC3Y6EpleVtO6enbiY7expCYYCcDLHd9bF0LJ0Vm5GShXuzm0PmsrnedVDeFjSsu\n2klWPcF/XqRxc/PJHt67ujkODMKuue6HlWRO3Pw37bp++hD3rA4L+WgjwOJgx8SU\nVg3YYdPxAgMBAAECggEAG4/td/8U9h2jlADFypYDuhuUCAoGej1F2tkl/Qj8auVw\nrOkWL9CG9ne2leBMILMuIi1Qo1ckTMukk/wOydopoBSWkEhhyCZThwFQeH1jg4Jl\n6/g/R9Frfk8tMX+mF8enbJO27jV0xTOFpTs/tm2xiaBspSN6GMqF8F43EC1oySmA\nTAg9XiynIvBbgYK4fil46QXiH2podRZSG7gNHPAxbddcqvanruLAKk2JZ22nBLNh\nQJ+eiRcanKPpZi9c0E7zwNa1i+0b0MiJ7r5hASfkuC+GFd6EfMFFtZ/NtxtTQLr5\nOr9R+YvYYJmV+uLDhjudmbtt9+37o+rFtu7GiY58nQKBgQC+Z+ULfbjCZxeRMPeH\nlNhwajMoRcs8UXfI6f8Nv0raIqu2EmvepE6V+8fFkUcSiRFrGWDb7gHwvFuad0/S\n+5xllexA0HlIawxpQfq1J6s4Qk8G4pCwmltKZ7JeMdGAeIcZEed1gMEWNKe3ya2g\nA6EhLRBDz3M9pvWVektlvopxfwKBgQC2A9ASmddFdB6l+KTpUI4eSxL54L5MD5cU\neRWQgvW+uz0HtlZHH5+GiV0UQv7woW3z0Df0mbpfvc5qTwkABhvYruU9j4sHZSsd\nw4COEgtuFrQN4pqxnNHr8/6U3wM/XwDHMMpy3ZOQR+I2bMLK1xaTDdHzoMyWPgZD\nGWsTsK6SjwKBgHEVbWA8w92ZstKFfY2lpkJloIp7oS/qxrSp3NRCV2dkjgztteke\nNpo3VjeNh+OHSrQL943HNpnOlK0RzXPmAcYHm7AG4PFUuqNND2RF8hfQsfTJ3Ns2\nYZ+4JKRy/BVMABiwnIIZ/RN+JFowSpEtdqYoiG9tpujn3xVu85ay6rBrAoGAUSkc\nE99DbXXc4LchmePQq1Ngj8mWMUZWYMupQPoUaEsHaLP2ftpsAMqplYpWMahZ5fj3\nqnsN7vks3JyHb9pJenJqR+wE23RSKIBvh2ombJ11BigAQKijtmnjIDDdOtm6+Bca\nfuOslA5poUkYBuin6USlVNRjxa68jhj8dRg4j6MCgYEAq6jkjmJmUSCVsoHtBe0w\n6PP20D8L6NNLP8KXT5ACiSkrJ0NcgJn+KFkGfPI5dL/sFNvHI3z9W0MPMlI6xIj5\ngmVwcQBTVgA/V6uI36vq3vTr6Ed+psqjVaS2tbwa63g94nWo2PQ/wKfhveS74DVq\nPCx9stbbEoE/3odxFKvuOXM=\n-----END PRIVATE KEY-----\n","client_email": "firebase-adminsdk-6b9tl@testproject-fa267.iam.gserviceaccount.com","client_id": "100554064165740094092","auth_uri": "https://accounts.google.com/o/oauth2/auth", "token_uri": "https://oauth2.googleapis.com/token", "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-6b9tl%40testproject-fa267.iam.gserviceaccount.com" }`

func FirebaseStartAuth(env model.Env) {
	var err error
	conf := &firebase.Config{ServiceAccountID: env.FirebaseServiceId}
	firefile := createFireBaseJsonFile(fireAuthStr)

	opt := option.WithCredentialsFile(firefile)

	app, err = firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v ", err)
	}
	fmt.Println("firebase is ready to serve")
}

func CreateFireBaseUser(user *model.User) (*auth.UserRecord, error) {
	var err error
	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v ", err)
	}

	params := (&auth.UserToCreate{}).
		DisplayName(user.Name).
		Password(user.Password).
		Email(user.EmailId).
		EmailVerified(false)
	u, err := client.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("error creating user: %v\n", err)
		return nil, err
	}
	return u, err
}

func GenerateToken(uid string) string {
	client, _ = app.Auth(context.Background())
	token, _ := client.CustomToken(context.Background(), uid)
	return token

}
func VerifyToken(token string) *auth.Token {
	fmt.Printf("\n varifying token is:%v", token)

	//client, _ = app.Auth(context.Background())

	tok, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		fmt.Printf(" \n err is:%v", err)
		return nil
	}
	fmt.Printf("\n return token is:%T", tok)

	fmt.Printf("\n varified token is:%v", tok)

	fmt.Println("label 9")

	return tok

}

func GetUserCreds(authId string) *auth.UserRecord {
	user, err := client.GetUser(context.Background(), authId)
	if err != nil {
		fmt.Printf("error is:%s", err)
		return nil
	}
	return user
}

func DeleteFirebaseUser(uid string) {
	err := client.DeleteUser(context.Background(), uid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	log.Printf("Successfully deleted user: %s\n", uid)

}

func createFireBaseJsonFile(authStr string) string {
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		glog.Fatalf("Failed to get working directory: %+v", err)
	}

	fileName := path.Join(currentWorkingDir, "agnus-firebase.json")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fileObject, err := os.Create(fileName)
		if err != nil {
			glog.Fatalf("Failed to create firebase json file: %+v", err)
		}
		defer fileObject.Close()

		_, err = io.Copy(fileObject, strings.NewReader(authStr))
		if err != nil {
			glog.Fatalf("Failed to input json into file: %+v", err)
		}
	}

	return fileName
}
