package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jptosso/coraza-center/utils"
	"github.com/mholt/archiver/v4"
)

type Remote struct {
	Server   string
	Username string
	Password string
}

func (r *Remote) Download(tag string, dst string) error {
	url := fmt.Sprintf("%s/v1/waf/%s", r.Server, tag)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", r.authString())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bts := bytes.NewBuffer(nil)
	if _, err := bts.ReadFrom(resp.Body); err != nil {
		return err
	}
	if err := utils.UntarTo(bts.Bytes(), dst); err != nil {
		return fmt.Errorf("Failed to untar downloaded file: %s", err.Error())
	}
	return nil
}

func (r *Remote) Upload(dir string, tag string) error {
	// upload tar as multipart
	var err error
	tar, err := dirToTar(dir)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	var fw io.Writer
	if x, ok := tar.(io.Closer); ok {
		defer x.Close()
	}
	// Add an image file
	if fw, err = w.CreateFormFile("file", "upload.tar.gz"); err != nil {
		return err
	}
	if _, err = io.Copy(fw, tar); err != nil {
		return err
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()
	transport := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{Transport: transport}
	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", r.Server+"/v1/waf/"+tag, &b)
	if err != nil {
		return err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", r.authString())
	req.Header.Set("If-Modified-Since", time.Now().Format(time.ANSIC))

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		bts, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("bad status: %s\nError: %s", res.Status, string(bts))
	}
	return nil
}

func (r *Remote) authString() string {
	// now we base64 encode data
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", r.Username, r.Password)))
}

func dirToTar(dir string) (io.Reader, error) {
	// we list all files in the directory
	files := map[string]string{}
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && !strings.Contains(path, "/.coraza") {
				files[path] = path[len(dir)+1:]
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	format := archiver.CompressedArchive{
		Compression: archiver.Gz{},
		Archival:    archiver.Tar{},
	}
	flist, err := archiver.FilesFromDisk(nil, files)
	if err != nil {
		return nil, err
	}
	out := &bytes.Buffer{}
	// create the archive
	err = format.Archive(context.Background(), out, flist)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewRemote(server string, username string, password string) *Remote {
	return &Remote{
		Server:   server,
		Username: username,
		Password: password,
	}
}
