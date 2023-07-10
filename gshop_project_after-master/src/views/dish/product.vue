<template>
  <div>
    <!-- 添加按钮 -->
    <br>
    <br>
    <el-button type="" icon="el-icon-thumb" @click="back" >返回</el-button>
    <el-button type="primary" icon="el-icon-plus" @click="showDialog" style="margin-left: 10px;">添加</el-button>
    <el-table 
      :data="productList"
      border style="width: 100%; margin: 20px 0">
      <el-table-column
        label="序号"
        align="center"
        type="index"
        width="80"
      >
      </el-table-column>
      
      <el-table-column prop="name" label="菜品名称" width="100" />
      <el-table-column prop="category_name" label="菜品种类" width="100" />
      <el-table-column prop="Info" label="详情" width="width" />
      <el-table-column prop="img_path" label="图片" width="150">
        <template slot-scope="{row}">
            <img :src="row.img_path" style="width:80px;height:60px"/>
        </template>
      </el-table-column>
      <el-table-column prop="price" label="原价" width="100" />
      <el-table-column prop="discount_price" label="现价" width="100" />
      <el-table-column prop="score" label="评分" width="100" />
      <el-table-column prop="title" label="其他" width="100" />


      <el-table-column prop="prop" label="操作" width="width">
          <!-- 由于要为每一个数据绑定,所以要循环,这样子才能传递数据 -->
          <template slot-scope="{row}">
            <!-- 传递当前循环的一个数据给函数 -->
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="showEditDialog(row)">编辑</el-button>
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="removeDialog(row)">删除</el-button>
            <el-button type="primary" icon="el-icon-view" size="mini" @click="showUploadDialog(row)">上传图片</el-button>
          </template>
      </el-table-column>
    </el-table>

    <el-pagination
      @size-change="handleSizeChange"
      @current-change="getProduct"
      style="text-align: center"
      :current-page="page"
      :page-sizes="[3, 6, 12]"
      :page-size="3"
      layout=" prev, pager, next, jumper,->,sizes,total"
      :total="total"
    >
    </el-pagination>

    <!-- 添加不需要id的,而修改是需要id的 -->
    <el-dialog :title="product.id?'修改菜品':'添加菜品'" :visible.sync="addDialogVisible">
      <!-- el-form通过:model="form" form为一个对象,用于将表单所有数据进行收集 -->
      <el-form :model="product" :rules="rules" style="width:80%" ref="product">
        <el-form-item v-if="product.id" label="菜品编号" prop="id" >
          <el-input v-model="product.id" :readonly="true"/>
        </el-form-item>
        <!-- 菜品名称 -->
        <el-form-item label="菜品名称" label-width="100px" prop="name">
          <el-input v-model="product.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="信息" label-width="100px" prop="Info">
          <el-input v-model="product.Info" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="原价" label-width="100px" prop="price">
          <el-input v-model="product.price" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="现价" label-width="100px" prop="discount_price">
          <el-input v-model="product.discount_price" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="其他" label-width="100px" prop="title">
          <el-input v-model="product.title" autocomplete="on"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false ">取 消</el-button>
        <el-button type="primary" @click="confirmAddOrEdit">确 定</el-button>
      </div>
    </el-dialog>

    <!-- 上传菜品图片dialog -->
    <el-dialog title="上传菜品图片" :visible.sync="uploadDialogVisible">
      <!-- el-form通过:model="form" form为一个对象,用于将表单所有数据进行收集 -->
      <el-form :model="product" :rules="upload_rules" style="width:80%" ref="product">
        <el-form-item label="菜品编号" label-width="100px" prop="id"  >
          <el-input v-model="product.id" :readonly="true"/>
        </el-form-item>
        <el-form-item label="菜品名称" label-width="100px" prop="name" >
          <el-input v-model="product.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="封面图：" label-width="100px" prop="" >
          <template>
          <div>
            <el-upload
              class="avatar-uploader"
              action
              label="描述"
              ref="upload"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :http-request="uploadSmallReq">
              <img v-if="product.img_path" :src="product.img_path" class="avatar">
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                <div class="el-upload__tip" slot="tip">点击头像上传，只能上传jpg/png文件，且不超过2M</div>
            </el-upload>
          </div>
        </template> 
        </el-form-item>

        <el-form-item label="详情图：" label-width="100px" prop="" >
          <template >
          <div>
            <el-upload 
              class="avatar-uploader"
              action
              label="描述"
              ref="upload"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :http-request="uploadBigReq1">
              <img v-if="big_urls[0]" :src="big_urls[0]" class="avatar" >
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                <div class="el-upload__tip" slot="tip">点击头像上传，只能上传jpg/png文件，且不超过2M</div>
            </el-upload>
            <el-upload 
              class="avatar-uploader"
              action
              label="描述"
              ref="upload"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :http-request="uploadBigReq2">
              <img v-if="big_urls[1]" :src="big_urls[1]" class="avatar" >
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                <div class="el-upload__tip" slot="tip">点击头像上传，只能上传jpg/png文件，且不超过2M</div>
            </el-upload>
            <!-- v-for=" (i) in 3" :key="i" -->
            <el-upload 
              class="avatar-uploader"
              action
              label="描述"
              ref="upload"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :http-request="uploadBigReq3">
              <img v-if="big_urls[2]" :src="big_urls[2]" class="avatar" >
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                <div class="el-upload__tip" slot="tip">点击头像上传，只能上传jpg/png文件，且不超过2M</div>
            </el-upload>
          </div>
        </template> 
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelUpload">取 消</el-button>
        <el-button type="primary" @click="confirmUpload">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Axios from 'axios';
import { Message } from 'element-ui';
export default {
  name: "Product",
  data() {
    return {
      //当前页
      page: 1,
      //每页显示数据
      limit: 3,
      //数据总数
      total: 0,
      //数据列表
      productList: [],
      //控制对话框的显示隐藏
      addDialogVisible:false,
      uploadDialogVisible:false,
      category_id:1,
      //收集表单数据
      product:{
        //菜品名称
        id:null,
        name:'',
        Info:'',
        img_path:'',
        price:'',
        discount_price:'',
        title:'无',
        score:'',
      },
      //imp_path
      avatar:'',//发给后端
      big_urls:[],//前端显示
      big_avatars:[],//发给后端
      //验证规则(由对象组成)
      rules:{
        // 菜品名称校验
        name:[
          {required:true,message:"请填写菜品!",trigger:'change'}
        ],
        Info:[
          {required:true,message:"请填写菜品详情!",trigger:'change'}
        ],
        price:[
          {required:true,message:"请填写价格!",trigger:'change'}
        ],
        discount_price:[
          {required:true,message:"请填写折后价!",trigger:'change'}
        ],
      },
      upload_rules:{ }
    };
  },
  mounted() {
    this.category_id = this.$route.query.category_id
    this.getProduct();  
  },
  methods: {

    // 上传图片之前格式/大小检验
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg' || file.type === 'image/png'
      if (!isJPG) {
        Message.error('上传图片只能是 JPG/PNG 格式!')
      }
      return isJPG ;
    },
    uploadSmallReq(option){
      const url = '/api1/v2/admin/generateurl'
      Axios.post(url,{filename:option.file.name})
      .then(response =>{
        if(response.data.status === 200){
            const oReq = new XMLHttpRequest()
            //前端把图片放put
            oReq.open('PUT',response.data.data.put,true)
            oReq.send(option.file)
            oReq.onload = () =>{
              this.product.img_path = response.data.data.get
              this.avatar = response.data.data.get
            }
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            Message.error('身份过期，请重新登录！')
            this.$store.dispatch(logout)
          }        
      })
      .catch(error =>{
        Message.error('上传失败！')
      })
    },
    uploadBigReq1(option){
      const url = '/api1/v2/admin/generateurl'
      Axios.post(url,{filename:option.file.name})
      .then(response =>{
        if(response.data.status === 200){
            const oReq = new XMLHttpRequest()
            //前端把图片放put
            oReq.open('PUT',response.data.data.put,true)
            oReq.send(option.file)
            oReq.onload = () =>{
              this.big_urls.splice(0,1,response.data.data.get)
              this.big_avatars.splice(0,1,response.data.data.get)
            }
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            Message.error('身份过期，请重新登录！')
            this.$store.dispatch(logout)
          }        
      })
      .catch(error =>{
        Message.error('上传失败！',error)
      })
    },
    uploadBigReq2(option){
      const url = '/api1/v2/admin/generateurl'
      Axios.post(url,{filename:option.file.name})
      .then(response =>{
        if(response.data.status === 200){
            const oReq = new XMLHttpRequest()
            //前端把图片放put
            oReq.open('PUT',response.data.data.put,true)
            oReq.send(option.file)
            oReq.onload = () =>{
              this.big_urls.splice(1,1,response.data.data.get)
              this.big_avatars.splice(1,1,response.data.data.get)
            }
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            Message.error('身份过期，请重新登录！')
            this.$store.dispatch(logout)
          }        
      })
      .catch(error =>{
        Message.error('上传失败！',error)
      })
    }, 
    uploadBigReq3(option){
      const url = '/api1/v2/admin/generateurl'
      Axios.post(url,{filename:option.file.name})
      .then(response =>{
        if(response.data.status === 200){
            const oReq = new XMLHttpRequest()
            //前端把图片放put
            oReq.open('PUT',response.data.data.put,true)
            oReq.send(option.file)
            oReq.onload = () =>{
              this.big_urls.splice(2,1,response.data.data.get)
              this.big_avatars.splice(2,1,response.data.data.get)
            }
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            Message.error('身份过期，请重新登录！')
            this.$store.dispatch(logout)
          }        
      })
      .catch(error =>{
        Message.error('上传失败！',error)
      })
    },     
    confirmUpload(){
      const url = '/api1/v2/admin/product/photo/add'
      Axios.post(url, {
        product_id:this.product.id,
        canteen_id:this.$store.state.canteen_id,
        big_urls:this.big_avatars,
        small_url:this.avatar,
      })
        .then((response) => {
          // 处理响应数据
          if(response.data.status === 200){
            Message.success("上传成功！")
            this.uploadDialogVisible = false
            this.product = {
              id:null,
              name:'',
              Info:'',
              img_path:'',
              price:'',
              discount_price:'',
              title:'无',
              score:'',
            }
            this.avatar=''
            this.big_urls = []
            this.big_avatars = []
            this.getProduct()
          }
          else if(response.data.status === 201){
            Message.error(response.data.msg)
          }
          else if(response.data.status === 404){
            this.$store.dispatch(logout)
          }
        })
        .catch(error => {
          // 处理请求错误
          this.notifyError('上传图片失败！')
        })   
    },
    cancelUpload(){
      this.uploadDialogVisible = false
        this.product = {
          id:null,
          name:'',
          Info:'',
          img_path:'',
          price:'',
          discount_price:'',
          title:'无',
          score:'',
        }
        this.avatar=''
        this.big_urls = []
        this.big_avatars = []
    },
    // 用户单击删除
    removeDialog(Info){
      this.$confirm(`确认删除${Info.name}吗?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const url = '/api1/v2/admin/product/del'
        Axios.post(url, {product_id:Info.id})
        .then(response => {
          // 处理响应数据
          if(response.data.status === 200){
            Message.success("删除菜品成功！")
            this.getProduct()
          }
          else if(response.data.status === 201){
              Message.error('有未完成的该商品订单，删除失败')
          }
          else if(response.data.status === 404){
            //token过期
            this.notifyError('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
          })
          .catch(error => {
          // 处理请求错误
          this.notifyError('删除失败！'+ error)
          })
      }).catch(() => {
        this.$message({
          type: 'Info',
          message: '已取消删除'
        });          
      });
    },
    // 用户单击确认添加/修改
    confirmAddOrEdit(){
      const {product} = this
      //id存在，则是更新
      if(product.id){
        const updateURL = '/api1/v2/admin/product/update'
        Axios.post(updateURL,{
          id:product.id,
          name:product.name,
          canteen_id:this.$store.state.canteen_id,
          category_id:this.category_id,
          Info:product.Info,
          price:product.price,
          discount_price:product.discount_price,
          title:product.title
        })
        .then((response) => {
          // 处理响应数据
          if(response.data.status === 200){
          // 处理响应数据
          Message.success('修改成功!')
          this.getProduct(product.id ? this.page : 1)
          this.product = {
            id:null,
            name:'',
            Info:'',
            img_path:'',
            price:'',
            discount_price:'',
            title:'无',
            score:'',
          }
          this.addDialogVisible = false
          }
          else if(response.data.status === 201){
            Message.error( response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            this.notifyError('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
        }
        ).catch(error => {
          // 处理请求错误
          Message.error('修改失败!')
        })
      }
      else{
        const addURL = '/api1/v2/admin/product/add'
        Axios.post(addURL,{
          name:product.name,
          canteen_id:this.$store.state.canteen_id,
          category_id:this.category_id,
          Info:product.Info,
          price:product.price,
          discount_price:product.discount_price,
          title:product.title
        })
        .then((response) => {
          if(response.data.status === 200){
          // 处理响应数据
          Message.success('添加成功!')
          this.getProduct(product.id ? this.page : 1)
          this.product = {
            id:null,
            name:'',
            Info:'',
            img_path:'',
            price:'',
            discount_price:'',
            title:'无',
            score:'',
          }
          this.addDialogVisible = false
          }
          else if(response.data.status === 201){
            Message.error( response.data.msg)
          }
          else if(response.data.status === 404){
            //token过期
            this.notifyError('登录已过期，请重新登录！')
            this.$store.dispatch(logout)
          }
        }
        ).catch(error => {
          // 处理请求错误
          Message.error('添加失败!')
        })       
      }
    },
    // 修改
    showEditDialog(row){
      //显示对话框
      this.addDialogVisible = true
      //浅拷贝
      this.product = {
        ...row
      }
    },
    showUploadDialog(row){
      this.uploadDialogVisible = true
      this.product = {
        ...row
      }
      if(this.product.img_path)
        this.avatar=this.product.img_path
    },
    // 显示添加dialog
    showDialog(){
      this.addDialogVisible = true;
      //清空
      this.product = {
        id:null,
        name:'',
        Info:'',
        img_path:'',
        price:'',
        discount_price:'',
        title:'无',
        score:'',
      }
    },
    // 获取页面菜品信息
    getProduct(page=1) {
      this.page = page;
      const url = '/api1/v2/admin/product/list'
      Axios.post(url, {page:this.page,limit:this.limit,canteen_id:this.$store.state.canteen_id,category_id:this.category_id})
        .then((response) => {
          // 处理响应数据
          if(response.data.status === 200){
            this.productList = response.data.data.productlist
            this.total = response.data.data.count
            
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
          Message.error('获取菜品失败！')
          })
    },
    // 每页显示数量被改变
    handleSizeChange(size){
      this.limit = size;
      this.getProduct();
    },
    back(){
      this.$router.push({name:'Category'})
    }
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
