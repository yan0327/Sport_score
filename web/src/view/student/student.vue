<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-input placeholder="搜索条件" v-model="searchInfo.name" />
        </el-form-item>    
        <el-form-item label="年龄">
          <el-input placeholder="搜索条件" v-model="searchInfo.age" />
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
      <el-table-column label="姓名" prop="name" width="120" /> 
      <el-table-column label="年龄" prop="age" width="120" /> 
      <el-table-column label="性别" prop="sex" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.sex,"sex")}}
        </template>
      </el-table-column><el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateStudent(scope.row)">变更</el-button>
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
        <el-form-item label="姓名:">
      
          <el-input v-model="formData.name" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="年龄:">
      
          <el-input v-model.number="formData.age" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="性别:">
      
          <el-select v-model="formData.sex" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in sexOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createStudent,
  deleteStudent,
  deleteStudentByIds,
  updateStudent,
  findStudent,
  getStudentList
} from '@/api/student' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'Student',
  mixins: [infoList],
  data() {
    return {
      listApi: getStudentList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
      sexOptions: [],
          
      formData: {
        name: '',
          age: 0,
          sex: 0,
          
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
    
    await this.getDict('sex')
      
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10   
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
        this.deleteStudent(row)
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
      const res = await deleteStudentByIds({ ids })
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
    async updateStudent(row) {
      const res = await findStudent({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.restudent
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        name: '',
          age: 0,
          sex: 0,
          
      }
    },
    async deleteStudent(row) {
      const res = await deleteStudent({ ID: row.ID })
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
          res = await createStudent(this.formData)
          break
        case "update":
          res = await updateStudent(this.formData)
          break
        default:
          res = await createStudent(this.formData)
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
