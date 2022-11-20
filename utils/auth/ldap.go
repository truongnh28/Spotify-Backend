package auth

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"strings"
	"time"

	"spotify/config"

	"github.com/go-ldap/ldap"
)

type ldapClient struct {
	cfg *config.LDAP
}

type BasicAuth struct {
	Username string
	Password string
}

// InitLDAP should call within init func of main
func InitLDAP(cfg interface{}) IAuth {
	cfgLDAP, ok := cfg.(*config.LDAP)
	if !ok {
		panic("can not parse config ldap")
	}
	return &ldapClient{
		cfg: cfgLDAP,
	}
}

func (c *ldapClient) initConn() (*ldap.Conn, error) {
	var (
		conn *ldap.Conn
		err  error
	)
	if c.cfg.UseTls {
		tlsCfg := &tls.Config{
			InsecureSkipVerify: false,
			ServerName:         strings.Split(c.cfg.Addr, ":")[0],
			MinVersion:         tls.VersionTLS12,
		}
		conn, err = ldap.DialTLS("tcp", c.cfg.Addr, tlsCfg)
	} else {
		conn, err = ldap.Dial("tcp", c.cfg.Addr)
	}
	return conn, err
}

func (c *ldapClient) Login(ctx context.Context, cert interface{}) (*UserInfo, error) {
	var (
		user *UserInfo
		cfg  = c.cfg
		err  error
	)

	basicAuth, ok := cert.(*BasicAuth)
	if !ok {
		return nil, errors.New("can not convert into basic auth")
	}

	conn, err := c.initConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	conn.SetTimeout(time.Duration(c.cfg.Timeout) * time.Second)
	err = conn.Bind(cfg.Username, cfg.Password)
	if err != nil {
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		cfg.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(&(objectClass=%s)(uid=%s))", cfg.ObjectClass, basicAuth.Username),
		[]string{"mail", "uid"},
		nil,
	)

	searchResponse, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if len(searchResponse.Entries) != 1 {
		return nil, nil
	}

	err = conn.Bind(searchResponse.Entries[0].DN, basicAuth.Password)
	if err != nil {
		return nil, nil
	}

	user = &UserInfo{
		Username: basicAuth.Username,
		Email:    searchResponse.Entries[0].GetAttributeValue("mail"),
	}
	return user, nil
}
