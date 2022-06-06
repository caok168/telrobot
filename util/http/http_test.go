package http

import (
	"encoding/json"
	"fmt"
	"telrobot/models"
	"testing"
)

func TestGetList(t *testing.T) {
	url := "http://agent-api.ai.telrobot.top/users"

	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjUxOWVmMmI0NTIwZDJhMzU3ZDNmZTdkZjE1Y2RmOGQ3NDJmOGNjNGViODVlNWI1ZWRmYjcyYjYyODg2NDU2NmYwOWFiZjM2OGRkNjViYjYwIn0.eyJhdWQiOiIxIiwianRpIjoiNTE5ZWYyYjQ1MjBkMmEzNTdkM2ZlN2RmMTVjZGY4ZDc0MmY4Y2M0ZWI4NWU1YjVlZGZiNzJiNjI4ODY0NTY2ZjA5YWJmMzY4ZGQ2NWJiNjAiLCJpYXQiOjE2NTM3MDE2ODQsIm5iZiI6MTY1MzcwMTY4NCwiZXhwIjoxNjg1MjM3Njg0LCJzdWIiOiJiNzM0NTYyMy0xMGRjLTRjNjktOGFhNy1jODdmMThmODljMzciLCJzY29wZXMiOltdfQ.cLqXS2KEOEvfdYOGFMeCqtphHb_JLubwjI1oKYxFgS0bIeMRolTSSWi1BE0HWO5DRmGAxrDOBS4iGqWtsb1FdUhbQnLYIe1aa9jOwGHm6kR_GJHHF3ETai8UGWqaTjnuiC74IsexyIlgr3Qj7w4kDuJuO60SnYcJTrxsIRRh736P71tux61PZrokM9UjAFqrD9Pn7kIHcsCuMGdExQm8MmFkDGC9VIMgfWrdtEAyZd-eEhTpqLlyA-ov_LCQeFiZHhfm18zWH_KFmvYj1ftqv_o8mQ5GcF8Q_43umiPZXH70rEaQutp79KNmt0g08_-Y3TwG2PY57P9WB0_boeRAeycoCWwUk_F0CmulM0eyBs-OIyI6HIxXQOQ2O53OYrPglGhoBklVa6Y7_ie3n94xu_I2x1QiDk6Uf_0NqH2V8CIUu_i1v-e-vQmYWAL_K9TGDOPBiq_J2soyuA7qHfd1kswF2O2nYQdJ6TlS2Fe-WYtN1PiJ3M9UPgKGm8JHekib-XONmhBCl8cdEhV0SXNPFhXx49CgD65CT8tVDB5VRYPWqPLztoRdkZ5LmrTv2XbpNSMBwT4CwXhaeLMPKxtl2HHTHmvPYMCZGaib3TAPdkXkpHKLAMNb91XMu9u8lz1Fq5PRlx5GssCdU_GpOCgww7gyWatFgvzArvaJf_dg60k"
	data, err := Get(url, "ContentType: application/json", token)
	if err != nil {
		t.Error(err.Error())
	}

	resp := models.ListUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(resp)
}

func TestGetById(t *testing.T) {
	//aed884ce-1aee-42cc-b19c-062e27a8f9f2
	id := "83e64bbe-01c1-437d-b67f-6184419d1db2"
	url := "http://agent-api.ai.telrobot.top/users/" + id + "/edit"

	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjUxOWVmMmI0NTIwZDJhMzU3ZDNmZTdkZjE1Y2RmOGQ3NDJmOGNjNGViODVlNWI1ZWRmYjcyYjYyODg2NDU2NmYwOWFiZjM2OGRkNjViYjYwIn0.eyJhdWQiOiIxIiwianRpIjoiNTE5ZWYyYjQ1MjBkMmEzNTdkM2ZlN2RmMTVjZGY4ZDc0MmY4Y2M0ZWI4NWU1YjVlZGZiNzJiNjI4ODY0NTY2ZjA5YWJmMzY4ZGQ2NWJiNjAiLCJpYXQiOjE2NTM3MDE2ODQsIm5iZiI6MTY1MzcwMTY4NCwiZXhwIjoxNjg1MjM3Njg0LCJzdWIiOiJiNzM0NTYyMy0xMGRjLTRjNjktOGFhNy1jODdmMThmODljMzciLCJzY29wZXMiOltdfQ.cLqXS2KEOEvfdYOGFMeCqtphHb_JLubwjI1oKYxFgS0bIeMRolTSSWi1BE0HWO5DRmGAxrDOBS4iGqWtsb1FdUhbQnLYIe1aa9jOwGHm6kR_GJHHF3ETai8UGWqaTjnuiC74IsexyIlgr3Qj7w4kDuJuO60SnYcJTrxsIRRh736P71tux61PZrokM9UjAFqrD9Pn7kIHcsCuMGdExQm8MmFkDGC9VIMgfWrdtEAyZd-eEhTpqLlyA-ov_LCQeFiZHhfm18zWH_KFmvYj1ftqv_o8mQ5GcF8Q_43umiPZXH70rEaQutp79KNmt0g08_-Y3TwG2PY57P9WB0_boeRAeycoCWwUk_F0CmulM0eyBs-OIyI6HIxXQOQ2O53OYrPglGhoBklVa6Y7_ie3n94xu_I2x1QiDk6Uf_0NqH2V8CIUu_i1v-e-vQmYWAL_K9TGDOPBiq_J2soyuA7qHfd1kswF2O2nYQdJ6TlS2Fe-WYtN1PiJ3M9UPgKGm8JHekib-XONmhBCl8cdEhV0SXNPFhXx49CgD65CT8tVDB5VRYPWqPLztoRdkZ5LmrTv2XbpNSMBwT4CwXhaeLMPKxtl2HHTHmvPYMCZGaib3TAPdkXkpHKLAMNb91XMu9u8lz1Fq5PRlx5GssCdU_GpOCgww7gyWatFgvzArvaJf_dg60k"
	data, err := Get(url, "ContentType: application/json", token)
	if err != nil {
		t.Error(err.Error())
	}

	resp := models.GetUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(resp)
}

func TestUpdateById(t *testing.T) {
	//aed884ce-1aee-42cc-b19c-062e27a8f9f2

	id := "83e64bbe-01c1-437d-b67f-6184419d1db2"
	url := "http://agent-api.ai.telrobot.top/users/" + id

	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjUxOWVmMmI0NTIwZDJhMzU3ZDNmZTdkZjE1Y2RmOGQ3NDJmOGNjNGViODVlNWI1ZWRmYjcyYjYyODg2NDU2NmYwOWFiZjM2OGRkNjViYjYwIn0.eyJhdWQiOiIxIiwianRpIjoiNTE5ZWYyYjQ1MjBkMmEzNTdkM2ZlN2RmMTVjZGY4ZDc0MmY4Y2M0ZWI4NWU1YjVlZGZiNzJiNjI4ODY0NTY2ZjA5YWJmMzY4ZGQ2NWJiNjAiLCJpYXQiOjE2NTM3MDE2ODQsIm5iZiI6MTY1MzcwMTY4NCwiZXhwIjoxNjg1MjM3Njg0LCJzdWIiOiJiNzM0NTYyMy0xMGRjLTRjNjktOGFhNy1jODdmMThmODljMzciLCJzY29wZXMiOltdfQ.cLqXS2KEOEvfdYOGFMeCqtphHb_JLubwjI1oKYxFgS0bIeMRolTSSWi1BE0HWO5DRmGAxrDOBS4iGqWtsb1FdUhbQnLYIe1aa9jOwGHm6kR_GJHHF3ETai8UGWqaTjnuiC74IsexyIlgr3Qj7w4kDuJuO60SnYcJTrxsIRRh736P71tux61PZrokM9UjAFqrD9Pn7kIHcsCuMGdExQm8MmFkDGC9VIMgfWrdtEAyZd-eEhTpqLlyA-ov_LCQeFiZHhfm18zWH_KFmvYj1ftqv_o8mQ5GcF8Q_43umiPZXH70rEaQutp79KNmt0g08_-Y3TwG2PY57P9WB0_boeRAeycoCWwUk_F0CmulM0eyBs-OIyI6HIxXQOQ2O53OYrPglGhoBklVa6Y7_ie3n94xu_I2x1QiDk6Uf_0NqH2V8CIUu_i1v-e-vQmYWAL_K9TGDOPBiq_J2soyuA7qHfd1kswF2O2nYQdJ6TlS2Fe-WYtN1PiJ3M9UPgKGm8JHekib-XONmhBCl8cdEhV0SXNPFhXx49CgD65CT8tVDB5VRYPWqPLztoRdkZ5LmrTv2XbpNSMBwT4CwXhaeLMPKxtl2HHTHmvPYMCZGaib3TAPdkXkpHKLAMNb91XMu9u8lz1Fq5PRlx5GssCdU_GpOCgww7gyWatFgvzArvaJf_dg60k"

	body := map[string]string{
	}
	body["name"] = "test444"
	body["email"] = "123333333@163.com"
	body["phone"] = "18613875370"
	body["password"] = "123456"
	body["password_confirmation"] = "123456"


	bodybyte, err := json.Marshal(body)
	if err != nil {
		t.Error(err.Error())
	}

	data, err := Put(url, bodybyte,"ContentType: application/json", token)
	if err != nil {
		t.Error(err.Error())
	}

	resp := models.UpdateUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(resp)
}

func TestDeleteById(t *testing.T) {
	//aed884ce-1aee-42cc-b19c-062e27a8f9f2

	id := "83e64bbe-01c1-437d-b67f-6184419d1db2"
	url := "http://agent-api.ai.telrobot.top/users/" + id

	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjUxOWVmMmI0NTIwZDJhMzU3ZDNmZTdkZjE1Y2RmOGQ3NDJmOGNjNGViODVlNWI1ZWRmYjcyYjYyODg2NDU2NmYwOWFiZjM2OGRkNjViYjYwIn0.eyJhdWQiOiIxIiwianRpIjoiNTE5ZWYyYjQ1MjBkMmEzNTdkM2ZlN2RmMTVjZGY4ZDc0MmY4Y2M0ZWI4NWU1YjVlZGZiNzJiNjI4ODY0NTY2ZjA5YWJmMzY4ZGQ2NWJiNjAiLCJpYXQiOjE2NTM3MDE2ODQsIm5iZiI6MTY1MzcwMTY4NCwiZXhwIjoxNjg1MjM3Njg0LCJzdWIiOiJiNzM0NTYyMy0xMGRjLTRjNjktOGFhNy1jODdmMThmODljMzciLCJzY29wZXMiOltdfQ.cLqXS2KEOEvfdYOGFMeCqtphHb_JLubwjI1oKYxFgS0bIeMRolTSSWi1BE0HWO5DRmGAxrDOBS4iGqWtsb1FdUhbQnLYIe1aa9jOwGHm6kR_GJHHF3ETai8UGWqaTjnuiC74IsexyIlgr3Qj7w4kDuJuO60SnYcJTrxsIRRh736P71tux61PZrokM9UjAFqrD9Pn7kIHcsCuMGdExQm8MmFkDGC9VIMgfWrdtEAyZd-eEhTpqLlyA-ov_LCQeFiZHhfm18zWH_KFmvYj1ftqv_o8mQ5GcF8Q_43umiPZXH70rEaQutp79KNmt0g08_-Y3TwG2PY57P9WB0_boeRAeycoCWwUk_F0CmulM0eyBs-OIyI6HIxXQOQ2O53OYrPglGhoBklVa6Y7_ie3n94xu_I2x1QiDk6Uf_0NqH2V8CIUu_i1v-e-vQmYWAL_K9TGDOPBiq_J2soyuA7qHfd1kswF2O2nYQdJ6TlS2Fe-WYtN1PiJ3M9UPgKGm8JHekib-XONmhBCl8cdEhV0SXNPFhXx49CgD65CT8tVDB5VRYPWqPLztoRdkZ5LmrTv2XbpNSMBwT4CwXhaeLMPKxtl2HHTHmvPYMCZGaib3TAPdkXkpHKLAMNb91XMu9u8lz1Fq5PRlx5GssCdU_GpOCgww7gyWatFgvzArvaJf_dg60k"
	data, err := Delete(url, nil,"ContentType: application/json", token)
	if err != nil {
		t.Error(err.Error())
	}

	resp := models.DeleteUserResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(resp)
}



func TestPostWithJson(t *testing.T) {
	url := "http://agent-api.ai.telrobot.top/users"

	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjUxOWVmMmI0NTIwZDJhMzU3ZDNmZTdkZjE1Y2RmOGQ3NDJmOGNjNGViODVlNWI1ZWRmYjcyYjYyODg2NDU2NmYwOWFiZjM2OGRkNjViYjYwIn0.eyJhdWQiOiIxIiwianRpIjoiNTE5ZWYyYjQ1MjBkMmEzNTdkM2ZlN2RmMTVjZGY4ZDc0MmY4Y2M0ZWI4NWU1YjVlZGZiNzJiNjI4ODY0NTY2ZjA5YWJmMzY4ZGQ2NWJiNjAiLCJpYXQiOjE2NTM3MDE2ODQsIm5iZiI6MTY1MzcwMTY4NCwiZXhwIjoxNjg1MjM3Njg0LCJzdWIiOiJiNzM0NTYyMy0xMGRjLTRjNjktOGFhNy1jODdmMThmODljMzciLCJzY29wZXMiOltdfQ.cLqXS2KEOEvfdYOGFMeCqtphHb_JLubwjI1oKYxFgS0bIeMRolTSSWi1BE0HWO5DRmGAxrDOBS4iGqWtsb1FdUhbQnLYIe1aa9jOwGHm6kR_GJHHF3ETai8UGWqaTjnuiC74IsexyIlgr3Qj7w4kDuJuO60SnYcJTrxsIRRh736P71tux61PZrokM9UjAFqrD9Pn7kIHcsCuMGdExQm8MmFkDGC9VIMgfWrdtEAyZd-eEhTpqLlyA-ov_LCQeFiZHhfm18zWH_KFmvYj1ftqv_o8mQ5GcF8Q_43umiPZXH70rEaQutp79KNmt0g08_-Y3TwG2PY57P9WB0_boeRAeycoCWwUk_F0CmulM0eyBs-OIyI6HIxXQOQ2O53OYrPglGhoBklVa6Y7_ie3n94xu_I2x1QiDk6Uf_0NqH2V8CIUu_i1v-e-vQmYWAL_K9TGDOPBiq_J2soyuA7qHfd1kswF2O2nYQdJ6TlS2Fe-WYtN1PiJ3M9UPgKGm8JHekib-XONmhBCl8cdEhV0SXNPFhXx49CgD65CT8tVDB5VRYPWqPLztoRdkZ5LmrTv2XbpNSMBwT4CwXhaeLMPKxtl2HHTHmvPYMCZGaib3TAPdkXkpHKLAMNb91XMu9u8lz1Fq5PRlx5GssCdU_GpOCgww7gyWatFgvzArvaJf_dg60k"

	body := map[string]string{
	}
	body["name"] = "test3"
	body["email"] = "123333333@163.com"
	body["phone"] = "18613875370"
	body["password"] = "123456"
	body["password_confirmation"] = "123456"


	bodybyte, err := json.Marshal(body)
	if err != nil {
		t.Error(err.Error())
	}

	res, err := PostWithJson(url, bodybyte, token)
	if err != nil {
		t.Error(err.Error())
	}

	resp := models.CreateUserResp{}

	err = json.Unmarshal(res, &resp)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(string(res))
	fmt.Println(resp)
}
