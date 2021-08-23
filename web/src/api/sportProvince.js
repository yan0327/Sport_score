import service from '@/utils/request'

// @Tags SportProvince
// @Summary 创建SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "创建SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportProvince/createSportProvince [post]
export const createSportProvince = (data) => {
  return service({
    url: '/sportProvince/createSportProvince',
    method: 'post',
    data
  })
}

// @Tags SportProvince
// @Summary 删除SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "删除SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportProvince/deleteSportProvince [delete]
export const deleteSportProvince = (data) => {
  return service({
    url: '/sportProvince/deleteSportProvince',
    method: 'delete',
    data
  })
}

// @Tags SportProvince
// @Summary 删除SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportProvince/deleteSportProvince [delete]
export const deleteSportProvinceByIds = (data) => {
  return service({
    url: '/sportProvince/deleteSportProvinceByIds',
    method: 'delete',
    data
  })
}

// @Tags SportProvince
// @Summary 更新SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "更新SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportProvince/updateSportProvince [put]
export const updateSportProvince = (data) => {
  return service({
    url: '/sportProvince/updateSportProvince',
    method: 'put',
    data
  })
}

// @Tags SportProvince
// @Summary 用id查询SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "用id查询SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportProvince/findSportProvince [get]
export const findSportProvince = (params) => {
  return service({
    url: '/sportProvince/findSportProvince',
    method: 'get',
    params
  })
}

// @Tags SportProvince
// @Summary 分页获取SportProvince列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SportProvince列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportProvince/getSportProvinceList [get]
export const getSportProvinceList = (params) => {
  return service({
    url: '/sportProvince/getSportProvinceList',
    method: 'get',
    params
  })
}