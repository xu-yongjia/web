<template>
  <div class="app-container">
    <el-form inline>
      <el-form-item>
         <el-input v-model="tempTruename_or_id" placeholder="配送员id或姓名" />
      </el-form-item>

      <el-button type="primary" icon="el-icon-search" @click="search">查询</el-button>
      <el-button type="default" @click="resetSearch">清空</el-button>
    </el-form>

    <div style="margin-bottom: 20px">
      <el-button type="primary" @click="showAddDeliver">添 加</el-button>
      <el-popconfirm :title="`确定删除吗?`" @onConfirm="removeDelivers()">
            <el-button :disabled="selectedIDs.length===0" slot="reference" type="danger" >批量删除</el-button>/>
      </el-popconfirm>
    </div>

    <el-table
      border
      stripe
      v-loading="listLoading"
      :data="delivers"
      @selection-change="handleSelectionChange">

      <el-table-column
        type="selection"
        width="55" />

      <el-table-column
        type="index"
        label="序号"
        width="80"
        align="center"
      />
      <el-table-column prop="deliver_id" label="配送员id" width="180" />
      <el-table-column prop="truename" label="配送员姓名" width="180" />
      <el-table-column prop="phone" label="电话号码" width="width"/>

      <el-table-column label="操作" width="230" align="center">
        <template slot-scope="{row}">
          <MyButton type="primary" size="mini" icon="el-icon-edit" title="修改配送员"
            @click="showUpdateDeliver(row)"/>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :total="total"
      :page-size="limit"
      :page-sizes="[3, 10, 20, 30, 40, 50, 100]"
      style="padding: 20px 0;"
      layout="prev, pager, next, jumper, ->, sizes, total"
      @current-change="getDelivers"
      @size-change="handleSizeChange"
    />

    <el-dialog :title="deliver.deliver_id ? '修改配餐员' : '添加配餐员'" :visible.sync="dialogDeliverVisible">
      <el-form ref="deliverForm" :model="deliver" :rules="deliverRules" label-width="120px">
        <!-- 若添加，不显示id，后端自动分配id -->
        <el-form-item v-if="deliver.deliver_id" label="配送员id" prop="deliver_id" >
          <el-input v-model="deliver.deliver_id" :readonly="true"/>
        </el-form-item>
        <el-form-item label="配送员姓名" prop="truename">
          <el-input v-model="deliver.truename"/>
        </el-form-item>
        <el-form-item label="电话号码">
          <el-input v-model="deliver.phone"/>
        </el-form-item>
        <el-form-item label="所属食堂" >
          <el-input v-model="canteenName" :readonly="true"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button :loading="loading" type="primary" @click="addOrUpdate">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import cloneDeep from 'lodash/cloneDeep'
import axios from 'axios'
import { Message } from 'element-ui'

export default {
  name: 'DeliverList',
  data () {
    return {
      canteenName:'',
      listLoading: false, // 是否显示列表加载的提示
      truename_or_id:'', // 包含请求搜索条件数据的对象
        
      tempTruename_or_id: '', // 收集搜索条件输入的对象
        
      selectedIDs: [], // 所有选择的deliver的id的数组
      delivers: [], // 当前页的delivers列表
      page: 1, // 当前页码
      limit: 10, // 每页数量
      total: 0, // 总数量
      deliver: {}, // 当前要操作的deliver
      dialogDeliverVisible: false, // 是否显示添加/修改的dialog
      deliverRules: { // 添加/修改表单的校验规则
        phone: [
          { required: true, validator: this.validatePhone }
        ],
        truename: [
          { required: true, message: '配送员姓名必须输入' },
        ],
      },
      loading: false, // 是否正在提交请求中
      isIndeterminate: false, // 是否是不确定的
      checkAll: false, // 是否全选
    }
  },
  mounted(){
    this.setCanteenName()
    this.getDelivers()
  },
  //发请求一般情况下，我们都是在mounted去发，但是也可以在created内部去发
  created () {
  },

  methods: {
    setCanteenName(){
      if(this.$store.state.canteen_id === 1){
        this.canteenName = '榕园食堂'
      }
      else if(this.$store.state.canteen_id === 2){
        this.canteenName = '岁月湖食堂'
      }
      else if(this.$store.state.canteen_id === 3){
        this.canteenName = '荔园食堂'
      }
      else if(this.$store.state.canteen_id === 4){
        this.canteenName = '若海食堂'
      }
      else if(this.$store.state.canteen_id === 5){
        this.canteenName = '瑾园食堂'
      }
    },
    /* 
    自定义电话号码校验
    */
    validatePhone (rule, value, callback) {
      const reg = /^1[3|4|5|6|7|8|9][0-9]\d{8}$/
        if (!value) {
          return callback(new Error('电话号码不能为空'))
        }
        if (!reg.test(value)) {
          return callback(new Error('请输入正确的电话号码'))
        }
        callback()
    },
    /* 
    根据输入进行搜索
    */
    search () {
      this.truename_or_id = this.tempTruename_or_id
      this.getDelivers()
    },

    /* 
    重置输入后搜索
    */
    resetSearch () {
      this.truename_or_id = ''
      this.tempTruename_or_id = ''
      this.getDelivers()
    },

    /* 
    ???显示添加用户的界面
    */
    showAddDeliver () {
      this.deliver = {}
      //显示“添加配送员”的对话框
      this.dialogDeliverVisible = true
      //将回调延迟到 DOM 更新之后执行，从而保证在更新后的 DOM 中操作组件
      this.$nextTick(() => this.$refs.deliverForm.clearValidate())
    },

    /* 
    删除所有选中的用户
    */
    removeDelivers (deliver) {
      const url='/api1/v2/admin/delivery/del'
      axios.post(url, {deliver_id:this.selectedIDs})
        .then((response) => {
          // 处理响应数据
          if(response.data.status === 200){
            Message.success('删除成功！')
            this.getDelivers()
            this.selectedIDs = []

          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            Message.error('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
          })
          .catch(error => {
          Message.error('删除配送员失败！')
          })
      },

    /* 
    列表选中状态发生改变的监听回调
    */
    handleSelectionChange (selection) {
      this.selectedIDs = selection.map(item => item.deliver_id)
    },

    /* 
    显示“更新用户”的界面
    */
    showUpdateDeliver (deliver) {
      this.deliver = cloneDeep(deliver)
      this.dialogDeliverVisible = true
    },
    /* 
    获取分页列表
    */
   //传参第几页
   //limit是每一页展示几个
    getDelivers (page=1) {
      this.page = page
      this.listLoading = true     
      const url='/api1/v2/admin/showDelivery'
      
      axios.post(url, {page:this.page,num_each_page:this.limit,canteen_id:this.$store.state.canteen_id,truename_or_id:this.truename_or_id})
        .then((response) => {
          // 处理响应数据
          if(response.data.status === 200){
            this.listLoading = false
            this.total = response.data.data.count
            const temp = response.data.data.deliverlist
            const result = temp.map(obj => {
              const { canteen_id, ...rest } = obj // 使用对象解构语法，将 name 属性值提取出来
              return { canteen_id: this.canteenName, ...rest } // 修改 name 属性值并返回新对象
              }
            )
            this.delivers = result
            this.selectedIDs = []
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            Message.error('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
        })
        .catch(error => {
          // 处理请求错误
          Message.error('请求配送员失败！')
        })   
    },

    /* 
    处理pageSize发生改变的监听回调
    */
    handleSizeChange (pageSize) {
      this.limit = pageSize
      this.getDelivers()
    },

    /* 
    取消用户的保存或更新
    */
    cancel () {
      this.dialogDeliverVisible = false
      this.deliver = {}
    },

    /* 
    保存或者更新用户
    */
    addOrUpdate () {
      const {deliver} = this
      this.loading = true
      //deliver_id存在，则是更新
      if(deliver.deliver_id){
        const updateURL = '/api1/v2/admin/delivery/update'
        axios.post(updateURL,{deliver_id:this.deliver.deliver_id,truename:this.deliver.truename,phone:this.deliver.phone,canteen_id:this.$store.state.canteen_id})
        .then((response) => {
          // 处理响应数据
          if(response.data.status === 200){
          // 处理响应数据
          this.loading = false
          this.$message.success('更新成功!')
          this.getDelivers(deliver.deliver_id ? this.page : 1)
          this.deliver = {}
          this.dialogDeliverVisible = false
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            Message.error('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
        }
        ).catch(error => {
          // 处理请求错误
          Message.error('修改失败!')
        })
      }
      else{
        const addURL = '/api1/v2/admin/delivery/add'
        axios.post(addURL,{canteen_id:this.$store.state.canteen_id,truename:this.deliver.truename,phone:this.deliver.phone,})
        .then((response) => {
          if(response.data.status === 200){
          // 处理响应数据
          this.loading = false
          this.$message.success('添加成功!')
          this.getDelivers(deliver.deliver_id ? this.page : 1)
          this.deliver = {}
          this.dialogDeliverVisible = false
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            Message.error('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
        }
        ).catch(error => {
          // 处理请求错误
          Message.error('添加失败!')
        })       
      }
    }
  }
}
  


</script>
