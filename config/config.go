package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type LDAP struct {
	Addr        string
	UseTls      bool
	Username    string
	Password    string
	BaseDN      string
	ObjectClass string
	Timeout     int64
	Enable      bool
}

type Auth struct {
	SecretKey      string `json:"secretKey"`
	ExpiredTime    int64  `json:"expiredTime"`
	RbacModelPath  string `json:"rbac-model-path"`
	RbacPolicyPath string `json:"rbac-policy-path"`
	CookiePath     string `json:"cookie-path"`
	CookieSecure   bool   `json:"cookie-secure"`
}

func AuthConfig() *Auth {
	return &Auth{
		SecretKey:      viper.GetString("app.auth.secretKey"),
		ExpiredTime:    viper.GetInt64("app.auth.expiredTime"),
		RbacModelPath:  viper.GetString("app.auth.rbac-model-path"),
		RbacPolicyPath: viper.GetString("app.auth.rbac-policy-path"),
		CookiePath:     viper.GetString("app.auth.cookie-path"),
		CookieSecure:   viper.GetBool("app.auth.cookie-secure"),
	}
}

type BaseConfig struct {
	BasePath  string
	BasicAuth string
}

type CloudinaryConfig struct {
	Name      string
	APIKey    string
	APISecret string
	Folder    string
}

func EnvCloudName() CloudinaryConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return CloudinaryConfig{
		Name:      os.Getenv("CLOUDINARY_CLOUD_NAME"),
		APIKey:    os.Getenv("CLOUDINARY_API_KEY"),
		APISecret: os.Getenv("CLOUDINARY_API_SECRET"),
		Folder:    os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),
	}
}
