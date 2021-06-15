package logic

import (
	"go.mod/dao"
	"go.mod/models"
	"go.mod/tool"
)

// @Title [Authorize needed] AddDetail
// @Description  Add Detail
func AddDetail(detail *models.Detail) {
	dao.DB.Create(detail)
}

// @Title [Authorize needed] UpdateDetail
// @Description  Update Detail
func UpdateDetail(detail *models.Detail) {

	dao.DB.Model(&detail).Updates(detail)
}

// @Title [Authorize needed] DeleteDetail
// @Description  Delete Detail
func DeleteDetail(ID string) {
	detail := models.Detail{}
	dao.DB.First(&detail, ID)
	exist, _ := tool.IsFileExist(detail.Cover)
	if exist {
		_ = tool.RemoveFile(detail.Cover)
	}
	dao.DB.Delete(&models.Detail{}, ID)
}

// @Title [Authorize needed] GetDetail
// @Description  返回文章给前端
func GetDetail() []models.Detail {
	var detail []models.Detail
	dao.DB.Find(&detail)
	return detail
}

// @Title [Authorize needed] IsDisplay
// @Description  更改是否显示
func IsDisplay(ID string, display string) {

	dao.DB.Model(&models.Detail{}).Where("id = ?", ID).Update("display", display)
}

// @Title [Authorize needed] DisplayDetail
// @Description  返回需要显示的动态
// @router /admin [post]
func DisplayDetail() []models.Detail {
	var detail []models.Detail
	dao.DB.Where("display = ?", "1").Find(&detail)
	return detail
}

// @Title [Authorize needed] GetDetailById
// @Description  通过id查找文章
func GetDetailById(id string) models.Detail {
	detail := models.Detail{}
	dao.DB.First(&detail, id)
	return detail
}

// @Title [Authorize needed] GetRelevantDetail
// @Description  find previous article and next article
func GetRelevantDetail(id string) []models.Detail {
	var details []models.Detail
	var previousDetail models.Detail
	var nextDetail models.Detail
	dao.DB.Where("id < ?", id).Last(&previousDetail)
	dao.DB.Where("id > ?", id).First(&nextDetail)
	details = append(details, previousDetail, nextDetail)
	return details
}
