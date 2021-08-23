import service from '@/utils/request'

// @Tags SportSchool
// @Summary 创建SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "创建SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportSchool/createSportSchool [post]
export const createSportSchool = (data) => {
  return service({
    url: '/sportSchool/createSportSchool',
    method: 'post',
    data
  })
}

// @Tags SportSchool
// @Summary 删除SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "删除SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportSchool/deleteSportSchool [delete]
export const deleteSportSchool = (data) => {
  return service({
    url: '/sportSchool/deleteSportSchool',
    method: 'delete',
    data
  })
}

// @Tags SportSchool
// @Summary 删除SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportSchool/deleteSportSchool [delete]
export const deleteSportSchoolByIds = (data) => {
  return service({
    url: '/sportSchool/deleteSportSchoolByIds',
    method: 'delete',
    data
  })
}

// @Tags SportSchool
// @Summary 更新SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "更新SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportSchool/updateSportSchool [put]
export const updateSportSchool = (data) => {
  return service({
    url: '/sportSchool/updateSportSchool',
    method: 'put',
    data
  })
}

// @Tags SportSchool
// @Summary 用id查询SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "用id查询SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportSchool/findSportSchool [get]
export const findSportSchool = (params) => {
  return service({
    url: '/sportSchool/findSportSchool',
    method: 'get',
    params
  })
}

// @Tags SportSchool
// @Summary 分页获取SportSchool列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SportSchool列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportSchool/getSportSchoolList [get]
export const getSportSchoolList = (params) => {
  return service({
    url: '/sportSchool/getSportSchoolList',
    method: 'get',
    params
  })
}