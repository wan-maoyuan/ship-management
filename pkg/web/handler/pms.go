package handler

import (
	"net/http"
	"ship-management/pkg/database"
	"ship-management/pkg/web/proto"

	"github.com/gin-gonic/gin"
)

func QueryPMSFirstCategory(c *gin.Context) {
	categorys, err := database.QueryAllPMSEquipmentFirstCategory()
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(categorys),
	)
}

func QueryPMSSecondCategory(c *gin.Context) {
	var err error

	var req = proto.QueryPMSSecondCategoryReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	categorys, err := database.QueryPMSEquipmentSecondCategoryByCode(req.FirstCode)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(categorys),
	)
}

func QueryPMSThirdCategory(c *gin.Context) {
	var err error

	var req = proto.QueryPMSThirdCategoryReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	categorys, err := database.QueryPMSEquipmentThirdCategoryByCode(req.FirstCode, req.SecondCode)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(categorys),
	)
}

func QueryPMSFourthCategory(c *gin.Context) {
	var err error

	var req = proto.QueryPMSFourthCategoryReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	categorys, err := database.QueryPMSEquipmentFourthCategoryByCode(req.FirstCode, req.SecondCode, req.ThirdCode)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(categorys),
	)
}

func QueryPMSJobOrder(c *gin.Context) {
	var err error

	var req = proto.QueryPMSJobOrderReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	orderList, err := database.QueryPMSJobOrder(req.Page, req.Size)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(orderList),
	)
}

func QueryPMSWorkDone(c *gin.Context) {
	var err error

	var req = proto.QueryPMSWorkDoneReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	workList, err := database.QueryPMSWorkDone(req.Page, req.Size)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(workList),
	)
}

func QueryPMSOverDueOrder(c *gin.Context) {
	var err error

	var req = proto.QueryPMSOverDueOrderReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	orderList, err := database.QueryPMSOverDueOrder(req.Page, req.Size)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(orderList),
	)
}
