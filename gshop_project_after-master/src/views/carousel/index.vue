<template>
  <div >
    <div style="text-align:center;">
      <p>新品详情</p>
        <textarea v-model="detail" name="" id="" cols="30" rows="10" style="resize: none;width: 1000px; height: 100px;"></textarea>
      <p>点击下方上传新品推荐的图片</p>
    </div>
    <el-upload
    style="text-align:center;"
        class="avatar-uploader"
        action
        label="描述"
        ref="upload"
        :show-file-list="false"
        :before-upload="beforeAvatarUpload"
        :http-request="uploadRequest">
        <img v-if="picUrl" :src="picUrl" class="avatar">
        <i v-else class="el-icon-plus avatar-uploader-icon"></i>
      <div class="el-upload__tip" slot="tip">点击头像上传，只能上传jpg/png文件，且不超过2M</div>
    </el-upload>
    <el-button type="primary" @click="confirm" style="margin-left: 550px;">确 定</el-button>
  </div>
</template>

<script>
import { Message } from 'element-ui';
import Axios from 'axios';
export default {
  name: 'Carousel',
  data(){
    return{
      selectedFile: null,
      //logo真实地址
      picUrl:'',//前端显示
      avatar:'',//发给后端
      detail:'',
    }
  },
  methods: {
    // 上传图片之前格式/大小检验
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg' || file.type === 'image/png'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isJPG) {
        Message.error('上传头像图片只能是 JPG/PNG 格式!')
      }
      if (!isLt2M) {
        Message.error('上传头像图片大小不能超过 2MB!')
      }
      return isJPG && isLt2M;
    },

    uploadRequest(option){
      const url = '/api1/v2/admin/generateurl'
      Axios.post(url,{filename:option.file.name})
      .then(response =>{
        if(response.data.status === 200){
            const oReq = new XMLHttpRequest()
            //前端把图片放put
            oReq.open('PUT',response.data.data.put,true)
            oReq.send(option.file)
            oReq.onload = () =>{
              this.picUrl = response.data.data.get
              this.avatar = response.data.data.get
              // this.avatar = response.data.data.key
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
    confirm(){
      if(this.avatar){
      const url = '/api1/v2/admin/saveCarouselUrl'
      Axios.post(url, {canteen_id:this.$store.state.canteen_id,carousel_url:this.avatar,board:this.detail})
        .then(response => {
          // 处理响应数据
          if(response.data.status === 200){
            Message.success("上传成功！")
            this.url=''
            this.avatar=''
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
          Message.error('上传图片失败！')
        }) 
    }  
    else{
      Message.error('图片为空！')
    }
    },
  },
}

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
    width: 1225px;
    height: 512px;
    display: block;
  }
</style>