<template>
    <div>
        <div class="product">
            <div class="pro-img">
                <img :src="$target + productDetails.product_picture" class="pro-img"/>
            </div>
            <div class="prodetail">
                <div class="content">
                    <h1 class="name">{{productDetails.product_name}}</h1>
                    <p class="intro">{{productDetails.product_intro}}</p>
                    <div class="price">
                      <span>{{productDetails.product_selling_price}}元</span>
                      <span
                        v-show="productDetails.product_price != productDetails.product_selling_price"
                        class="del"
                      >{{productDetails.product_price}}元</span>
                    </div>      
                </div>                
            </div>          
        </div>
    <div>
    
      <StarRating :rating="starIndex"/>
      <!-- 星级评价 -->
      <div class="review-form">
        <textarea placeholder="请输入评价内容" v-model="text"></textarea>
        <!-- 输入评价 -->
        <button @click="submit">提交</button>
        <!-- 提交按钮 -->
      </div>
    </div>

</div>
  </template>
  
  <script>
  import StarRating from '../components/StarRating.vue';
  
  export default {
    components: {
      StarRating,
    },
    data() {
      return {
        starIndex: 0,
        text: '',
      };
    },
    methods: {
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
    },
  };
  </script>
  