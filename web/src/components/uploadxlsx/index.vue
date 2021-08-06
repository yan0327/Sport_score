<template>
  <div class="hello">
    <div>
      <el-button id="import" type="primary"  class="clearfix" @click="selectFile">导入</el-button>
      <el-button id="export" type="primary"  class="clearfix" @click="dialogFormVisible=true">导出</el-button>
    </div>
    <el-table id="table" :data="tableData" border style="width: 100%" >
      <template v-if="title.length>0">
        <el-table-column
          :prop="item"
          :label="item"
          v-for="item in title"
          :key="item"
          align="center"
        ></el-table-column>
      </template>
    </el-table>
    <el-dialog title="导出参数" :visible.sync="dialogFormVisible" width="30%">
      <el-form :model="form" :rules="rules">
        <el-form-item label="文件名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="form.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="后缀名" :label-width="formLabelWidth" prop="suffix">
          <el-select v-model="form.suffix" placeholder="请选择文件后缀名">
            <el-option label="xlsx" value="xlsx"></el-option>
            <el-option label="csv" value="csv"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="exportExcel">确 定</el-button>
      </div>
    </el-dialog>
    <input type="file" id="file" @change="importFile" />
  </div>
</template>

<script>
import {
  createSport,
} from '@/api/sport' //  此处请自行替换地址
import untils from "@/utils/excel";
export default {
  name: 'Uploadxlsx',
  props: {
    msg: String
  },
  data() {
    return {
      title: [],
      tableData: [
      ],
      dialogFormVisible: false,
      form: {
        name: "",
        suffix: "csv"
      },
      rules: {
        name: [{ required: true, message: "请输入文件名称", trigger: "blur" }]
      },
      formLabelWidth: "100px"
    };
  },
  methods: {
    exportExcel() {
      if (this.form.name == "") {
        this.$message.error("请先输入文件名");
        return false;
      }
      untils.exportFromTable({
        id: 'table',
        name: this.form.name,
        suffix: this.form.suffix,
        merges: [
        ]
      });
      this.form = {
        name: '',
        suffix: 'csv'
      };
      this.dialogFormVisible = false;
    },
    selectFile() {
      const el = document.getElementById("file")
      el.click()
    },
    async importFile() {
        let res
      const el = document.getElementById("file")
      const file = el.files[0]
      var data = await untils.importFromLocal(file)
      console.log(data)
      var n = 0;
      for (; n < data.length; n++) {
        res = await createSport(data[n])
      }
      this.title = data.title
      this.tableData = data.body
      el.value = ''
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
#export {
  float: right;
  margin: 12px 0;
}
#import {
  float: left;
  margin: 12px 0;
}
#file {
  display: none;
}
.el-select {
  width: 100%;
}
</style>
