<template>
  <div class="app-container">

    <el-table
      border
      stripe
      v-loading="listLoading"
      :data="users"
      @selection-change="handleSelectionChange">

      <el-table-column
        type="index"
        label="序号"
        width="80"
        align="center"
      />

      <el-table-column prop="user_id" label="账号" width="150" />
      <el-table-column prop="user_name" label="账号名" width="150"/>
      <el-table-column prop="phone" label="电话号码" width="180"/>
      <el-table-column prop="addresss" label="地址">
        <template slot-scope="{row}">
          <div>
            <ul>
              <li  v-for="(item,id) in row.addresss" :key="id" >
                <span>收货人：{{item.user_name}}</span>
                <br>
                <span>收货人电话号码：{{item.phone}}</span>
                <br>
                <span>收货地址：{{item.address}}</span>
              </li>
            </ul>
          </div>
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
      @current-change="getUsers"
      @size-change="handleSizeChange"
    />
  </div>
</template>

<script>
import cloneDeep from 'lodash/cloneDeep'
import axios from 'axios'
import { Message } from 'element-ui';

export default {
  name: 'UserList',

  data () {
    return {
      listLoading: false, // 是否显示列表加载的提示
      searchObj: { // 包含请求搜索条件数据的对象
        user_id: 1
      },
      tempSearchObj: { // 收集搜索条件输入的对象
        user_id: 1
      },
      users: [], // 当前页的用户列表
      page: 1, // 当前页码
      limit: 10, // 每页数量
      total: 0, // 总数量
    }
  },

  created () {
    this.getUsers()
  },

  methods: {
    /* 
    根据输入进行搜索
    */
    handleSelectionChange(){

    },

    getUsers (page=1) {
      this.page = page
      const {limit} = this
      this.listLoading = true     
      const url='/api1/v2/admin/showUser'
      
      axios.post(url, {page:this.page,num_each_page:limit})
        .then(response => {
          // 处理响应数据\
          if(response.data.status === 200){
            this.listLoading = false
            this.users = response.data.data.userlist
            this.total = response.data.data.count
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
          Message.error('请求用户失败！')
        })   
    },

    /* 
    处理pageSize发生改变的监听回调
    */
    handleSizeChange (pageSize) {
      this.limit = pageSize
      this.getUsers()
    },

  }
}

</script>
