<template>
        <el-card shadow="never">
            <!-- 评论头 -->
            <div shadow="never" >
                <span v-if="pUser">
                    <el-popover  placement="right" width="400" trigger="click">
                      <p>{{ user }}</p>
                      <el-button type="text" size="medium" style="font-family: 楷体" slot="reference">{{ user.name }}</el-button>
                    </el-popover>
                    回复
                    <el-popover placement="right"  width="400" trigger="click">
                      <p>{{ pUser }}</p>
                      <el-button type="text" size="medium" style="font-family: 楷体" slot="reference">{{pUser.name}}</el-button>
                    </el-popover>
                    {{ comment.commentDate }}
                </span>
                <span class="name" v-else>
                    <div class="photo-wall">
                        <img v-for="(photo, index) in comment.commentPhoto" 
                        :src="photo" 
                        :key="index"
                        alt="photo">
                    </div>                   
                    {{ comment.user_name }}
                    于{{ comment.created_at }}
                    发布评价
                </span>

                <!-- 星级 -->
                <div class="score">
                    <Star :score="Number(comment.score)"></Star>
                </div>
            </div>
            <!-- 评论内容 -->
            <p style="padding: 45px 0px 0px 0px;font-size: 30px;font-family: STKaiti;">{{ comment.ProductComment }}</p>           
        </el-card>
</template>

<script>
import Star from "../Star.vue";
export default {
    name:'commentInfo',
    components:{
        Star
    },
    props:{
        comment:{
            type:Object,
            request:true
        },
        pUser:{
            type:Object,
            request: true,
        }
    },
    data(){
        return {
            activeNames:['1'],
            user:{},
            showSon:false
        }
    },
    created() {
        this.getUser(this.comment.user_name)
        this.showSonInfo()
    }
}
</script>

<style scoped>
    span{
        padding-left: 10px;
        color: grey;
        font-size: 15px;

    }
    .el-card__body {
        padding: 0px 0px 0px 5px;
        border: 1px solid #bbbbbb;
        font-family: 楷体;
    }
    .score{
        padding-top: 10px;
    }
</style>
