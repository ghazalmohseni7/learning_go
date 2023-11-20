package Camera

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"startProject/utils"
	"time"
)

func ListCamera(w http.ResponseWriter, r *http.Request) {

	var cameras []CameraSchema
	errorConnection := utils.Connect("TEST")
	if errorConnection != nil {
		http.Error(w, "can not connect to db\n"+errorConnection.Error(), http.StatusBadRequest)
		return
	}
	camera_collection := utils.GetCollection("camera")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findResult, findError := camera_collection.Find(ctx, bson.D{})
	if findError != nil {
		http.Error(w, "can not find anything\n"+findError.Error(), http.StatusBadRequest)
	}
	for findResult.Next(context.Background()) {
		var camera CameraSchema
		errorJsonConvert := findResult.Decode(&camera)
		if errorJsonConvert != nil {
			http.Error(w, "can not convert the data to json \n"+errorJsonConvert.Error(), http.StatusBadRequest)
			return
		}
		cameras = append(cameras, camera)
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    cameras,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}
func InsertCamera(w http.ResponseWriter, r *http.Request) {

	var postdata CameraSchema
	var validate = validator.New()
	errconnections := utils.Connect("TEST")
	if errconnections != nil {
		http.Error(w, "can not connect to db\n"+errconnections.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)

	errdecoder := decoder.Decode(&postdata)

	if errdecoder != nil {
		http.Error(w, "Error happed during decoding json\n"+errdecoder.Error(), http.StatusBadRequest)
		return
	}

	validateerror := validate.Struct(postdata)
	if validateerror != nil {
		http.Error(w, validateerror.Error(), http.StatusBadRequest)
		return
	}

	camera_collection := utils.GetCollection("camera")
	insertResult, inserterror := camera_collection.InsertOne(context.Background(), postdata)
	if inserterror != nil {
		http.Error(w, "can not add the data\n"+inserterror.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"success":    true,
		"message":    "Data inserted successfully",
		"insertedID": insertResult.InsertedID,
	}
	responseJson, errorResponseJson := json.Marshal(response)
	if errorResponseJson != nil {
		http.Error(w, "Error encoding JSON\n"+errorResponseJson.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(responseJson)

}
func UpdateCamera(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	recordId := r.URL.Query().Get("id")
	if recordId == "" {
		http.Error(w, "please enter the id of record\n", http.StatusBadRequest)
		return
	}
	objectID, err := primitive.ObjectIDFromHex(recordId)
	if err != nil {
		http.Error(w, "invalid ObjectId format\n", http.StatusBadRequest)
		return
	}
	//fmt.Println(recordId)
	var updateDate CameraSchema
	var validate = validator.New()
	errorConnection := utils.Connect("TEST")
	if errorConnection != nil {
		http.Error(w, "can not connect to db\n"+errorConnection.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println("here is the r.body : \n", r.Body)
	decoder := json.NewDecoder(r.Body)
	errorDecoder := decoder.Decode(&updateDate)
	if errorDecoder != nil {
		http.Error(w, "some error in data you send\n"+errorDecoder.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println("here is the data i got from json ")
	//fmt.Println(updateDate)
	errorValidation := validate.Struct(updateDate)
	if errorValidation != nil {
		http.Error(w, "send the correct data\n"+errorValidation.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println("here is the validated data")
	//fmt.Println(updateDate)
	camera_collection := utils.GetCollection("camera")
	filter := bson.D{{"_id", objectID}}
	udpate := bson.D{
		{"$set", bson.D{
			{"camera", updateDate.Name},
			{"active", updateDate.Active},
			{"softdelete", updateDate.SoftDelete},
			{"userid", updateDate.UserId},
			{"packageid", updateDate.PackageId},
		}},
	}
	//fmt.Println(udpate)
	//fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVvvvv")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	updateResult, errorUpdate := camera_collection.UpdateOne(ctx, filter, udpate)
	//fmt.Println(errorUpdate.Error())
	if errorUpdate != nil {
		http.Error(w, "some error during update \n"+errorUpdate.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println("record is updateddddd")
	fmt.Println(updateResult)

	response := map[string]interface{}{
		"success": true,
		"message": "Data updated successfully",
	}

	responseJson, errorResponseJson := json.Marshal(response)
	if errorResponseJson != nil {
		http.Error(w, "Error encoding JSON\n"+errorResponseJson.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
func DeleteCamera(w http.ResponseWriter, r *http.Request) {

	recordId := r.URL.Query().Get("id")
	if recordId == "" {
		http.Error(w, "enter the valid id\n", http.StatusBadRequest)
		return
	}
	errorConnection := utils.Connect("TEST")
	if errorConnection != nil {
		http.Error(w, "can not connect to db\n"+errorConnection.Error(), http.StatusBadRequest)
		return
	}
	camera_collection := utils.GetCollection("camera")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ObjectId, errorObjectId := primitive.ObjectIDFromHex(recordId)
	if errorObjectId != nil {
		http.Error(w, "recordid can not be converted\n"+errorObjectId.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.D{{"_id", ObjectId}}
	_, deleteError := camera_collection.DeleteOne(ctx, filter)
	if deleteError != nil {
		http.Error(w, "can not delete the record\n"+deleteError.Error(), http.StatusBadRequest)
		return
	} else {
		response := map[string]interface{}{
			"message": "Data deleted successfully",
			"success": true,
		}
		w.Header().Set("Content-Type", "application/json")
		responseJson, errorResponseJson := json.Marshal(response)
		if errorResponseJson != nil {
			http.Error(w, "Error encoding JSON\n"+errorResponseJson.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	}

}
func RetrieveCamera(w http.ResponseWriter, r *http.Request) {

	recordId := r.URL.Query().Get("id")
	if recordId == "" {
		http.Error(w, "send valid id \n", http.StatusBadRequest)
		return
	}
	errorConnection := utils.Connect("TEST")
	if errorConnection != nil {
		http.Error(w, "can not connect to db\n"+errorConnection.Error(), http.StatusBadRequest)
		return
	}
	camera_collection := utils.GetCollection("camera")
	ObjectId, errorObjectId := primitive.ObjectIDFromHex(recordId)
	if errorObjectId != nil {
		http.Error(w, "can not conver id\n"+errorObjectId.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.D{{"_id", ObjectId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findResult := camera_collection.FindOne(ctx, filter)
	w.Header().Set("Content-Type", "application/json")
	var camera CameraSchema
	errorJsonConvert := findResult.Decode(&camera)
	if errorJsonConvert != nil {
		http.Error(w, "can not convert the data to json \n"+errorJsonConvert.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"data":    camera,
		"success": true,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}

}

//mordeshore patial update ro bebarn
//func UpdateCamera(w http.ResponseWriter, r *http.Request) {
//	recordId := r.URL.Query().Get("id")
//	if recordId == "" {
//		http.Error(w, "send the id of camera too", http.StatusBadRequest)
//		return
//	}
//	fmt.Println(recordId)
//	var updateData PatchCameraData
//	var validate = validator.New()
//	errconnections := utils.Connect()
//	if errconnections != nil {
//		http.Error(w, "can not connect to db\n"+errconnections.Error(), http.StatusInternalServerError)
//		return
//	}
//	fmt.Println("here is the r.body : ", r.Body)
//	decoder := json.NewDecoder(r.Body)
//	errorderocder := decoder.Decode(&updateData)
//	if errorderocder != nil {
//		http.Error(w, "something happen \n"+errorderocder.Error(), http.StatusBadRequest)
//		return
//	}
//	errorvalidation := validate.Struct(updateData)
//	if errorvalidation != nil {
//		http.Error(w, "enter valid data\n"+errorvalidation.Error(), http.StatusBadRequest)
//	}
//	camera_collection := utils.GetCollection("camera")
//	filter := bson.D{{"_id", recordId}}
//	var update bson.D
//	for i := 0; i < reflect.ValueOf(updateData).NumField(); i++ {
//		field := updateData.Field(i)
//		fmt.Println("NNNNNNNNNNNNNNNNNN")
//		fmt.Println(field)
//		fmt.Println("NNNNNNNNNNNNNNNNNN")
//		update = bson.D{
//			{"$set", bson.D{
//				{field, updateData[field]},
//				// Add other fields to update as needed
//			}},
//		}
//	}
//	//update := bson.D{
//	//	{"$set", bson.D{
//	//		{"fieldToUpdate", updateData.FieldToUpdate},
//	//		// Add other fields to update as needed
//	//	}},
//	//}
//	updateError := camera_collection.FindOneAndUpdate(ctx, filter, update)
//	if updateError != nil {
//		http.Error(w, "can not update the record \n"+updateError.Error(), http.StatusBadRequest)
//		return
//	}
//
//}
