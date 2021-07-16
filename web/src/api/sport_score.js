import service from '@/utils/request'

// @Tags Sport_score
// @Summary 创建Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "创建Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport_score/createSport_score [post]
export const createSport_score = (data) => {
  return service({
    url: '/sport_score/createSport_score',
    method: 'post',
    data
  })
}

// @Tags Sport_score
// @Summary 删除Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "删除Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sport_score/deleteSport_score [delete]
export const deleteSport_score = (data) => {
  return service({
    url: '/sport_score/deleteSport_score',
    method: 'delete',
    data
  })
}

// @Tags Sport_score
// @Summary 删除Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sport_score/deleteSport_score [delete]
export const deleteSport_scoreByIds = (data) => {
  return service({
    url: '/sport_score/deleteSport_scoreByIds',
    method: 'delete',
    data
  })
}

// @Tags Sport_score
// @Summary 更新Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "更新Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sport_score/updateSport_score [put]
export const updateSport_score = (data) => {
  return service({
    url: '/sport_score/updateSport_score',
    method: 'put',
    data
  })
}

// @Tags Sport_score
// @Summary 用id查询Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "用id查询Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sport_score/findSport_score [get]
export const findSport_score = (params) => {
  return service({
    url: '/sport_score/findSport_score',
    method: 'get',
    params
  })
}

// @Tags Sport_score
// @Summary 分页获取Sport_score列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Sport_score列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport_score/getSport_scoreList [get]
export const getSport_scoreList = (params) => {
  return service({
    url: '/sport_score/getSport_scoreList',
    method: 'get',
    params
  })
}