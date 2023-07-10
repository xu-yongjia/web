<!--
 * @Descripttion:星级评价显示组件
-->

<template> 
        <div class="star">
            <span class="star-item" v-for="(itemClass, index) in itemClasses" :key="index" :class="itemClass">
            </span>
            <span class="s">{{ score }}</span>
        </div>
</template>

<script>
    // 星星长度 
const LENGTH = 5

// 星星的状态
const CLS_ON = 'on'
const CLS_HALF = 'half'
const CLS_OFF = 'off'

export default {
    name: 'Star',
    props: {
        score: Number
    },
    computed: {
        itemClasses () {
            let result = []
            let score = Math.floor(this.score * 2) / 2

            // 半星 (通过跟1取余判断是否为小数)
            let hasDecimal = score % 1 !== 0

            // 全星 （向下取整，获取全星部分）
            let integer = Math.floor(score)

            // 遍历全星
            for(let i = 0; i < integer; i++){
                result.push(CLS_ON)
            }

            // 处理半星
            if(hasDecimal){
                result.push(CLS_HALF)
            }

            // 补齐
            while(result.length < LENGTH){ // 到这里还不够五颗星，则凑空星
                result.push(CLS_OFF)
            }  
            return result
        }
    }
}

</script>
<style>
.star{
    height: 30px;
    padding-top: 5px;
}

.star .s{
  padding-top: 5px;
  display: inline-block;
  margin: 10px;
  margin-top: 10px;
  margin-left: 10px;
  font-size: 28px;
  color: rgb(225, 149, 34);
} 


.star-item {
  display: inline-block;
  width: 26px;
  height: 24px;
  background-size: contain;
}

.star-item.on {
    background-image: url(../assets/half.png);
}
.star-item.half {
    background-image: url(../assets/full.png);
}
.star-item.off {
    background-image: url(../assets/empty.png);
}
</style>