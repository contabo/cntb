package util

import (
	"encoding/json"
	"time"

	"github.com/cheggaaa/pb"
)

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}

func ShowUploadDetails(fileSize int64) *pb.ProgressBar{
	progress := pb.New64(fileSize)
	progress.Start()
	progress.SetRefreshRate(time.Second)
	progress.SetUnits(pb.U_BYTES)
	progress.Prefix("Uploading")

	return progress
}