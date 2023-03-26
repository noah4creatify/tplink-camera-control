package pkg

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"
)

var (
	ErrReadConfigFile = errors.New("read config file error")
	ErrAddressNotSet  = errors.New("address not set")
	ErrPassWordNotSet = errors.New("password not set")
	ErrUnAuthorized   = errors.New("unauthorized")
	ErrTurnCameraOff  = errors.New("turn camera off error")
	ErrTurnCameraOn   = errors.New("turn camera on error")
)

type Router struct {
	config *ConfigOptions
	pubKey *rsa.PublicKey
}

func (r *Router) PostData(data []byte) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return doPostRequest(ctx, fmt.Sprintf(PostDsUrl, r.config.Address, r.config.Stok), data)
}

func (r *Router) RenewPostData(resData, data []byte) ([]byte, error) {
	if err := r.RenewStok(resData); err != nil {
		return nil, err
	}
	return r.PostData(data)
}

func (r *Router) RetryPostDataWhenNotAuth(data []byte) ([]byte, error) {
	resData, err := r.PostData(data)
	if err != nil {
		switch err {
		case ErrUnAuthorized:
			return r.RenewPostData(resData, data)
		}
		return nil, err
	}
	return resData, nil
}

func (r *Router) RenewStok(data []byte) (err error) {
	unAuthRes := struct {
		Data struct {
			Code        int      `json:"code"`
			EncryptType []string `json:"encrypt_type"`
			Key         string   `json:"key"`
			Nonce       string   `json:"nonce"`
		} `json:"data"`
		ErrorCode int `json:"error_code"`
	}{}
	if err = json.Unmarshal(data, &unAuthRes); err != nil {
		return
	}
	if unAuthRes.ErrorCode != ErrUnAuthorizedCode {
		return fmt.Errorf("unexpected error code: %d", unAuthRes.ErrorCode)
	}
	tpEncryptPasswd := securityEncode(r.config.PassWord)
	postPasswd, err := encrypt(r.pubKey, tpEncryptPasswd+":"+unAuthRes.Data.Nonce)
	if err != nil {
		return err
	}

	payload := fmt.Sprintf(PayloadLogin, r.config.UserName, postPasswd)
	resData, err := r.PostData([]byte(payload))
	if err != nil {
		return err
	}

	stokRes := struct {
		Stok      string `json:"stok"`
		ErrorCode int    `json:"error_code"`
	}{}

	if err = json.Unmarshal(resData, &stokRes); err != nil {
		return err
	}

	return r.config.UpdateSaveStok(stokRes.Stok)
}

func (r *Router) GetBaseInfo() {
	resData, err := r.RetryPostDataWhenNotAuth([]byte(PayloadGetBasicInfo))
	if err != nil {
		log.Fatalf("post base info error: %v", err)
	}

	var res struct {
		DeviceInfo struct {
			BasicInfo map[string]any `json:"basic_info"`
		} `json:"device_info"`
		ErrorCode int64 `json:"error_code"`
	}
	if err = json.Unmarshal(resData, &res); err != nil {
		log.Fatalf("unmarshal base info error: %v", err)
	}

	for k, v := range res.DeviceInfo.BasicInfo {
		if n, ok := v.(string); ok {
			v, _ = url.QueryUnescape(n)
		}
		fmt.Printf("%v:%v\n", k, v)
	}
}

func (r *Router) IsLenMaskOn() bool {
	resData, err := r.RetryPostDataWhenNotAuth([]byte(PayloadGetLensmask))
	if err != nil {
		log.Fatalf("post len_mask info error: %v", err)
	}

	var res struct {
		LensMask struct {
			LensMaskInfo struct {
				Enabled string `json:"enabled"`
			} `json:"lens_mask_info"`
		} `json:"lens_mask"`
		ErrorCode int64 `json:"error_code"`
	}

	if err = json.Unmarshal(resData, &res); err != nil {
		log.Fatalf("unmarshal len_mask info error: %v", err)
	}
	return res.LensMask.LensMaskInfo.Enabled == "on"
}

func (r *Router) SetLenMaskOff() error {
	_, err := r.RetryPostDataWhenNotAuth([]byte(PayloadSetLensmaskOff))
	return err
}

func (r *Router) SetLenMaskOn() error {
	_, err := r.RetryPostDataWhenNotAuth([]byte(PayloadSetLensmaskOn))
	return err
}

func (r *Router) GotoPreset(id string) error {
	_, err := r.RetryPostDataWhenNotAuth([]byte(fmt.Sprintf(PayloadGotoPreset, id)))
	return err
}

func (r *Router) TurnOnCamera() error {
	if err := r.SetLenMaskOff(); err != nil {
		return err
	}
	if err := r.GotoPreset("1"); err != nil {
		return err
	}
	time.Sleep(time.Second)
	if r.IsLenMaskOn() {
		return ErrTurnCameraOn
	}
	return nil
}

func (r *Router) TurnOffCamera() error {
	if err := r.SetLenMaskOn(); err != nil {
		return err
	}
	time.Sleep(time.Second)
	if !r.IsLenMaskOn() {
		return ErrTurnCameraOff
	}
	return nil
}

func NewRouter(options *ConfigOptions) (*Router, error) {
	pubKey, err := newPubKey(ipc44awN, ipc44awE)
	if err != nil {
		return nil, err
	}

	return &Router{
		config: options,
		pubKey: pubKey,
	}, nil
}
