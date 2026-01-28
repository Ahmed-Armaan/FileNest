package storage

//
//import (
//	"fmt"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/Ahmed-Armaan/FileNest/storage/helper"
//	"github.com/gin-gonic/gin"
//)
//
//func TestFileUpload(t *testing.T) {
//	// create new upload
//	uploadId, objectKey, err := helper.CreateNewUpload(ctx, s3Client, bucketName, "__test__/")
//	if err != nil {
//		t.Fatalf("New upload init failed %s\n", err)
//	}
//
//	if len(uploadId) == 0 {
//		t.Fatal("No uploadId generated")
//	}
//	if len(objectKey) == 0 {
//		t.Fatal("No objectKey generated")
//	}
//
//	// get upload url
//	r := gin.New()
//	r.GET("/get_url", GetUploadUrl)
//
//	reqParams := fmt.Sprintf("/get_url?uploadId=%s&objectKey=%s&partNumber=%d", uploadId, objectKey, 1)
//	req := httptest.NewRequest("GET", reqParams, nil)
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//}
