package uploadphotos

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/daniel-dsouza/test/app/common"
)

// UploadPhotoController uploads struct
type UploadPhotoController struct {
	common.Controller
}

// Upload uploads a photo
func (c *UploadPhotoController) Upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("profile")
	filename := header.Filename
	log.Println(header.Filename)
	out, err := os.Create("./tmp/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	common.UploadToS3(file, "/profile-photos/53")
	c.SendJSON(w, r, filename, http.StatusOK)
}

//Download a photo
func (c *UploadPhotoController) Download(w http.ResponseWriter, r *http.Request) {
	file := common.DownloadFromS3("/profile-photos/53")
	c.SendJPEG(w, r, file, http.StatusOK)
}
