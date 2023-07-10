<template>
  <div class="app-container">
    <!-- 订单搜索 -->
    <el-form inline>
      <el-form-item>
         <el-input v-model="tempSearchOrder_id" placeholder="订单编号" />
      </el-form-item>

      <el-button type="primary" icon="el-icon-search" @click="search">查询</el-button>
      <el-button type="default" @click="resetSearch">清空</el-button>
    </el-form>

    <el-radio-group v-model="srchStatus" >
      <el-radio  label="" >全部</el-radio>
      <el-radio  label="未支付" >未支付</el-radio>
      <el-radio  label="已支付" >已支付</el-radio>
      <el-radio  label="送餐中" >送餐中</el-radio>
      <el-radio  label="已送达" >已送达</el-radio>
    </el-radio-group>

    

    <div style="margin-bottom: 20px">
      <el-button type="primary" @click="inputDeliver">指派配送员</el-button>
    </div>

    <el-table
      border
      stripe
      v-loading="listLoading"
      :data="orders"
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
      <el-table-column prop="order_id" label="订单编号" width="150" />
      <el-table-column prop="created_at" label="创建时间" width="180" :formatter="formatDate"/>
      <el-table-column prop="updated_at" label="更新时间" width="180" :formatter="formatDate"/>
      <el-table-column prop="order_productlist" label="订单详情" width="150" >
        <template slot-scope="{row}">
            <div>
              <ul>
                <li  v-for="(item,id) in row.order_productlist" :key="id" >                 
                  <span>{{item.product_name}}*{{item.num}}</span>
                  <br>
                </li>
              </ul>
            </div>
        </template>
      </el-table-column>
      <el-table-column prop="user_name" label="收货人" width="150" />
      <el-table-column prop="address" label="收货地址" width="150" />
      <el-table-column prop="user_phone" label="收货人电话号码" width="150" />
      <el-table-column prop="status" label="订单状态" width="150" />
      <el-table-column prop="deliver_id" label="配送员编号" width="150" />
      <el-table-column prop="deliver_name" label="配送员姓名" width="150" />
      <el-table-column prop="deliver_phone" label="配送员电话号码" width="150" />
            
    </el-table>

    <el-pagination
      :current-page="page"
      :total="total"
      :page-size="limit"
      :page-sizes="[3, 10, 20, 30, 40, 50, 100]"
      style="padding: 20px 0;"
      layout="prev, pager, next, jumper, ->, sizes, total"
      @current-change="getOrders"
      @size-change="handleSizeChange"
    />
    <!-- 指派配送员的弹窗 -->
    <el-dialog title='指派配送员' :visible.sync="distributeVisible">
      <template>
        <div class="custom-input">
          <h3>请输入配送员的编号</h3>
          <el-input
            v-model="deliver_id"
            v-on="$listeners"
            :class="{ 'custom-input__error': showError }"
            :show-word-limit="showLimit"
            clearable
            size="medium"
            prepend-icon="el-icon-edit"
            placeholder='id'
            @input="handleInput"
            @focus="handleFocus"
            @blur="handleBlur"
          />
          <span v-show="showError" class="custom-input__message">{{ errorMessage }}</span>
        </div>
      </template>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button  type="primary" @click="distribute">确 定</el-button>
      </div>
    </el-dialog>
    
  </div>
</template>

<script>
import cloneDeep from 'lodash/cloneDeep'
import { Message } from 'element-ui'
import Axios from 'axios'
import { ElRadio } from 'element-ui';
export default {
  name: 'OrderList',
  props:{
    errorMessage: {
        type: String,
        default: '输入内容不能为空'
      }
  },

  data () {
    return {
      srchStatus:'',
      distributeVisible: false,
      listLoading: false, // 是否显示列表加载的提示
      deliver_id:'',
      orders: [], // 当前页的订单列表
      selectedIds:[],
      page: 1, // 当前页码
      limit: 10, // 每页数量
      total: 0, // 总数量
      searchOrder_id:null,
      tempSearchOrder_id:null,
      showError: false,
      showLimit: false
      
    }
  },

  created () {
    this.getOrders()
  },
  watch:{
    srchStatus:{
      handler(newvalue,oldvalue){
        this.getOrders()
      }
    },
    immediate: true,
  },
  methods: {
    formatDate(row, column, cellValue) {
      const date = new Date(cellValue)
      return date.toLocaleString('zh-CN', { hour12: false })
    },
    // 显示指派配送员的dialog
    inputDeliver(){
        if(this.selectedIds.length != 0){
          this.distributeVisible = true
        }
        else{
          Message.error('请勾选订单！')
        }
    },
    distribute(){
      const tmp = parseInt(this.deliver_id)
      this.deliver_id = tmp
      const url='/api1/v2/admin/assignDelivery'
      Axios.post(url, {order_id_list:this.selectedIds,deliver_id:this.deliver_id})
        .then(response => {
          // 处理响应数据
          if(response.data.status === 200){
            this.selectedIds = []
            Message.success('指派配送员成功！')
            this.distributeVisible = false
            this.deliver_id=''
            this.getOrders()
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg+'请重新输入')
          }
          else if(response.data.status === 404){
            Message.error('登录已过期，请重新登录')
            this.$store.dispatch(logout)
          }
        })
        .catch(error => {
          // 处理请求错误
          Message.error('请求订单失败！')
        })
    },
    /* 
    根据输入进行搜索
    */
    search () {
      const num = parseInt(this.tempSearchOrder_id)
      this.searchOrder_id = num
      this.getOrders()
    },

    /* 
    重置输入后搜索
    */
    resetSearch () {
      this.searchOrder_id = null
      this.tempSearchOrder_id = null
      this.getOrders()
    },

    /* 
    列表选中状态发生改变的监听回调
    */
    handleSelectionChange (selection) {
      this.selectedIds = selection.map(item => item.order_id)
    },
    handlegetOrders(){
      this.getOrders()
    },
    /* 
    获取分页列表
    */
    getOrders (page=1) {
      this.page = page
      this.listLoading = true
      const tmp = parseInt(this.tempSearchOrder_id)
      this.searchOrder_id = tmp
      const url = '/api1/v2/admin/showOrder'
      Axios.post(url, {page:this.page, num_each_page:this.limit, canteen_id:this.$store.state.canteen_id, status:this.srchStatus,search_order_id:this.searchOrder_id})
        .then(response => {
          // 处理响应数据
          if(response.data.status === 200){
            this.listLoading = false
            this.orders = response.data.data.orderlist
            this.total = response.data.data.count
            this.selectedIds = []
            this.searchOrder_id = null
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            Message.error('登录已过期，请重新登录')
            this.$store.dispatch(logout)
          }
        })
        .catch(error => {
          // 处理请求错误
          Message.error('请求订单失败！')
        })   

    },

    /* 
    处理pageSize发生改变的监听回调
    */
    handleSizeChange (pageSize) {
      this.limit = pageSize
      this.getOrders()
    },

    /* 
    取消指派配送员
    */
    cancel () {
      this.deliver_id = ''
      this.distributeVisible = false
    },
    handleInput(value) {
      this.$emit('input', value)
      this.showError = value.trim() === ''
    },
    handleFocus() {
      this.showLimit = true
      this.showError = false
    },
    handleBlur() {
      this.showLimit = false
      this.showError = this.value.trim() === ''
    }
  }
}
</script>
<style scoped>
.custom-input {
  position: relative;
}
.custom-input__error {
  border-color: #f56c6c !important;
}
.custom-input__icon {
  margin-left: 5px;
  color: #999;
  cursor: pointer;
}
.custom-input__icon--active {
  color: #409eff;
}
.custom-input__message {
  position: absolute;
  bottom: -20px;
  left: 0;
  font-size: 12px;
  color: #f56c6c;
}
.custom-input__tooltip {
  margin-left: 5px;
  cursor: pointer;
}
</style>