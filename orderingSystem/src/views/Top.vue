<!--
 * @Descripttion:榜单页面组件 
-->
<template>
<div class="pcontent">
  <div id="container">
  <div class="content">
  <div class="title">
    <h4>本周好评榜</h4>
  </div>  
    <ul>
      <li class="header">
            <div class="pro-rank"></div>
            <div class="pro-img"></div>
            <div class="pro-name">商品名称</div>
            <div class="pro-price">单价</div>
            <div class="pro-evalu">分数</div>
            <div class="pro-place">来自</div>
      </li>
    <li class="product-list" v-for="(product,i) in toplist" :key="i">
      <div class="pro-rank">
        <template v-if="i === 0">
          <img src="../assets/imgs/first-img.png">
        </template>
        <template v-else-if="i === 1">
          <img src="../assets/imgs/second-img.png">
        </template>
        <template v-else-if="i === 2">
          <img src="../assets/imgs/third-img.png">
        </template>        
        <template v-else>
          <div class="pro-rank1">
          <span class="rankelse">{{ i + 1 }}</span>
        </div>
        </template>
      </div>

        <div class="pro-img">
          <router-link :to="{ path: '/goods/details', query: {productID:product.id} }">
            <img :src="product.img_path" />
          </router-link>
        </div>
        <div >
          <router-link class="pro-name1"
            :to="{ path: '/goods/details', query: {productID:product.product_id} }"
          >{{product.name}}</router-link>
        </div>
        <div class="pro-price1">￥{{product.discount_price}}</div>
        <div class="pro-evalu1">{{product.avg_score}}</div> 
        <div class="pro-place1">{{placemap[product.canteen_id]}}</div>
    </li>
    </ul>
  </div>
</div>
</div>
</template>

<script>
export default {
  data() {
    return {         
        toplist:[], 
        //榜单商品列表
        placemap:["无","榕园","岁月湖","荔园","若海","瑾园"]
    };
  },

  activated() {
    //请求榜单
    this.$axios
      .post("/api/rankings", {
      })
      .then(res => {
        if (res.data.status == "200") {
          this.toplist = res.data.data.items;
        } else {

          this.notifyError(res.data.msg);
        }
      })
      .catch(err => {
        return Promise.reject(err);
      });
  },
}
</script>

<style>
#container {
  width: 70%;
  margin: 0 auto;
  /*其他样式*/
}

.pcontent{
  left: 50%;
  width: auto;
}

.title h4 {
 font-family: STXinwei;
  color:white;
  text-align: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  font-size: 48px; 
  padding-top:14px;
 }  
.content{
  background: #ffffffcc;
}

.content .title{
  height: 76px;
  background-color:rgba(0, 88, 38, 0.7);
}

.content ul .header {
  height: 85px;
  padding-right: 26px;
  color: #424242;
}
 .content ul .product-list {
  height: 85px;
  padding: 15px 26px 15px 0;
  border-top: 1px solid #e0e0e0;
}
.pro-rank {
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 20px;
  padding-right: 40px;
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; 

}

.pro-rank1 {
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 20px;
  padding-right: 40px;
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; 
}

.content ul .pro-rank img{
  height: 80px;
  width: 80px;
}
.content ul .pro-rank span{
  float: left;
    height: 80px;
    width: 80px;
    padding-left: 20px;
    padding-bottom: -1px;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
    font-family: fangsong;
}
 .pro-img {
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 16px;
}



 .content ul .pro-img img {
  height: 80px;
  width: 80px;
}
 .pro-name {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 28px;

  color: black;
  font-family: STKaiti;
    font-size: 23px;
    font-weight: 900;
}

.product-list .pro-name1{
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 28px;
  font-size: 26px;
  color: black;
}

.pro-evalu1 {
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 123px;
  padding-top: 28px;
  font-size: 26px;
  color: black;
}

.content ul .pro-price {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 93px;

  color: black;
  font-family: STKaiti;
    font-size: 23px;
    font-weight: 900;
}

.product-list .pro-price1{
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 94px;
  font-size: 26px;
  color: black;
}


 .content ul .pro-name a {
  color: #424242;
}
 .content ul .pro-name a:hover {
  color: #ff6700;
}
.content ul .pro-evalu {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 78px;

  color: black;
  font-family: STKaiti;
    font-size: 23px;
    font-weight: 900;
}



.content ul .pro-place {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 79px;

  color: black;

  font-family: STKaiti;
    font-size: 23px;
    font-weight: 900;
}

.product-list .pro-place1 {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 35px;
  font-size: 26px;
  color: black;
}

.rankelse{
  font-size: 60px;
  height: 85px;
  width: 60px;
  padding-left: 80px;
  text-align: center;
}
</style>