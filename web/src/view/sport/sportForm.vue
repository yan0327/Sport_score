<template>
  <div>
      <uploadxlsx />
    <el-form :model="formData" label-position="right" label-width="160px">
      
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
    <el-input v-model="formData.sex" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="总分:">
    
      <el-input-number v-model="formData.totalScore" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="过程评价分:">
    
      <el-input-number v-model="formData.processeValuation" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="等级:">
    <el-input v-model="formData.grade" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="一类项目:">
    <el-input v-model="formData.itemone" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="成绩1:">
    
      <el-input-number v-model="formData.gradeone" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="分数1:">
    
      <el-input-number v-model="formData.scoreOne" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="二类项目:">
    <el-input v-model="formData.itemTwo" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="成绩2:">
    
      <el-input-number v-model="formData.gradeTwo" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="分数2:">
    
      <el-input-number v-model="formData.scoreTwo" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="三类项目:">
    <el-input v-model="formData.itemThree" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="成绩3:">
    
      <el-input-number v-model="formData.gradeThree" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="分数3:">
    
      <el-input-number v-model="formData.scoreThree" :precision="2" clearable></el-input-number>
    </el-form-item>

    <el-form-item label="体育成绩哈希值:">
    <el-input v-model="formData.hash256" clearable placeholder="请输入" />
    </el-form-item>

    <el-form-item label="交易哈希:">
    <el-input v-model="formData.transhash" clearable placeholder="请输入" />
    </el-form-item>

    <el-form-item>
        <el-button size="mini" type="primary" @click="save">保存</el-button>
        <el-button size="mini" type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  createSport,
  updateSport,
  findSport
} from '@/api/sport' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import uploadxlsx from '@/components/uploadxlsx'
import Web3 from "web3";
import EthereumTx from "ethereumjs-tx";

export default {
  name: 'Sport',
  mixins: [infoList],
  components:{
    uploadxlsx
  },
  data() {
    return {
      type: '',
      msg: "This is demo net work",
      formData: {
            school: '',
            class: '',
            testid: '',
            name: '',
            sex: '',
            totalScore: 0,
            processeValuation: 0,
            grade: '',
            itemone: '50米跑',
            gradeone: 0,
            scoreOne: 0,
            itemTwo: '跳绳',
            gradeTwo: 0,
            scoreTwo: 0,
            itemThree: '篮球运球',
            gradeThree: 0,
            scoreThree: 0,
            hash256: '',
            transhash: '',
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findSport({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.resport
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createSport(this.formData)
          break
        case 'update':
          res = await updateSport(this.formData)
          break
        default:
          res = await createSport(this.formData)
          break
      }
      console.log(this.formData)
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    },
    onComplete(data){
      console.log('data', data);
    }
  },
  async mounted() {
    web3 = new Web3(new Web3.providers.HttpProvider("https://kovan.infura.io/v3/e6b151ba42004b5ebce395b52fa4de91"));
        //智能合约地址

    /*
    if (window.ethereum) {
      window.web3 = new Web3(ethereum);
      try {
        const accounts = await ethereum.enable();
        console.log(accounts);
        const provider = window['ethereum']
        console.log(provider)
        console.log(provider.chainId)
        //this.web3TimerCheck(window.web3);
        const web3 = new Web3(provider)
        console.log(web3)
        window.defaultAccount = accounts[0].toLowerCase()
        console.log(window.defaultAccount)
        
        var abi2 = require("./contract_abi2.json");
        const address2 = "0xb1d17b075d13ee1ec7a686d692809182ac9f19f0"
        window.myContract2 = new web3.eth.Contract(abi2.abi, address2)
        console.log(window.myContract2)


      } catch (error) {
        console.log(error)
      }
    } else if (window.web3) {
      window.web3 = new Web3(web3.currentProvider);
      //this.web3TimerCheck(window.web3);
    } else {
      this.web3 = null;
      this.Log(this.MetamaskMsg.METAMASK_NOT_INSTALL, "NO_INSTALL_METAMASK");
      console.error(
        "Non-Ethereum browser detected. You should consider trying MetaMask!"
      );
    }*/
  
  },
}
</script>

<style>
</style>