<!--
 * @Description: 我的订单页面组件
 -->
 <template>
  <div class="order">
    <!-- 我的订单头部 -->
    <div class="order-header">
      <div class="order-header-content">
        <p>
          <i class="el-icon-s-order" style="font-size: 30px;color: #ff6700;"></i>
          我的订单
        </p>
      </div>
    </div>
    <!-- 我的订单头部END -->

    <!-- 我的订单主要内容 -->
    <div class="order-content" v-if="orders.length>0">
      <div class="content" v-for="(item,index) in orders" :key="index">
        <ul>
          <!-- 我的订单表头 -->
          <li class="order-info">
            <div class="order-id">订单编号: {{item.order_id}}</div>
            <div class="order-time">订单时间: {{item.created_at | dateFormat}}</div>
          </li>
          <li class="header">
            <div class="pro-img"></div>
            <div class="pro-name">商品名称</div>
            <!-- <div class="pro-state">订单状态</div> -->
            <div class="pro-price">单价</div>
            <div class="pro-num">数量</div>
            <div class="pro-total">小计</div>
            <div class="pro-state">订单状态</div>
            <div class="pro-courier">配送员</div>
            <div class="pro-courierphone">配送员电话</div>
            <div v-show="isfinished" class="pro-evalu"></div>
          </li>
          <!-- 我的订单表头END -->

          <!-- 订单列表 -->
          <li class="product-list" v-for="(product,i) in item.order_productlist" :key="i">
            <div class="pro-img">
              <router-link :to="{ path: '/goods/details', query: {productID:product.product_id} }">
                <img :src="product.img_path" />
              </router-link>
            </div>
            <div class="pro-name">
              <router-link
                :to="{ path: '/goods/details', query: {productID:product.product_id} }"
              >{{product.product_name}}</router-link>
            </div>
            <div class="pro-price">{{product.product_price}}元</div>
            <div class="pro-num">{{product.num}}</div>
            <div class="pro-total pro-total-in">{{product.product_price*product.num}}元</div>
            
            <!-- 这里的插值应该为“order.state”，上面的就是order.product.product_num等等 为显示字先写死-->
            <div class="pro-state">{{item.status}}</div>
            <div class="pro-courier">{{item.deliver_name}}</div>
            <div class="pro-courierphone">{{item.deliver_phone}}</div>
            
            <div v-show="item.status === '已送达'" class="pro-evalu">
            <!-- <div v-show="isfinished" class="pro-evalu">去评价</div> -->
            <router-link :to="{ path: '/judge', query: {orderID:item.order_id,productID:product.product_id} }" class="pro-evalu"> 
                去评价
                <i class="el-icon-d-arrow-right"></i>
            </router-link>
            </div>

            

            <div v-show="item.status === '未支付'" class="pro-evalu">
          
          <div  class="pro-evalu"> 
            <p @click="openweb(item.order_id)">
                去支付                
              <i class="el-icon-d-arrow-right"></i>
            </p>              
          </div>

          </div>

          <div v-show="item.status === '送餐中'" class="pro-evalu">
          
          <div  class="pro-evalu"> 
            <p @click="opensure(item.order_id)">
                确认收货                
            </p>              
          </div>

          </div>

          </li>
        </ul>
        <div class="order-bar">
          <div class="order-bar-left">
            <span class="order-total">
              共
              <span class="order-total-num">{{total[index].totalNum}}</span> 件商品
            </span>
          </div>
          <div class="order-bar-right">
            <span>
              <span class="total-price-title">合计：</span>
              <span class="total-price">{{total[index].totalPrice}}元</span>
            </span>
          </div>
          <!-- 订单列表END -->
        </div>
      </div>
      <div style="margin-top:-40px;"></div>
    </div>
    <!-- 我的订单主要内容END -->

    <!-- 订单为空的时候显示的内容 -->
    <div v-else class="order-empty">
      <div class="empty">
        <h2>您的订单还是空的！</h2>
        <p>快去选购吧！</p>
      </div>
    </div>
    <!-- 订单为空的时候显示的内容END -->
  </div>
</template>
<script>
export default {
  data() {
    return {
      orders: [], // 订单列表
      total: [], // 每个订单的商品数量及总价列表
      isfinished:true,
      type:'',
      source:''
    };
  },

  // 获取订单数据
  activated() {
    if (this.$route.query.type != undefined) {
      this.type = this.$route.query.type
    } else {
      this.type = '全部'
    }
    this.$axios
      .post("/api/users/getorder", {
        user_id: this.$store.getters.getUser.user_id,
        status:this.type
      })
      .then(res => {
        if (res.data.status === 200) {//status                    
         this.orders = res.data.data.orderlist;
        } else {
          this.notifyError(res.data.msg);
        }
      })
      .catch(err => {
        return Promise.reject(err);
      });
  },

  // 通过订单信息，计算出每个订单的商品数量及总价
  watch: {
    orders: function(val) {
     
      let total = [];
      
      for (let i = 0; i < val.length; i++) {
        const element = val[i];
     
        let totalNum = 0;
        let totalPrice = 0;
        
        for (let j = 0; j < element.order_productlist.length; j++) {
          const temp = element.order_productlist[j];
    
          totalNum += temp.num;
       
          totalPrice += temp.product_price * temp.num;
         
        }
        total.push({ totalNum, totalPrice });
      }
      this.total = total;
    }
  },

  methods:{
    openweb(orderid){
     
      this.$axios
        .post("/api/users/payments", {
           order_id:orderid
        })
        .then(res => {

          switch (res.data.status) {
            // “001”代表结算成功
            case 200:
            
              this.source=res.data.data
            
              window.location.href=this.source
              break;
            default:
              // 提示失败信息
              this.notifyError(res.data.msg);
          }
        })
        .catch(err => {
          return Promise.reject(err);
        });
      
       
    },

    opensure(orderid){
     
     this.$axios
       .post("/api/users/recieveOrder", {
          order_id:parseInt(orderid),
     

       })
       .then(res => {
         switch (res.data.status) {
          case 200:      
          this.notifySucceed(res.data.msg);
          
          this.$axios
      .post("/api/users/getorder", {
        user_id: this.$store.getters.getUser.user_id,
        status:this.type
      })
      .then(res => {
        if (res.data.status === 200) {//status                    
         this.orders = res.data.data.orderlist;
        } else {
          this.notifyError(res.data.msg);
        }
      })
      .catch(err => {
        return Promise.reject(err);
      });


          break;
          default:
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
<style scoped>
.order {
  background-color: #f5f5f5;
  padding-bottom: 20px;
}
/* 我的订单头部CSS */
.order .order-header {
  height: 64px;
  border-bottom: 2px solid #ff6700;
  background-color: #fff;
  margin-bottom: 20px;
}
.order .order-header .order-header-content {
  width: 1225px;
  margin: 0 auto;
}
.order .order-header p {
  font-size: 28px;
  line-height: 58px;
  float: left;
  font-weight: normal;
  color: #424242;
}
/* 我的订单头部CSS END */
.order .content {
  width: 1225px;
  margin: 0 auto;
  background-color: #fff;
  margin-bottom: 50px;
}

.order .content ul {
  background-color: #fff;
  color: #424242;
  line-height: 85px;
}
/* 我的订单表头CSS */
.order .content ul .order-info {
  height: 60px;
  line-height: 60px;
  padding: 0 26px;
  color: #424242;
  border-bottom: 1px solid #ff6700;
}
.order .content ul .order-info .order-id {
  float: left;
  color: #ff6700;
}
.order .content ul .order-info .order-time {
  float: right;
}

.order .content ul .header {
  height: 85px;
  padding-right: 26px;
  color: #424242;
}
/* 我的订单表头CSS END */

/* 订单列表CSS */
.order .content ul .product-list {
  height: 85px;
  padding: 15px 26px 15px 0;
  border-top: 1px solid #e0e0e0;
}
.order .content ul .pro-img {
  float: left;
  height: 85px;
  width: 120px;
  padding-left: 80px;
}
.order .content ul .pro-img img {
  height: 80px;
  width: 80px;
}

.order .content ul .pro-num {
  float: left;
  width: 132px;
}


.order .content ul .pro-name {
  float: left;
  width: 112px;
}
.order .content ul .pro-name a {
  color: #424242;
}
.order .content ul .pro-name a:hover {
  color: #ff6700;
}
.order .content ul .pro-price {
  float: left;
  width: 78px;
  padding-right: 18px;
  text-align: center;
}
.order .content ul .pro-state {
  float: left;
  width: 125px;
  padding-right: 18px;
  text-align: center;
}

.pro-courier {
  float: left;
  width: 86px;
  padding-right: 18px;
  text-align: center;
}

.pro-courierphone {
  float: left;
  width: 160px;
  padding-right: 18px;
  text-align: center;
}
.order .content ul .pro-evalu {
  float: left;
  width: 160px;
  padding-right: 18px;
  text-align: center;
  font-size: 15px; /* 字体大小 */
  color: rgb(237, 162, 111); /* 文字颜色 */
  text-decoration: none; /* 去掉下划线 */
}

.order .content ul .pro-evalu:hover {
  color: rgb(243, 83, 39);    /* 鼠标悬浮时文字颜色 */
  cursor: pointer;    /* 鼠标悬浮时鼠标形状 */
}

.order .content ul .pro-num {
  float: left;
  width: 88px;
  text-align: center;
}
.order .content ul .pro-total {
  float: left;
  width: 100px;
  /* padding-right: 10px; */
  text-align: center;
}
.order .content ul .pro-total-in {
  color: #ff6700;
}

.order .order-bar {
  width: 1185px;
  padding: 0 20px;
  border-top: 1px solid #ff6700;
  height: 50px;
  line-height: 50px;
  background-color: #fff;
}
.order .order-bar .order-bar-left {
  float: left;
}
.order .order-bar .order-bar-left .order-total {
  color: #757575;
}
.order .order-bar .order-bar-left .order-total-num {
  color: #ff6700;
}
.order .order-bar .order-bar-right {
  float: right;
}
.order .order-bar .order-bar-right .total-price-title {
  color: #ff6700;
  font-size: 14px;
}
.order .order-bar .order-bar-right .total-price {
  color: #ff6700;
  font-size: 30px;
}
/* 订单列表CSS END */

/* 订单为空的时候显示的内容CSS */
.order .order-empty {
  width: 1225px;
  margin: 0 auto;
}
.order .order-empty .empty {
  height: 300px;
  padding: 0 0 130px 558px;
  margin: 65px 0 0;
  background: url(../assets/imgs/cart-empty.png) no-repeat 124px 0;
  color: #b0b0b0;
  overflow: hidden;
}
.order .order-empty .empty h2 {
  margin: 70px 0 15px;
  font-size: 36px;
}
.order .order-empty .empty p {
  margin: 0 0 20px;
  font-size: 20px;
}
/* 订单为空的时候显示的内容CSS END */
</style>