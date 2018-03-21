package model

import (
	"fmt"
	"superBrain/types"
	"github.com/golang/glog"
)

func GetTestData()([]types.Test, error){
	var testData []types.Test
	hrows, herr := DB.Query(fmt.Sprintf(`SELECT ID, NAME FROM SYS_ROLE`))
	if herr != nil {
		glog.Errorf("GetTestData: sql return err %v\n", herr)
		return nil, herr
	}
	for hrows.Next() {
		var hrow = new(types.Test)
		if herr = hrows.Scan(&hrow.Id, &hrow.Name); herr != nil {
			return nil, herr
		}
		testData = append(testData, *hrow)
	}
	if hrows.Err() != nil {
		glog.Errorf("GetTestData: sql rows.Err() %v\n", herr)
		return nil, hrows.Err()
	}
	return  testData, nil
}
