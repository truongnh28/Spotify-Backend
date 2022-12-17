package client

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang/glog"
	"sync"
)

type CloudinaryAPI interface {
	UploadImage(ctx context.Context, imagePath string) (*uploader.UploadResult, error)
	UploadMusic(ctx context.Context, musicPath string) (*uploader.UploadResult, error)
}
type cloudinaryAPI struct {
	client *cloudinary.Cloudinary
}

var cloudinaryInstance *cloudinaryAPI
var cloudinaryMutex sync.Mutex

func GetCloudinaryAPI() CloudinaryAPI {
	if cloudinaryInstance != nil {
		return cloudinaryInstance
	}
	cloudinaryMutex.Lock()
	defer cloudinaryMutex.Unlock()
	const cldUrl = "cloudinary://512616158545567:mClhxuKZ9F-EsP4Kjm_s5qccdvk@dbk0cmzcb"
	var cld, err = cloudinary.NewFromURL(cldUrl)
	if err != nil {
		//log.Fatalf("Failed to intialize Cloudinary, %v", err)
		panic(fmt.Errorf("unable to connect to cloudinary: %v", err.Error()))
	}
	cloudinaryInstance = &cloudinaryAPI{
		client: cld,
	}
	return cloudinaryInstance
}

func (c *cloudinaryAPI) UploadImage(ctx context.Context, imagePath string) (*uploader.UploadResult, error) {
	resp, err := c.client.Upload.Upload(ctx, imagePath, uploader.UploadParams{})
	if err != nil {
		glog.Errorln("Upload cloudinary fail: ", err)
		return nil, err
	}
	return resp, nil
}

func (c *cloudinaryAPI) UploadMusic(ctx context.Context, musicPath string) (*uploader.UploadResult, error) {
	resp, err := c.client.Upload.Upload(ctx, musicPath, uploader.UploadParams{})
	if err != nil {
		glog.Errorln("Upload cloudinary fail: ", err)
		return nil, err
	}
	return resp, nil
}
