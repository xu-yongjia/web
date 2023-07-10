<!--
 * @Description: 首页组件
 -->
<template>
  <div class="home" id="home" name="home">
    <div class="layer">
      <p class="font">今日推荐</p>
      <br>          
      <ul>
        <li v-for="(item,index) of border" :key="index" class="recomm">         
          <span>
          *{{ item }}
          </span>
          <br>
          <br>
        </li>        
      </ul>
  </div>
    
    <!-- 轮播图 -->  
    <div class="block">
      <div class="xinpin">
        <span>新品推荐</span>
      </div>
      <div >
      <el-carousel height="460px" >
        <el-carousel-item v-for="item in carousel" :key="item.id">
       
          <img style="height:460px;" :src="item.img_path" />       
        </el-carousel-item>
      </el-carousel>
      </div>
    </div>
    <!-- 轮播图END -->
    <div class="main-box">
      <div class="main">
      

        <!-- 地点展示区域 -->
        <div class="phone">
          <div class="box-hd">
            <div class="title">前往点单</div>
           
          </div>
          <div class="box-bd">
           
            <div class="list">
              <MyList :list="placelist" :isMore="true" :isnotPlace="true"></MyList>
            </div>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      carousel: "", // 轮播图数据
      phoneList: "", // 手机商品列表
      miTvList: "", // 小米电视商品列表
      applianceList: "", // 家电商品列表
      applianceHotList: "", //热门家电商品列表
      accessoryList: "", //配件商品列表
      accessoryHotList: "", //热门配件商品列表
      protectingShellList: "", // 保护套商品列表
      chargerList: "", //充电器商品列表
      applianceActive: 1, // 家电当前选中的商品分类
      accessoryActive: 1, // 配件当前选中的商品分类
      placelist:"",
      border:"111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
    };
  },
 
  created() {
    // 获取轮播图数据
    this.$axios
      .get("/api/carousels", {})
      .then(res => {
        this.carousel = res.data.data.carousel;
        this.border = res.data.data.board
      })
      .catch(err => {
     
        return Promise.reject(err);
      });
      this.getplacemessage()
     
  },
  methods: {
    

    //地点信息
    getplacemessage(){
      var one = { 
        product_picture: "../assets/imgs/rongyuan.jpg",
        product_name:"榕园",
        product_title:"1111",
        place_id:"1"
      }
      var two = { 
        product_picture: "../assets/imgs/suiyuehu.jpg",
        product_name:"岁月湖",
        product_title:"1111",
        place_id:"2"
      }
      var three = { 
        product_picture: "../assets/imgs/liyuan.jpg",
        product_name:"荔园",
        product_title:"1111",
        place_id:"3"
      }
      var four = { 
        product_picture: "../assets/imgs/ruohai.jpg",
        product_name:"若海",
        product_title:"1111",
        place_id:"4"
      }
      var five = { 
        product_picture: "../assets/imgs/jinyuan.jpg",
        product_name:"瑾园",
        product_title:"1111",
        place_id:"5"
      }
      var l=[one,two,three,four,five]
      this.placelist=l;
    },

    
  }
};
</script>
<style scoped>
@import "../assets/css/index.css";
/* .block {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
} */

.font{
  color: #ded5d5;
    font-size: 27px;
    text-align: center;
    font-family: KaiTi;
}

.board{
  width: 300px;
  height: 400px;
  background-color: #114938;
}

.xinpin{
  padding-top: 12px;
  padding-bottom: 15px;
}


.title{
  color: #114938;
  font-size: 58px;
  padding-left: 10px;
  font-family: STXinwei;
  font-weight: 600;
  text-shadow: 5px 2px 3px rgb(79 76 76 / 50%);
 padding-bottom: 20px;
}

.home{
  background-color: #ffffffbf;
}

.block{
  max-width: 695px
}

 .block span {
  font-family: STXinwei;
  font-size: 50px;
  text-align: center; /* 水平居中 */
  /* line-height: 100vh; 垂直居中 */
  margin-bottom: 20px;
  color: rgb(4, 57, 11);
  padding-left: 249px;
  
 }  

 element.style {
    height: 547px;
}

.carousel-container {
  width:  1000px; /* 设置父容器的宽度 */
}
.layer{
  position: absolute;
  top: 30%;
    left: 5%;
    /* z-index: 9999; */
    width: 336px;
    height: auto;
    white-space: normal;
    word-wrap: break-word;
    background-color: #20ba6cde;
    border-radius: 27px;
  
}

.layer:hover{
  z-index: 2;
  -webkit-box-shadow: 0 15px 30px rgba(0, 0, 0, 0.1);
  box-shadow: 0 15px 30px rgba(0, 0, 0, 0.1);
  -webkit-transform: translate3d(0, -2px, 0);
  transform: translate3d(0, -2px, 0);
}

.recomm{
  font-family: STXinwei;
    font-size: 34px;
    color: white;
}
</style>