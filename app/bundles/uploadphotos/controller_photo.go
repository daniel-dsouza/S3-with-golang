package uploadphotos

import (
	"log"
	"net/http"

	"fmt"

	"github.com/daniel-dsouza/test/app/common"
)

// UploadPhotoController uploads struct
type UploadPhotoController struct {
	common.Controller
}

// Upload uploads a photo
func (c *UploadPhotoController) Upload(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	file, header, err := r.FormFile("profile")
	if err != nil {
		log.Fatal(err)
	}

	filename := header.Filename
	log.Println(header.Filename)
	//  log.Println(fmt.Sprintf("/profile-photos/%s", user))
	//	out, err := os.Create("./tmp/" + filename)
	//	if err != nil {
	//		log.Fatal(err)
	//
	//
	//	defer out.Close()
	//	_, err = io.Copy(out, file)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	common.UploadToS3(file, fmt.Sprintf("/profile-photos/%s", user))
	c.SendJSON(w, r, filename, http.StatusOK)
}

//Download a photo
func (c *UploadPhotoController) Download(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	log.Println(user)
	file := common.DownloadFromS3(fmt.Sprintf("/profile-photos/%s", user))
	c.SendJPEG(w, r, file, http.StatusOK)
}
