package proto

type QueryPMSSecondCategoryReq struct {
	FirstCode string `form:"first_code"`
}

type QueryPMSThirdCategoryReq struct {
	FirstCode  string `form:"first_code"`
	SecondCode string `form:"second_code"`
}

type QueryPMSFourthCategoryReq struct {
	FirstCode  string `form:"first_code"`
	SecondCode string `form:"second_code"`
	ThirdCode  string `form:"third_code"`
}

type QueryPMSJobOrderReq struct {
	Page uint `form:"page"`
	Size uint `form:"size"`
}

type QueryPMSWorkDoneReq struct {
	Page uint `form:"page"`
	Size uint `form:"size"`
}

type QueryPMSOverDueOrderReq struct {
	Page uint `form:"page"`
	Size uint `form:"size"`
}
