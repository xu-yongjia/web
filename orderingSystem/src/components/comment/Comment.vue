<template>
    <div>
        <span class="text-center">用户评价</span>
        <h2>(共{{ commentList.length }}条评价)</h2>
        <!-- <hr/> -->
        <CommentInfo v-for="(comment,index) in commentList" :key="index" :comment="comment" :pUser="null" />
    </div>
</template>

<script>
    import CommentInfo from "./CommentInfo.vue";
export default {
    props:["commentList"],
    components:{
        CommentInfo,
    },
    data(){
        return {
            activeNames:['1'],
            user:{},
        }
    },
}
</script>

<style scoped>
/* CSS 样式表 */

.text-center{
  color:#0e510b;
  font-size: 60px;
  font-family: STXinwei;
  padding-left: 15px;
}

h1 {
  font-size: 32px;  /* 字体大小 */
  font-weight: bold;  /* 粗体字 */
  color: #dc9222;  /* 颜色 */
  margin-top: 30px;  /* 上外边距 */
  margin-bottom: 25px;  /* 下外边距 */
}
h2 {
    padding-top: 27px;
    padding-left: 25px;
    color: gray;
    margin-bottom: 44px; 
}
</style>