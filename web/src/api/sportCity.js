import service from '@/utils/request'

// @Tags SportCity
// @Summary 创建SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "创建SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportCity/createSportCity [post]
export const createSportCity = (data) => {
  return service({
    url: '/sportCity/createSportCity',
    method: 'post',
    data
  })
}

// @Tags SportCity
// @Summary 删除SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "删除SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportCity/deleteSportCity [delete]
export const deleteSportCity = (data) => {
  return service({
    url: '/sportCity/deleteSportCity',
    method: 'delete',
    data
  })
}

// @Tags SportCity
// @Summary 删除SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportCity/deleteSportCity [delete]
export const deleteSportCityByIds = (data) => {
  return service({
    url: '/sportCity/deleteSportCityByIds',
    method: 'delete',
    data
  })
}

// @Tags SportCity
// @Summary 更新SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "更新SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportCity/updateSportCity [put]
export const updateSportCity = (data) => {
  return service({
    url: '/sportCity/updateSportCity',
    method: 'put',
    data
  })
}

// @Tags SportCity
// @Summary 用id查询SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "用id查询SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportCity/findSportCity [get]
export const findSportCity = (params) => {
  return service({
    url: '/sportCity/findSportCity',
    method: 'get',
    params
  })
}

// @Tags SportCity
// @Summary 分页获取SportCity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SportCity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportCity/getSportCityList [get]
export const getSportCityList = (params) => {
  return service({
    url: '/sportCity/getSportCityList',
    method: 'get',
    params
  })
}