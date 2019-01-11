package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sendpulse/api/customlogger"
	"strconv"
	"sync"
	"time"

	"sendpulse/api/client/responsestructs"
)

const (
	apiUrl = "https://api.sendpulse.com/"
	authPath = "oauth/access_token"
	adressbooksPath = "addressbooks"
	campaignsPath = "campaigns"
	campaignInfoPath = "campaigns/"
	reconnectInterval = 3
)

type Logger interface {
	Log(message string)
	Alert(err error)
}

func New(clientId string, clientSecret string, logger Logger) *Client{
	c := Client{
		grantType:"client_credentials",
		clientId:clientId,
		clientSecret:clientSecret,
		wg:&sync.WaitGroup{},
	}
	if logger!=nil{
		c.logger = logger
	}else{
		c.logger = &customlogger.DefaultLogger{}
	}
	return &c
}

type Client struct{
	grantType		string
	clientId		string
	clientSecret	string
	actualToken		string
	tokenExpiresIn	time.Time
	wg				*sync.WaitGroup
	logger			Logger
}

func (c *Client) updateToken() error{
	params := make(map[string]string)
	params["grant_type"] = "client_credentials"
	params["client_id"] = c.clientId
	params["client_secret"] = c.clientSecret
	data := []byte(getRawParams(params))
	r := bytes.NewReader(data)
	res, err := http.Post(apiUrl + authPath,"application/x-www-form-urlencoded",r)
	if err!=nil{
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err!=nil{
		return err
	}

	if res.StatusCode!=200{
		return errors.New("error during token getting: " + string(body))
	}
	atr := responsestructs.AccessToken{}
	err = json.Unmarshal(body,&atr)
	if err!=nil{
		return err
	}
	c.actualToken = atr.AccessToken
	c.tokenExpiresIn = time.Now()
	if atr.ExpiresIn>100{
		c.tokenExpiresIn = time.Now().Add(time.Second*time.Duration(atr.ExpiresIn-100))
	}
	return nil
}

func (c *Client) ensureToken() {
	c.wg.Add(1)
	defer c.wg.Done()

	for {
		err := c.updateToken()
		if err==nil{
			return
		}
		c.logger.Alert(err)
		time.Sleep(time.Duration(reconnectInterval)*time.Second)
	}

}

func (c *Client) sendRequest(method string, path string, body []byte) (*http.Response, error){
	c.wg.Wait()
	if len(c.actualToken)==0||c.tokenExpiresIn.Before(time.Now()){
		c.ensureToken()
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, apiUrl + path,bytes.NewReader(body))
	if err!=nil{
		return nil, err
	}
	req.Header.Add("Authorization","Bearer " + c.actualToken)
	resp, err := client.Do(req)
	if err!=nil{
		return nil, err
	}
	return resp, nil
}

func (c *Client) sendGet(path string, params map[string]string) (*http.Response, error){
	if params!=nil&&len(params)>0{
		path = path + "?" + getRawParams(params)
	}
	return c.sendRequest("GET",path,nil)
}

func (c *Client) GetAddressBooks() (*responsestructs.AddressBooks, error){
	var books responsestructs.AddressBooks
	_, err := c.GetEntity(&books,adressbooksPath)
	if err != nil{
		return nil, err
	}
	return &books, nil
}

func (c *Client) GetCampaigns() (*responsestructs.Campaigns, error) {
	var campaigns responsestructs.Campaigns
	_, err := c.GetEntity(&campaigns,campaignsPath)
	if err != nil{
		return nil, err
	}
	return &campaigns, nil
}

func (c *Client) GetCampaignInfo(campaign *responsestructs.Campaign) error {
	var tcampaign responsestructs.Campaign
	_, err := c.GetEntity(&tcampaign,campaignInfoPath + strconv.Itoa(campaign.Id))
	if err != nil{
		return err
	}
	*campaign = tcampaign
	return nil
}

func (c *Client) GetEntity(dataStruct interface{}, path string) (interface{}, error){
	resp, err := c.sendGet(path,nil)
	if err!=nil{
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil, err
	}


	err = json.Unmarshal(body,dataStruct)
	if err !=nil{
		return nil, err
	}
	return dataStruct, nil
}

func getRawParams(params map[string]string) string{
	result := ""
	for k, v := range params{
		if len(result)>0{
			result = result + "&"
		}
		result = result + k + "=" + v
	}
	return result
}

