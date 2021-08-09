import service from '@/utils/request'

// @Tags Sport
// @Summary 创建Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "创建Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport/createSport [post]
export const createSport = (data) => {
  return service({
    url: '/sport/createSport',
    method: 'post',
    data
  })
}

// @Tags Sport
// @Summary 删除Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "删除Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sport/deleteSport [delete]
export const deleteSport = (data) => {
  return service({
    url: '/sport/deleteSport',
    method: 'delete',
    data
  })
}

// @Tags Sport
// @Summary 删除Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sport/deleteSport [delete]
export const deleteSportByIds = (data) => {
  return service({
    url: '/sport/deleteSportByIds',
    method: 'delete',
    data
  })
}

// @Tags Sport
// @Summary 更新Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "更新Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sport/updateSport [put]
export const updateSport = (data) => {
  return service({
    url: '/sport/updateSport',
    method: 'put',
    data
  })
}

// @Tags Sport
// @Summary 更新SportByHash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "用H查询Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sport/updateSport [put]
export const findSportByHash = (params) => {
  return service({
    url: '/sport/findSportByHash',
    method: 'get',
    params
  })
}

// @Tags Sport
// @Summary 用id查询Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "用id查询Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sport/findSport [get]
export const findSport = (params) => {
  return service({
    url: '/sport/findSport',
    method: 'get',
    params
  })
}

// @Tags Sport
// @Summary 分页获取Sport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Sport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport/getSportList [get]
export const getSportList = (params) => {
  return service({
    url: '/sport/getSportList',
    method: 'get',
    params
  })
}