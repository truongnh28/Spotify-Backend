package dto

import "mime/multipart"

type UploadIn struct {
	File multipart.File
}
