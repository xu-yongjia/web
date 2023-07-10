<!--
 * @Descripttion:商品评价页面组件 
-->
<template>
  <div class="judgemain">
    <div>
      <div class="title">
        <p>商品评价</p>
      </div>
        <div class="product">
            <div class="pro-img">
                <img :src="productDetails.img_path" class="pro-img"/>
            </div>
            <div class="prodetail">
                <div class="content">
                 
                    <p class="name">{{productDetails.name}}</p>
                    <p class="intro">商品简介：{{productDetails.info}}</p>
        
                    <div class="price">
                      <span>价格：{{productDetails.price}}元</span>
                      <span
                        v-show="productDetails.product_price != productDetails.product_selling_price"
                        class="del"
                      >{{productDetails.discount_price}}元</span>
                    </div>

                    <div class="canteen">
                      <span>所属食堂：{{placemap[productDetails.canteen_id]}}</span>                      
                    </div>

                    <div class="category">
                      <span>所属分类：{{productDetails.category_name}}</span>                     
                    </div>
                </div>              
            </div>         
        </div>
      </div>      

       <div class="judgeinput">
            <div style="letter-spacing: 1px">
              <div class="scorehead">
                  <p class="caipin">
                    请为菜品打分：
                  </p>
                  <el-radio-group class="el-radio-group"  v-model="Sc" >
              <el-radio class="el-radio__label"  label="0" >0分</el-radio>
              <el-radio label="1" >1分</el-radio>
              <el-radio label="2" >2分</el-radio>
              <el-radio label="3" >3分</el-radio>
              <el-radio label="4" >4分</el-radio>
              <el-radio label="5" >5分</el-radio>
               </el-radio-group>

              </div>
                <div class="head">
                  <p class="caipin">
                    请为菜品作出评价：
                  </p>
                    <input
                      type="text"
                      placeholder="请输入评价 . . ."
                      ref="input2"            
                      v-model="Comment"          
                    />
                    <div class="submitbut">
                        <button @click="submitcomment">发布评价</button>
                    </div>
                </div>
            </div>
            
        </div>
        
    </div>
</template>

<script>
export default {

  components: {
  //  StarRating,
    },

    data(){
        return{
          
            user_id:'',
            scoreInput:'',
            imageUrl: '',
            Comment:"",
            productDetails:"",
            starLevel: '',
            starIndex: 0,
            text: '',
            Sc:'',
            form: {
       // id: 0,
       // nickname: '',
        avatar: '',
       // user_name: ''
      },

      placemap:["无","榕园","岁月湖","荔园","若海","瑾园"]
        };
    },

   
    created() {
        this.getDetailsPicture(this.$route.query.productID);
        this.user_id=this.$store.getters.getUser.user_id;
    },
    computed: {
    score: {
      get() {
        // 将输入框的值转换为数字并返回
        return parseFloat(this.scoreInput);
      },
      set(value) {
        // 将数字值转换为字符串并赋给输入框
        this.scoreInput = value.toString();
      }
    }
  },

  

   methods:{
    //提交评价
    submitcomment(){
      this.$axios
        .post("/api/users/submitComment", {
        user_id:this.user_id,
        product_id: parseInt(this.$route.query.productID),
        order_id:parseInt(this.$route.query.orderID),
        score: this.Sc,
        product_comment: this.Comment,  
        })
        .then(res => {
          console.log(res)  
        })
        .catch(err => {
          return Promise.reject(err);
        });
        this.score=''
        this.Comment=''
    },
  
    onStarClick(index) {
        this.starIndex = index;
      },
      submit() {
        const formData = new FormData();
        formData.append('stars', this.starIndex);
        formData.append('text', this.text);
        // 将评级和文本信息打包成 FormData 发送给服务器
        // TODO: 发送数据到服务器
      },


    getDetailsPicture(val) {
      this.$axios
        .post("/api/product/getDetails", {
          productID: val
        })
        .then(res => {
            console.log(res)  
            this.productDetails = res.data.data;
           
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },

    change(index) {
      this.starIndex = index;
      switch (this.starIndex) {
        case 1:
          this.starLevel = '很差';
          break;
        case 2:
          this.starLevel = '较差';
          break;
        case 3:
          this.starLevel = '一般';
          break;
        case 4:
          this.starLevel = '很好';
          break;
        case 5:
          this.starLevel = '非常好';
          break;
        default:
          this.starLevel = '';
          break;
      }
    },

    validateScore() {
      // 在输入框的 @input 事件中触发该方法进行验证
      // 使用正则表达式检查输入是否为数字，并限制在 0 到 5 之间
      const isValid = /^([0-5](\.\d*)?)?$/.test(this.scoreInput);      
      if (!isValid) {
        // 如果输入不符合规则，则清空输入框的值
        this.scoreInput = "";
      }
    }
    }
}
</script>

<style>

.caipin{
  font-family: STXinwei;
  font-size: 35px;
  padding-left: 79px;
  padding-bottom: 5px;
}

.el-radio-group {
    font-size: 0;
    padding-left: 71px;
    padding-top: 1px;
    padding-bottom: 15px;
}

.el-radio__label {
    padding-left: 10px;
    font-size: 24px;
    padding-top: 10px;
}

.judgeinput{
  margin-top: 20px;
}

.judgemain {
  background-color: #ffffffd1;
  padding-left: 81px;
}

.title{
  font-size: 35px;
    font-family: STXinwei;
    padding-bottom: 25px;
    padding-top: 25px;
    color: #215e07;
}
/* 图片 */

.avatar-uploader {
  width: 1085px;
  height: 80px;
  display: flex;
  justify-content: flex-start;
  margin-right: 0; 
}

.avatar {
  width: 100px; 
}


.pro-img{
    width: 462px;
    height: 462px;
    float: left;
    border: 1px solid #e7e7e7;
}
.prodetail{
    width: 526px;
    height: 462px;
    margin-left: 463px;
    border-top: 1px solid #e7e7e7;
    padding-left: 68px;
}
.pro-name{
  float: left;
  width: 160px;
  padding-right: 18px;
  margin:50px;
  text-align: center;
}
.pro-price{
  float: left;
  width: 160px;
  padding-right: 18px;
  text-align: center;
}
.content {
    height: 462.89px;
}

.content .name {
  font-size: 25px;
  color: black;
  padding-top: 20px;
  padding-left: 15px;
  font-weight: bold;
}

.content .intro {
  font-family: "楷体", KaiTi, serif;
  font-size: 25px;
  color: #b0b0b0;
  padding-top: 28px;
  padding-bottom: 15px;
  padding-left: 15px;
}
.content .store {
  color: #ff6700;
  padding-top: 10px;
}
.content .price {
  display: block;
  font-size: 24px;
  color: #ff6700;
  border-bottom: 1px solid #e0e0e0;
  padding: 25px 15px 25px;
}

.content .canteen {
  display: block;
  font-size: 24px;
  color: #ff6700;
  border-bottom: 1px solid #e0e0e0;
  padding: 25px 15px 25px;
}

.content .category {
  padding-left: 15px;
  display: block;
  font-size: 24px;
  color: #ff6700;
  border-bottom: 1px solid #e0e0e0;
  padding: 25px 15px 25px;
}


.content .price .del {
  font-size: 14px;
  margin-left: 10px;
  color: #b0b0b0;
  text-decoration: line-through;
}
.content .pro-list {
  background: #f9f9fa;
  padding: 30px 60px;
  margin: 50px 0 50px;
}
.content .pro-list span {
  line-height: 30px;
  color: #616161;
}
.content .pro-list .pro-price {
  float: right;
}
.content .pro-list .pro-price .pro-del {
  margin-left: 10px;
  text-decoration: line-through;
}
.content .pro-list .price-sum {
  color: #ff6700;
  font-size: 24px;
  padding-top: 20px;
}
.content .button {
  height: 55px;
  margin: 10px 0 20px 0;
}
.content .button .el-button {
  float: left;
  height: 55px;
  font-size: 16px;
  color: #fff;
  border: none;
  text-align: center;
}
.content .button .shop-cart {
  width: 340px;
  background-color: #ff6700;
}
.content .button .shop-cart:hover {
  background-color: #f25807;
}

.content .button .like {
  width: 260px;
  margin-left: 40px;
  background-color: #b0b0b0;
}
.content .button .like:hover {
  background-color: #757575;
}
.content .pro-policy li {
  float: left;
  margin-right: 20px;
  color: #b0b0b0;
}


/* 评论区 */
.scorehead {
  background-color: rgb(248, 248, 248);
  position: relative;
  height: 75px;
  border-radius: 5px;
  padding-bottom: 15px;
  padding-top: 9px;
}

.head {
  background-color: rgb(248, 248, 248);
  position: relative;
  height: 290px;
  border-radius: 5px;
}

.head img {
  width: 55px;
  height: 55px;
  border-radius: 50%;
  position: absolute;
  top: 10px;
  left: 13px;
}
/* 评论框 */
.scorehead input {
  position: absolute;
  top: 13px;
  left: 80px;
  height: 45px;
  border-radius: 5px;
  outline: none;
  width: 90%;
  font-size: 20px;
  padding: 0 20px;
  border: 2px solid #f8f8f8;
}

.head input {
  position: absolute;
    top: 44px;
    left: 80px;
    height: 172px;
    border-radius: 5px;
    outline: none;
    width: 90%;
    font-size: 20px;
    padding: 0px 18px;
    border: 2px solid #f8f8f8;
}

/* 发布评论按钮 */
.submitbut button {
  position: absolute;
  top: 229px;
  left: 82px;
  width: 120px;
  height: 48px;
  border: 0;
  border-radius: 5px;
  font-size: 20px;
  font-weight: 500;
  color: #fff;
  background-color: rgb(230, 159, 17);
  cursor: pointer;
  letter-spacing: 2px;
}


.custom-file-upload {
  display: inline-block;
  padding: 6px 12px;
  cursor: pointer;
  border: 1px solid #ccc;
  border-radius: 4px;
  background-color: #fff;
}
</style>