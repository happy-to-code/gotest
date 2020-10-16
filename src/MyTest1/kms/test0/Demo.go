package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"strings"
)

var createKey = `
{
	"keySpec":"%s",
	"password":"%s",
	"pubFormat":"%s"
}
`
var getPubKey = `
{
	"keyId":"%s"
}
`

type CreateKeyRes struct {
	Status int    `json:"status"`
	Data   KeyMsg `json:"data"`
}
type KeyMsg struct {
	KeyId        string `json:"keyId"`
	KeySpec      string `json:"keySpec"`
	Origin       string `json:"origin"`
	CreationDate string `json:"creationDate"`
	PublicKey    string `json:"publicKey"`
}
type GetPubKeyRes struct {
	Status int `json:"status"`
	Data   Pub `json:"data"`
}
type Pub struct {
	PublicKey string `json:"publicKey"`
	X         string `json:"x"`
	Y         string `json:"y"`
	Curve     string `json:"curve"`
}

func main() {
	body := fmt.Sprintf(createKey, "sm2", "123", "sm2_pem")

	fmt.Printf("----%+v\n", body)

	keyMsg, _ := CreateKey("123456")
	fmt.Printf("keyMsg:%+v\n", keyMsg)
	fmt.Println(keyMsg.KeyId, "<===>", keyMsg.PublicKey)

}

func CreateKey(pwd string) (*KeyMsg, error) {
	body := fmt.Sprintf(createKey, "sm2", pwd, "sm2_pem")
	client := NewClient()
	code, res, err := client.Post("http://10.bb.json.3.250:8080/v1/createKey", body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if code != 200 {
		log.Error("post code:", code)
		reply := new(error)
		err = json.Unmarshal(res, reply)
		if err != nil {
			log.Error("json:", err)
			return nil, err
		}
		log.Error("reply:", reply)
		return nil, *reply
	}
	reply := new(CreateKeyRes)
	err = json.Unmarshal(res, reply)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &reply.Data, nil
}

type Client interface {
	Post(url, body string) (int, []byte, error)
	Get(url string) (int, []byte, error)
}
type client struct {
	client *http.Client
}

type TYPE int

const (
	POST TYPE = iota
	GET
)

func NewClient() Client {
	client := new(client)
	client.client = &http.Client{}
	return client
}

func (c *client) Post(url, body string) (int, []byte, error) {
	request, err := c.request(POST, url, body)
	if err != nil {
		return 0, nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	return c.do(request)
}

func (c *client) Get(url string) (int, []byte, error) {
	request, err := c.request(GET, url, "")
	if err != nil {
		return 0, nil, err
	}

	return c.do(request)
}

func (c *client) request(typ TYPE, url, body string) (*http.Request, error) {
	switch typ {
	case GET:
		return http.NewRequest("GET", url, nil)
	case POST:
		return http.NewRequest("POST", url, strings.NewReader(body))
	}
	return nil, errors.New("error TYPE!")
}

func (c *client) do(req *http.Request) (int, []byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, b, nil
}
