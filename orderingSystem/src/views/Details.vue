<!--
 * @Description: 商品详情页面组件
 -->
<template>
  <div id="details">
    <!-- 导航栏 -->
  <div class="sidebar">
    <ul>
    <li><a href="#details" >点菜</a></li>
    <li><a href="#Anchor2" >评价</a></li>
    </ul>
</div>
    <!-- 主要内容 -->   
    <div class="main">
    <div class="gaishu">
      <div class="total">
        <span>商品概述</span>
      </div>
      <!-- 左侧商品轮播图 -->
      <div class="block">
        <el-carousel height="560px" v-if="productPicture.length>=1">
          <el-carousel-item v-for="item in productPicture" :key="item.id">
            <img style="height:560px;" :src="item.img_path" :alt="2" />
          </el-carousel-item>
        </el-carousel>
      </div>
      <!-- 左侧商品轮播图END -->

      <!-- 右侧内容区 -->
      <div class="content">
        <h1 class="name">{{productDetails.name}}</h1>
        <p class="intro">{{productDetails.info}}</p>
        <div class="price">
          <span>{{productDetails.discount_price}}元</span>
          <span
            v-show="productDetails.product_price != productDetails.product_selling_price"
            class="del"
          >{{productDetails.price}}元</span>
        </div>
        <div class="canteen">
            <span>菜品来自：{{placemap[productDetails.canteen_id]}}</span>
            
        </div>
        <div class="category">
            <span>菜品类型：{{productDetails.category_name}}</span>
                      
        </div>
        <div class="pro-list">
          <span class="pro-name">{{productDetails.name}}</span>
          <span class="pro-price">
            <span>{{productDetails.discount_price}}元</span>
            <span
              v-show="productDetails.discount_price != productDetails.price"
              class="pro-del"
            >{{productDetails.price}}元</span>
          </span>
          <p class="price-sum">总计 : {{productDetails.discount_price}}元</p>
        </div>
        <!-- 内容区底部按钮 -->
        <div class="button">
          <el-button class="shop-cart" :disabled="dis" @click="addShoppingCart">加入购物车</el-button>
          <el-button class="like" @click="addCollect">喜欢</el-button>
        </div>
      </div>
      <!-- 右侧内容区END -->
    </div>
  </div>
    <!-- 评论区 -->
    <div id="Anchor2" class="box-bd">
        <div class="comment">
          <Comment :commentList="commentList"></Comment>
        </div>
    </div>
    <!-- 主要内容END -->
  </div>
</template>
<script>
import { mapActions } from "vuex";
import Comment from "../../src/components/comment/Comment.vue"
export default {
  components:{
    Comment
  },
  data() {
    return {
      commentList:[
            ], 
      dis: false, // 控制“加入购物车按钮是否可用”
      productID: "", // 商品id
      productDetails: "", // 商品详细信息
      productPicture: "", // 商品图片
      placemap:["无","榕园","岁月湖","荔园","若海","瑾园"]
    };
  },
  // 通过路由获取商品id
  activated() {
    if (this.$route.query.productID != undefined) {
      this.productID = this.$route.query.productID;
    }
  },
  watch: {
    // 监听商品id的变化，请求后端获取商品数据
    productID: function(val) {
      this.getDetails(val);
      this.getDetailsPicture(val);
      this.getDetailsComment(val)
    }
  },
  methods: {
    ...mapActions(["unshiftShoppingCart", "addShoppingCartNum"]),
    // 获取商品详细信息
    getDetails(val) {
      this.$axioss
        .post("/api/product/getDetails", {
          productID: val
        })
        .then(res => {
          this.productDetails = res.data.data;
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    // 获取商品图片
    getDetailsPicture(val) {
      this.$axios
        .post("/api/product/getDetailsPicture", {
          productID: parseInt(val)
        })
        .then(res => {
          this.productPicture = res.data.data;
       
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    //获取商品评价
    getDetailsComment(val) {
      this.$axios
        .post("/api/product/getDetailsComment", {
          productID: parseInt(val)
        })
        .then(res => {
          this.commentList = res.data.data.items;
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },

    // 加入购物车
    addShoppingCart() {
      // 判断是否登录,没有登录则显示登录组件
      if (!this.$store.getters.getUser) {
        this.$store.dispatch("setShowLogin", true);
        return;
      }
      this.$axios
        .post("/api/users/shoppingCart/addShoppingCart", {
          user_id: this.$store.getters.getUser.user_id,
          product_id: this.productID
        })
        .then(res => {
        
          switch (res.status) {
            case 200:
              // 新加入购物车成功           
              this.notifySucceed(res.data.msg);
              this.unshiftShoppingCart(res.data.data.shoppingCartData[0]);
              location.reload(true);
              break;
            case 201:
              // 该商品已经在购物车，数量+1
              this.addShoppingCartNum(this.productID);
              this.notifySucceed(res.data.msg);
              location.reload(true);
              break;
           
            default:
              this.notifyError(res.data.msg);
          }
        })
        .catch(err => {
          return Promise.reject(err);
        });
    },
    addCollect() {
      // 判断是否登录,没有登录则显示登录组件
      if (!this.$store.getters.getUser) {
        this.$store.dispatch("setShowLogin", true);
        return;
      }
      this.$axios
        .post("/api/users/collect/addCollect", {
          user_id: this.$store.getters.getUser.user_id,
          product_id: parseInt(this.productID)
          
        })
        .then(res => {
          if (res.data.status == 200) {
            // 添加收藏成功
            this.notifySucceed(res.data.msg);
          } else {
            // 添加收藏失败
            this.notifyError(res.data.msg);
          }
        })
        .catch(err => {
          return Promise.reject(err);
        });
    }
  }
};
</script>
<style>

h1 {
  font-size: 32px;  /* 字体大小 */
  font-weight: bold;  /* 粗体字 */
  color: #dc9222;  /* 颜色 */
  margin-top: 20px;  /* 上外边距 */
  margin-bottom: 25px;  /* 下外边距 */
}

/* 头部CSS */
#details .page-header {
  height: 64px;
  margin-top: -20px;
  z-index: 4;
  background: #fff;
  border-bottom: 1px solid #e0e0e0;
  -webkit-box-shadow: 0px 5px 5px rgba(0, 0, 0, 0.07);
  box-shadow: 0px 5px 5px rgba(0, 0, 0, 0.07);
}
#details .page-header .title {
  width: 1225px;
  height: 64px;
  line-height: 64px;
  font-size: 18px;
  font-weight: 400;
  color: #212121;
  margin: 0 auto;
}
#details .page-header .title p {
  float: left;
}
#details .page-header .title .list {
  height: 64px;
  float: right;
}
#details .page-header .title .list li {
  float: left;
  margin-left: 20px;
}
#details .page-header .title .list li a {
  font-size: 14px;
  color: #616161;
}
#details .page-header .title .list li a:hover {
  font-size: 14px;
  color: #ff6700;
}
/* 头部CSS END */

/* 主要内容CSS */

.main{
  background-color: #f6f5f5e0;
}

.main .total span{
  color:#0e510b;
  font-size: 60px;
  font-family: STXinwei;
  padding-left: 15px;
}

.name{
  font-size: 30px;
  font-family: "楷体", KaiTi, serif;
}

.intro{
  font-family: "楷体", KaiTi, serif;
  font-size: 25px;
  color: #5f5e5e;
  padding-top: 28px;
  padding-bottom: 15px;
  padding-left: 0px;
}
.price{
  display: block;
  font-size: 37px;
  color: #ff6700;
  border-bottom: 1px solid #e0e0e0;
  padding: 0px 15px 25px;
}
.canteen{
  display: block;
  font-size: 20px;
  color: #000000;
  border-bottom: 1px solid #e0e0e0;
  padding: 0px 0px 25px;
}
.category{
  display: block;
  font-size: 20px;
  color: #000000;
  border-bottom: 1px solid #e0e0e0;
  padding: 0px 0px 25px;
}


#details .main {
  width: 1310px;
  height: 650px;
  padding-top: 30px;
  margin: 0 auto;
}
#details .main .block {
  float: left;
  width: 560px;
  height: 560px;
}
#details .el-carousel .el-carousel__indicator .el-carousel__button {
  background-color: rgba(163, 163, 163, 0.8);
}
#details .main .content {
  float: left;
  margin-left: 25px;
  width: 640px;
}
#details .main .content .name {
  height: 30px;
  line-height: 30px;
  font-size: 30px;
  font-weight: normal;
  color: #212121;
}
#details .main .content .intro {
  color: #595656;
  padding-top: 0px;
}
#details .main .content .store {
  color: #ff6700;
  padding-top: 10px;
}
#details .main .content .price {
  display: block;
  font-size: 37px;
  color: #ff6700;
  border-bottom: 1px solid #e0e0e0;
  padding: 0px 0 25px;
}
#details .main .content .price .del {
  font-size: 14px;
  margin-left: 10px;
  color: #b0b0b0;
  text-decoration: line-through;
}
#details .main .content .pro-list {
  background: #f9f9fa;
  padding: 30px 60px;
  margin:21px 0 50px;
}
#details .main .content .pro-list span {
  line-height: 30px;
  color: #616161;
}
#details .main .content .pro-list .pro-price {
  float: right;
}
#details .main .content .pro-list .pro-price .pro-del {
  margin-left: 10px;
  text-decoration: line-through;
}
#details .main .content .pro-list .price-sum {
  color: #ff6700;
  font-size: 24px;
  padding-top: 20px;
}
#details .main .content .button {
  height: 55px;
  margin: 10px 0 20px 0;
}
#details .main .content .button .el-button {
  float: left;
  height: 55px;
  font-size: 16px;
  color: #fff;
  border: none;
  text-align: center;
}
#details .main .content .button .shop-cart {
  width: 340px;
  background-color: #ff9900;
}
#details .main .content .button .shop-cart:hover {
  background-color: #f25807;
}

#details .main .content .button .like {
  width: 260px;
  margin-left: 40px;
  background-color: #b0b0b0;
}
#details .main .content .button .like:hover {
  background-color: #757575;
}
#details .main .content .pro-policy li {
  float: left;
  margin-right: 20px;
  color: #b0b0b0;
}
/* 主要内容CSS END */
#details .box-bd {
  background-color: #f6f5f5e0;
  width: 1310px;
  height: 560px;
  padding-top: 30px;
  margin: 0 auto;
}

.sidebar {
    position: fixed;
    top: 50%;
    right: 0;
    transform: translateY(-50%);
    background-color: #136331e0;
    padding: 13px;
    border-radius: 20px;
}

.sidebar ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.sidebar li {
  font-size: 25px;
  color: #ffffff;
  margin-bottom: 10px;
}

.sidebar a {
    text-decoration: none;
    font-family: STXinwei;
    color: #fff;
}

</style>