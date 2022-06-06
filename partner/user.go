package partner

import (
	"encoding/json"
	"fmt"
	"telrobot/models"
	app "telrobot/util/common"
	"telrobot/util/http"
)

var token = app.Env("token", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjUxOWVmMmI0NTIwZDJhMzU3ZDNmZTdkZjE1Y2RmOGQ3NDJmOGNjNGViODVlNWI1ZWRmYjcyYjYyODg2NDU2NmYwOWFiZjM2OGRkNjViYjYwIn0.eyJhdWQiOiIxIiwianRpIjoiNTE5ZWYyYjQ1MjBkMmEzNTdkM2ZlN2RmMTVjZGY4ZDc0MmY4Y2M0ZWI4NWU1YjVlZGZiNzJiNjI4ODY0NTY2ZjA5YWJmMzY4ZGQ2NWJiNjAiLCJpYXQiOjE2NTM3MDE2ODQsIm5iZiI6MTY1MzcwMTY4NCwiZXhwIjoxNjg1MjM3Njg0LCJzdWIiOiJiNzM0NTYyMy0xMGRjLTRjNjktOGFhNy1jODdmMThmODljMzciLCJzY29wZXMiOltdfQ.cLqXS2KEOEvfdYOGFMeCqtphHb_JLubwjI1oKYxFgS0bIeMRolTSSWi1BE0HWO5DRmGAxrDOBS4iGqWtsb1FdUhbQnLYIe1aa9jOwGHm6kR_GJHHF3ETai8UGWqaTjnuiC74IsexyIlgr3Qj7w4kDuJuO60SnYcJTrxsIRRh736P71tux61PZrokM9UjAFqrD9Pn7kIHcsCuMGdExQm8MmFkDGC9VIMgfWrdtEAyZd-eEhTpqLlyA-ov_LCQeFiZHhfm18zWH_KFmvYj1ftqv_o8mQ5GcF8Q_43umiPZXH70rEaQutp79KNmt0g08_-Y3TwG2PY57P9WB0_boeRAeycoCWwUk_F0CmulM0eyBs-OIyI6HIxXQOQ2O53OYrPglGhoBklVa6Y7_ie3n94xu_I2x1QiDk6Uf_0NqH2V8CIUu_i1v-e-vQmYWAL_K9TGDOPBiq_J2soyuA7qHfd1kswF2O2nYQdJ6TlS2Fe-WYtN1PiJ3M9UPgKGm8JHekib-XONmhBCl8cdEhV0SXNPFhXx49CgD65CT8tVDB5VRYPWqPLztoRdkZ5LmrTv2XbpNSMBwT4CwXhaeLMPKxtl2HHTHmvPYMCZGaib3TAPdkXkpHKLAMNb91XMu9u8lz1Fq5PRlx5GssCdU_GpOCgww7gyWatFgvzArvaJf_dg60k")

type Client struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

func NewClient() *Client {
	url := "http://agent-api.ai.telrobot.top"
	return &Client{Url: url, Token: token}
}

func (c *Client) GetList() {
	url := c.Url + "/users"
	data, err := http.Get(url, "ContentType: application/json", token)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	resp := models.ListUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	fmt.Println(resp)
}

func (c *Client) Get(id string) models.GetUserResp {
	//id := "83e64bbe-01c1-437d-b67f-6184419d1db2"
	url := c.Url + "/users/" + id + "/edit"

	data, err := http.Get(url, "ContentType: application/json", token)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	resp := models.GetUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	fmt.Println(resp)

	return resp
}

func (c *Client) Create(model models.User) (string, error) {
	body := map[string]string{
	}
	body["name"] = model.Name
	body["email"] = model.Email
	body["phone"] = model.Phone
	body["password"] = model.Password
	body["password_confirmation"] = model.PasswordConfirmation

	bodybyte, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	url := c.Url + "/users"

	res, err := http.PostWithJson(url, bodybyte, token)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	resp := models.CreateUserResp{}

	err = json.Unmarshal(res, &resp)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	fmt.Println(string(res))
	fmt.Println(resp)

	return resp.Data.Id, nil
}

func (c *Client) Update(id string, model models.User) error {
	//id := "83e64bbe-01c1-437d-b67f-6184419d1db2"
	url := c.Url + "/users/" + id

	body := map[string]string{
	}
	body["name"] = model.Name
	body["email"] = model.Email
	body["phone"] = model.Phone
	body["password"] = model.Password
	body["password_confirmation"] = model.PasswordConfirmation

	bodybyte, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	data, err := http.Put(url, bodybyte, "ContentType: application/json", token)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	resp := models.UpdateUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	fmt.Println(resp)

	return nil
}

func (c *Client) Remove(id string) error {
	//id := "83e64bbe-01c1-437d-b67f-6184419d1db2"
	url := c.Url + "/users/" + id

	data, err := http.Delete(url, nil, "ContentType: application/json", token)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	resp := models.DeleteUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	fmt.Println(resp)

	return nil
}
