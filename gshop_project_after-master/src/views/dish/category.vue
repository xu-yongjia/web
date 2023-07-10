<template>
  <div>
    <!-- 添加按钮 -->
    <el-button type="primary" icon="el-icon-plus" @click="showDialog" >添加</el-button>
    <el-table 
      :data="categoryList"
      border style="width: 100%; margin: 20px 0">
      <el-table-column
        label="序号"
        align="center"
        type="index"
        width="80"
      >
      </el-table-column>

      <el-table-column prop="category_name" label="种类名称" width="width">
      </el-table-column>


      <el-table-column prop="prop" label="操作" width="width">
          <!-- 由于要为每一个数据绑定,所以要循环,这样子才能传递数据 -->
          <template slot-scope="{row}">
            <!-- 传递当前循环的一个数据给函数 -->
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="showEditDialog(row)">编辑</el-button>
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="removeDialog(row)">删除</el-button>
            <!-- <router-link :to=""> -->
            <el-button type="primary" icon="el-icon-view" size="mini" @click="viewProducts(row.category_id)">查看菜品</el-button>
            <!-- </router-link> -->
          </template>
      </el-table-column>
    </el-table>

    <!-- 添加不需要id的,而修改是需要id的 -->
    <el-dialog :title="category.category_id?'修改种类':'添加种类'" :visible.sync="addDialogVisible">
      <!-- el-form通过:model="form" form为一个对象,用于将表单所有数据进行收集 -->
      <el-form :model="category" :rules="rules" style="width:80%" ref="category">
        <!-- 种类id -->
        <el-form-item v-if="category.category_id" label="种类id" prop="category_id">
          <el-input v-model="category.category_id" :readonly="true"/>
        </el-form-item>
        <!-- 种类名称 -->
        <el-form-item label="种类名称" prop="category_name">
          <el-input v-model="category.category_name" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="confirmAddOrEdit">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Axios from 'axios';
import { Message } from 'element-ui';
export default {
  name: "Category",
  data() {
    return {
      //当前页
      page: 1,
      //每页显示数据
      limit: 3,
      //数据列表
      categoryList: [],
      //控制对话框的显示隐藏
      addDialogVisible:false,
      //收集表单数据
      category:{
        //种类名称
        category_id:null,
        category_name:''
      },
      //验证规则(由对象组成)
      rules:{
        // 种类名称校验
        category_name:[
          {required:true,message:"请填写种类!",trigger:'change'}
        ],
      }
    };
  },
  mounted() {
    this.getCategory();
  },
  methods: {
    // 用户单击删除
    removeDialog(info){
      this.$confirm(`确认删除${info.category_name}吗?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const url = '/api1/v2/admin/category/hasDel'
        Axios.post(url, {category_id:info.category_id,canteen_id:this.$store.state.canteen_id})
        .then(response => {
          // 处理响应数据
          if(response.data.status === 200){
            Message.success("删除种类成功！")
            this.getCategory()
          }
          else if(response.data.status === 201){
              this.$confirm("该类别中包含菜品信息，是否同步删除",'提示',{
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                const url2='/api1/v2/admin/category/del'
                Axios.post(url2, {category_id:info.category_id,canteen_id:this.$store.state.canteen_id})
                .then(response => {
                  // 处理响应数据
                  if(response.data.status === 200){
                    Message.success("该种类的所有菜品删除成功！")
                    this.getCategory()
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
                  Message.error('失败！')
                  })
              }).catch(() => {
                  this.$message({
                    type: 'info',
                    message: '已取消删除'
                  })
            })
          }
          else if(response.data.status === 404){
            //token过期
            Message.error('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
          })
          .catch(error => {
          // 处理请求错误
          Message.error('失败！')
          })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });          
      });
    },
    //查看该类别所有菜品
    viewProducts(id){
      this.$router.push({
              name:'Product',
              query:{
                category_id:id
              },
            })
    },
    // 用户单击确认添加/修改
    confirmAddOrEdit(){
      const {category} = this
      //deliver_id存在，则是更新
      if(category.category_id){
        const updateURL = '/api1/v2/admin/category/update'
        Axios.post(updateURL,{category_id:category.category_id,canteen_id:this.$store.state.canteen_id,category_name:category.category_name})
        .then((response) => {
          // 处理响应数据
          if(response.data.status === 200){
          // 处理响应数据
          this.$message.success('修改成功!')
          this.getCategory(category.category_id ? this.page : 1)
          this.category = {
            category_id:null,
            category_name:''
          }
          this.addDialogVisible = false
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
        const addURL = '/api1/v2/admin/category/add'
        Axios.post(addURL,{canteen_id:this.$store.state.canteen_id,category_name:category.category_name})
        .then((response) => {
          if(response.data.status === 200){
          // 处理响应数据
          this.$message.success('添加成功!')
          this.getCategory(category.category_id ? this.page : 1)
          this.category = {
            category_id:null,
            category_name:''
          }
          this.addDialogVisible = false
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            this.notifyError('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
        }
        ).catch(error => {
          // 处理请求错误
          this.$message.error('添加失败!')
        })       
      }
    },
    // 修改
    showEditDialog(row){
      //显示对话框
      this.addDialogVisible = true;
      //浅拷贝
      this.category = {
        ...row
      }
    },
    // 显示添加dialog
    showDialog(){
      this.addDialogVisible = true;
      //清空
      this.category = {
        category_id:null,
        category_name:''
      }
    },
    // 获取页面种类信息
    getCategory(page=1) {
      this.page = page;
      const url = '/api1/v2/admin/category/list'
      Axios.post(url, {canteen_id:this.$store.state.canteen_id})
        .then(response => {
          // 处理响应数据
          if(response.data.status === 200){
            this.categoryList = response.data.data.category_list
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            this.notifyError('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
          })
          .catch(error => {
          // 处理请求错误
          this.notifyError('获取菜品种类失败！')
          })
    },
    // 每页显示数量被改变
    handleSizeChange(size){
      this.limit = size;
      this.getCategory();
    },
  },
};
</script>

<style>
.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
</style>
