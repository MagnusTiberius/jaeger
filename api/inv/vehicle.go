package inv

type VehicleEntity struct {
	ManufacturerCode  			string
	ModelCode 					string
	TrimCode 					string
}


func NewVehicleEntity() *VehicleEntity {
	v := new(VehicleEntity)
	return v
}