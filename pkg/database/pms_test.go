package database

import (
	"ship-management/pkg/model"
	"testing"
	"time"
)

func TestInsertPMSJobOrder(t *testing.T) {
	dsn := "sqlserver://sa:smart_ship123..@172.31.126.33:1433?database=ISM_Vsl_9977551&encrypt=disable"
	if err := InitDatabase(dsn); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	for i := 0; i < 100; i++ {
		if err := InsertPMSJobOrder(&model.PMSJobOrder{
			VesselName:         "test",
			VesselIMO:          "123124234",
			OrderID:            "sadjhasji",
			Task:               "211",
			FirstCode:          "12",
			SecondCode:         "123",
			ThirdCode:          "sad",
			FourthCode:         "fgdsg",
			Equipment:          "sadsa",
			EquipmentComponent: "das",
			Remarks:            "sadsa",
			PlannedStartTime:   time.Now().UTC(),
			CreateTime:         time.Now().UTC(),
			Completed:          false,
			Deleted:            false,
		}); err != nil {
			t.Errorf("InsertPMSJobOrder: %v", err)
			continue
		}
	}
}

func TestQueryPMSJobOrder(t *testing.T) {
	dsn := "sqlserver://sa:smart_ship123..@172.31.126.33:1433?database=ISM_Vsl_9977551&encrypt=disable"
	if err := InitDatabase(dsn); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	for page := 1; page < 13; page++ {
		orders, err := QueryPMSJobOrder(uint(page), 10)
		if err != nil {
			t.Errorf("QueryPMSJobOrder: %v", err)
			continue
		}

		for _, order := range orders {
			t.Logf("order id: %d", order.ID)
		}
	}
}

func TestInsertPMSWorkDone(t *testing.T) {
	dsn := "sqlserver://sa:smart_ship123..@172.31.126.33:1433?database=ISM_Vsl_9977551&encrypt=disable"
	if err := InitDatabase(dsn); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	for i := 0; i < 100; i++ {
		if err := InsertPMSWorkDone(&model.PMSWorkDone{
			VesselName:         "test",
			VesselIMO:          "123124234",
			OrderID:            "sadjhasji",
			Task:               "211",
			FirstCode:          "12",
			SecondCode:         "123",
			ThirdCode:          "sad",
			FourthCode:         "fgdsg",
			Equipment:          "sadsa",
			EquipmentComponent: "das",
			Result:             "success",
			Remarks:            "nothing",
			SpeedHours:         100,
			PlannedStartTime:   time.Now().UTC(),
			EndTime:            time.Now().UTC(),
		}); err != nil {
			t.Errorf("InsertPMSWorkDone: %v", err)
			continue
		}
	}
}

func TestQueryPMSWorkDone(t *testing.T) {
	dsn := "sqlserver://sa:smart_ship123..@172.31.126.33:1433?database=ISM_Vsl_9977551&encrypt=disable"
	if err := InitDatabase(dsn); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	for page := 1; page < 13; page++ {
		orders, err := QueryPMSWorkDone(uint(page), 10)
		if err != nil {
			t.Errorf("QueryPMSWorkDone: %v", err)
			continue
		}

		for _, order := range orders {
			t.Logf("order id: %d", order.ID)
		}
	}
}

func TestInsertPMSOverDueOrder(t *testing.T) {
	dsn := "sqlserver://sa:smart_ship123..@172.31.126.33:1433?database=ISM_Vsl_9977551&encrypt=disable"
	if err := InitDatabase(dsn); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	for i := 0; i < 100; i++ {
		if err := InsertPMSOverDueOrder(&model.PMSOverDueOrder{
			VesselName:         "test",
			VesselIMO:          "123124234",
			OrderID:            "sadjhasji",
			Task:               "211",
			FirstCode:          "12",
			SecondCode:         "123",
			ThirdCode:          "sad",
			FourthCode:         "fgdsg",
			Equipment:          "sadsa",
			EquipmentComponent: "das",
			Reason:             "no replace",
			Remarks:            "nothing",
			SpeedHours:         100,
			PlannedStartTime:   time.Now().UTC(),
		}); err != nil {
			t.Errorf("InsertPMSOverDueOrder: %v", err)
			continue
		}
	}
}

func TestQueryPMSOverDueOrder(t *testing.T) {
	dsn := "sqlserver://sa:smart_ship123..@172.31.126.33:1433?database=ISM_Vsl_9977551&encrypt=disable"
	if err := InitDatabase(dsn); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	for page := 1; page < 13; page++ {
		orders, err := QueryPMSOverDueOrder(uint(page), 10)
		if err != nil {
			t.Errorf("QueryPMSOverDueOrder: %v", err)
			continue
		}

		for _, order := range orders {
			t.Logf("order id: %d", order.ID)
		}
	}
}
