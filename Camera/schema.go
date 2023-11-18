package Camera

type CameraSchema struct {
	Camera     string `json:"camera" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Active     bool   `json:"active" validate:"required"`
	SoftDelete bool   `json:"soft_delete"`
	UserId     int32  `json:"user_id" validate:"gte=0"`
	PackageId  int32  `json:"package_id" validate:"gte=0"`
}
type PatchCameraData struct {
	Name       *string `json:"name"`
	Active     *bool   `json:"active"`
	SoftDelete *bool   `json:"soft_delete"`
	UserId     *int32  `json:"user_id" validate:"omitempty,gte=0"`
	PackageId  *int32  `json:"package_id" validate:"omitempty,gte=0"`
}
