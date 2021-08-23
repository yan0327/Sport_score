import service from '@/utils/request'

// @Tags SportStudentScore
// @Summary 创建SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "创建SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudentScore/createSportStudentScore [post]
export const createSportStudentScore = (data) => {
  return service({
    url: '/sportStudentScore/createSportStudentScore',
    method: 'post',
    data
  })
}

// @Tags SportStudentScore
// @Summary 删除SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "删除SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportStudentScore/deleteSportStudentScore [delete]
export const deleteSportStudentScore = (data) => {
  return service({
    url: '/sportStudentScore/deleteSportStudentScore',
    method: 'delete',
    data
  })
}

// @Tags SportStudentScore
// @Summary 删除SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportStudentScore/deleteSportStudentScore [delete]
export const deleteSportStudentScoreByIds = (data) => {
  return service({
    url: '/sportStudentScore/deleteSportStudentScoreByIds',
    method: 'delete',
    data
  })
}

// @Tags SportStudentScore
// @Summary 更新SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "更新SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportStudentScore/updateSportStudentScore [put]
export const updateSportStudentScore = (data) => {
  return service({
    url: '/sportStudentScore/updateSportStudentScore',
    method: 'put',
    data
  })
}

// @Tags SportStudentScore
// @Summary 用id查询SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "用id查询SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportStudentScore/findSportStudentScore [get]
export const findSportStudentScore = (params) => {
  return service({
    url: '/sportStudentScore/findSportStudentScore',
    method: 'get',
    params
  })
}

// @Tags SportStudentScore
// @Summary 分页获取SportStudentScore列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SportStudentScore列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudentScore/getSportStudentScoreList [get]
export const getSportStudentScoreList = (params) => {
  return service({
    url: '/sportStudentScore/getSportStudentScoreList',
    method: 'get',
    params
  })
}