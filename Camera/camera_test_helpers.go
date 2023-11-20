package Camera

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"startProject/utils"
	"testing"
)

type IDS struct {
	ID string `json:"id"`
}

func SetUpCamera(datas []CameraSchema) ([]IDS, error) {
	errorConnection := utils.Connect("TEST")
	idObj := make([]IDS, 0, len(datas))
	if errorConnection != nil {
		return nil, errorConnection
	}

	utils.CreateCollection("camera")
	camera_collection := utils.GetCollection("camera")
	for index := range datas {
		data := datas[index]
		resultInsert, errorInsert := camera_collection.InsertOne(context.Background(), bson.M{
			"camera":      data.Camera,
			"name":        data.Name,
			"active":      data.Active,
			"soft_delete": data.SoftDelete,
			"user_id":     data.UserId,
			"package_id":  data.PackageId,
		})
		if errorInsert != nil {
			return nil, errorInsert
		}
		insertedID, ok := resultInsert.InsertedID.(primitive.ObjectID)
		if !ok {
			return nil, fmt.Errorf("InsertedID is not an ObjectID")
		}
		idAsString := insertedID.Hex()

		idObj = append(idObj, IDS{ID: idAsString})
	}

	return idObj, nil
}
func TearDown() {
	utils.EmptyTestDB()
}

func TestRetrieveCamera(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RetrieveCamera(w, r)
	}))
	defer TearDown()
	defer server.Close()
	cameraData := []CameraSchema{
		{
			Camera:     "test_camera1",
			Name:       "test_name1",
			Active:     true,
			SoftDelete: false,
			UserId:     12,
			PackageId:  12,
		},
		{
			Camera:     "test_camera2",
			Name:       "test_name2",
			Active:     true,
			SoftDelete: false,
			UserId:     13,
			PackageId:  14,
		},
	}

	ids, errorSetUp := SetUpCamera(cameraData)
	if errorSetUp != nil {
		fmt.Println(errorSetUp.Error())
		//return
	}

	for index := range ids {

		request, _ := http.NewRequest("GET", "/camera/id/?id="+ids[index].ID, nil)
		response := httptest.NewRecorder()
		server.Config.Handler.ServeHTTP(response, request)
		if http.StatusOK != response.Code {
			t.Errorf("Failed need 200 get %d", response.Code)
		}
		//assert.Equal(t, http.StatusOK, response.Code)
	}

}

func TestListCamera(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ListCamera(w, r)
	}))
	defer TearDown()
	defer server.Close()
	cameraData := []CameraSchema{
		{
			Camera:     "test_camera1",
			Name:       "test_name1",
			Active:     true,
			SoftDelete: false,
			UserId:     12,
			PackageId:  12,
		},
		{
			Camera:     "test_camera2",
			Name:       "test_name2",
			Active:     true,
			SoftDelete: false,
			UserId:     13,
			PackageId:  14,
		},
	}
	_, errorSetUp := SetUpCamera(cameraData)
	if errorSetUp != nil {
		fmt.Println(errorSetUp.Error())
		return
	}

	request, _ := http.NewRequest("GET", "/camera/", nil)
	response := httptest.NewRecorder()
	server.Config.Handler.ServeHTTP(response, request)
	if http.StatusOK != response.Code {
		t.Errorf("Failed need 200 get %d", response.Code)
	}
	//assert.Equal(t, http.StatusOK, response.Code)

}

func TestInsertCamera(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		InsertCamera(w, r)
	}))
	defer TearDown()
	defer server.Close()

	testbody := `{"camera":"test1","name":"test1","active":true,"soft_delete":true,"user_id":1000, "package_id":1000}`
	request, _ := http.NewRequest("POST", "/camera/insert/", bytes.NewBufferString(testbody))
	response := httptest.NewRecorder()
	server.Config.Handler.ServeHTTP(response, request)
	//assert.Equal(t, http.StatusCreated, response.Code)
	if http.StatusCreated != response.Code {
		t.Errorf("Failed need 201 get %d", response.Code)
	}
}

func TestUpdateCamera(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateCamera(w, r)
	}))
	defer TearDown()
	defer server.Close()
	cameraData := []CameraSchema{
		{
			Camera:     "test_camera1",
			Name:       "test_name1",
			Active:     true,
			SoftDelete: false,
			UserId:     12,
			PackageId:  12,
		},
		{
			Camera:     "test_camera2",
			Name:       "test_name2",
			Active:     true,
			SoftDelete: false,
			UserId:     13,
			PackageId:  14,
		},
	}
	resultSetUp, errorSetUp := SetUpCamera(cameraData)
	if errorSetUp != nil {
		fmt.Println(errorSetUp.Error())
		return
	}
	for index := range resultSetUp {
		testbody := `{"camera":"test1","name":"test1","active":true,"soft_delete":true,"user_id":1000, "package_id":1000}`
		request, _ := http.NewRequest("PUT", "/camera/update/?id="+resultSetUp[index].ID, bytes.NewBufferString(testbody))
		response := httptest.NewRecorder()
		server.Config.Handler.ServeHTTP(response, request)
		if http.StatusOK != response.Code {
			t.Errorf("Failed need 200 get %d", response.Code)
		}

	}
}

func TestDeleteCamera(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteCamera(w, r)
	}))
	defer TearDown()
	defer server.Close()
	cameraData := []CameraSchema{
		{
			Camera:     "test_camera1",
			Name:       "test_name1",
			Active:     true,
			SoftDelete: false,
			UserId:     12,
			PackageId:  12,
		},
		{
			Camera:     "test_camera2",
			Name:       "test_name2",
			Active:     true,
			SoftDelete: false,
			UserId:     13,
			PackageId:  14,
		},
	}
	resultSetUp, errorSetUp := SetUpCamera(cameraData)
	if errorSetUp != nil {
		fmt.Println(errorSetUp.Error())
		return
	}
	for index := range resultSetUp {
		request, _ := http.NewRequest("DELETE", "/camera/delete/?id="+resultSetUp[index].ID, nil)
		response := httptest.NewRecorder()
		server.Config.Handler.ServeHTTP(response, request)
		if http.StatusOK != response.Code {
			t.Errorf("Failed need 200 get %d", response.Code)
		}
	}

}

func RunAllTests(t *testing.T) {
	t.Run("RetrieveCamera", TestRetrieveCamera)
	t.Run("ListCamera", TestListCamera)
	t.Run("InsertCamera", TestInsertCamera)
	t.Run("UpdateCamera", TestUpdateCamera)
	t.Run("DeleteCamera", TestDeleteCamera)
	return

}
