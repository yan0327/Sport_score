<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="学校">
          <el-input placeholder="搜索条件" v-model="searchInfo.school" />
        </el-form-item>    
        <el-form-item label="班级">
          <el-input placeholder="搜索条件" v-model="searchInfo.class" />
        </el-form-item>    
        <el-form-item label="考号">
          <el-input placeholder="搜索条件" v-model="searchInfo.testid" />
        </el-form-item>    
        <el-form-item label="姓名">
          <el-input placeholder="搜索条件" v-model="searchInfo.name" />
        </el-form-item>    
            <el-form-item label="性别" prop="sex">
            <el-select v-model="searchInfo.sex" clear placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>   
        <el-form-item label="总分">
          <el-input placeholder="搜索条件" v-model="searchInfo.totalScore" />
        </el-form-item>    
        <el-form-item label="过程评价分">
          <el-input placeholder="搜索条件" v-model="searchInfo.processeValuation" />
        </el-form-item>    
        <el-form-item label="等级">
          <el-input placeholder="搜索条件" v-model="searchInfo.grade" />
        </el-form-item>    
        <el-form-item label="一类项目">
          <el-input placeholder="搜索条件" v-model="searchInfo.itemone" />
        </el-form-item>    
        <el-form-item label="成绩1">
          <el-input placeholder="搜索条件" v-model="searchInfo.gradeone" />
        </el-form-item>    
        <el-form-item label="分数1">
          <el-input placeholder="搜索条件" v-model="searchInfo.scoreOne" />
        </el-form-item>    
        <el-form-item label="二类项目">
          <el-input placeholder="搜索条件" v-model="searchInfo.itemTwo" />
        </el-form-item>        
        <el-form-item label="三类项目">
          <el-input placeholder="搜索条件" v-model="searchInfo.itemThree" />
        </el-form-item>      
        <el-form-item label="分数3">
          <el-input placeholder="搜索条件" v-model="searchInfo.scoreThree" />
        </el-form-item>    
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="学校" prop="school" width="120" /> 
      <el-table-column label="班级" prop="class" width="120" /> 
      <el-table-column label="考号" prop="testid" width="120" /> 
      <el-table-column label="姓名" prop="name" width="120" /> 
      <el-table-column label="性别" prop="sex" width="120">
        <template slot-scope="scope">{{scope.row.sex|formatBoolean}}</template>
      </el-table-column>
      <el-table-column label="总分" prop="totalScore" width="120" /> 
      <el-table-column label="过程评价分" prop="processeValuation" width="120" /> 
      <el-table-column label="等级" prop="grade" width="120" /> 
      <el-table-column label="一类项目" prop="itemone" width="120" /> 
      <el-table-column label="成绩1" prop="gradeone" width="120" /> 
      <el-table-column label="分数1" prop="scoreOne" width="120" /> 
      <el-table-column label="二类项目" prop="itemTwo" width="120" /> 
      <el-table-column label="成绩2" prop="gradeTwo" width="120" /> 
      <el-table-column label="分数2" prop="scoreTwo" width="120" /> 
      <el-table-column label="三类项目" prop="itemThree" width="120" /> 
      <el-table-column label="成绩3" prop="gradeThree" width="120" /> 
      <el-table-column label="分数3" prop="scoreThree" width="120" /> <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateSport_score(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学校:">
      
          <el-input v-model="formData.school" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="班级:">
      
          <el-input v-model="formData.class" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="考号:">
      
          <el-input v-model="formData.testid" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="姓名:">
      
          <el-input v-model="formData.name" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="性别:">
      
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.sex" clearable ></el-switch>
      </el-form-item>
        <el-form-item label="总分:">
      
          <el-input-number v-model="formData.totalScore" :precision="2" clearable />
       </el-form-item>
        <el-form-item label="过程评价分:">
      
          <el-input-number v-model="formData.processeValuation" :precision="2" clearable />
       </el-form-item>
        <el-form-item label="等级:">
      
          <el-input v-model="formData.grade" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="一类项目:">
      
          <el-input v-model="formData.itemone" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="成绩1:">
      
          <el-input-number v-model="formData.gradeone" :precision="2" clearable />
       </el-form-item>
        <el-form-item label="分数1:">
      
          <el-input-number v-model="formData.scoreOne" :precision="2" clearable />
       </el-form-item>
        <el-form-item label="二类项目:">
      
          <el-input v-model="formData.itemTwo" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="成绩2:">
      
          <el-input-number v-model="formData.gradeTwo" :precision="2" clearable />
       </el-form-item>
        <el-form-item label="分数2:">
      
          <el-input-number v-model="formData.scoreTwo" :precision="2" clearable />
       </el-form-item>
        <el-form-item label="三类项目:">
      
          <el-input v-model="formData.itemThree" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="成绩3:">
      
          <el-input-number v-model="formData.gradeThree" :precision="2" clearable />
       </el-form-item>
        <el-form-item label="分数3:">
      
          <el-input-number v-model="formData.scoreThree" :precision="2" clearable />
       </el-form-item>
     </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createSport_score,
  deleteSport_score,
  deleteSport_scoreByIds,
  updateSport_score,
  findSport_score,
  getSport_scoreList
} from '@/api/sport_score' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'Sport_score',
  mixins: [infoList],
  data() {
    return {
      listApi: getSport_scoreList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
      formData: {
        school: '',
          class: '',
          testid: '',
          name: '',
          sex: false,
          totalScore: 0,
          processeValuation: 0,
          grade: '',
          itemone: '',
          gradeone: 0,
          scoreOne: 0,
          itemTwo: '',
          gradeTwo: 0,
          scoreTwo: 0,
          itemThree: '',
          gradeThree: 0,
          scoreThree: 0,
          
      }
    }
  },
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time);
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss');
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  async created() {
    await this.getTableData()
    
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10    
      if (this.searchInfo.sex === ""){
        this.searchInfo.sex=null
      }              
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteSport_score(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteSport_scoreByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSport_score(row) {
      const res = await findSport_score({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.resport_score
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        school: '',
          class: '',
          testid: '',
          name: '',
          sex: false,
          totalScore: 0,
          processeValuation: 0,
          grade: '',
          itemone: '',
          gradeone: 0,
          scoreOne: 0,
          itemTwo: '',
          gradeTwo: 0,
          scoreTwo: 0,
          itemThree: '',
          gradeThree: 0,
          scoreThree: 0,
          
      }
    },
    async deleteSport_score(row) {
      const res = await deleteSport_score({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1 ) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case "create":
          res = await createSport_score(this.formData)
          break
        case "update":
          res = await updateSport_score(this.formData)
          break
        default:
          res = await createSport_score(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
