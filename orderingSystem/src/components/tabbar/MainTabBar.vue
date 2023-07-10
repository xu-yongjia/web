<!--
 * @Descripttion:顶部导航栏组件 
-->

<template>
    <tab-bar>
        <tab-bar-item path="/home">
            <template v-slot:item-icon><img class="logo-image" src="../../assets/imgs/sysu.png" alt=""></template>
            <template v-slot:item-icon-active><img class="logo-image" src="../../assets/imgs/sysu.png" alt=""></template>
        </tab-bar-item>
        <tab-bar-item path="/home" activeColor="mediumturquoise">
            <template v-slot:item-text><div class="title-style">首页</div></template>
            <template v-slot:item-icon><img src="" alt=""></template>
            <template v-slot:item-icon-active><img class="hengxian" src="../../assets/img/hengxian.png" alt=""></template>
        </tab-bar-item>
        <tab-bar-item path="/goods" activeColor="mediumturquoise">
            <template v-slot:item-text><div class="title-style">全部商品</div></template>
            <template v-slot:item-icon><img src="" alt=""></template>
            <template v-slot:item-icon-active><img class="hengxian" src="../../assets/img/hengxian.png" alt=""></template>
        </tab-bar-item>
        <tab-bar-item path="/community" activeColor="mediumturquoise">
            <template v-slot:item-text><div class="title-style">好评榜</div></template>
            <template v-slot:item-icon><img src="" alt=""></template>
            <template v-slot:item-icon-active><img class="hengxian" src="../../assets/img/hengxian.png" alt=""></template>
        </tab-bar-item>       
    </tab-bar>
</template>

<script>
    import TabBar from "./TabBar";
    import TabBarItem from "./TabBarItem";

    export default {
        name: "MainTabBar",
        components: {
            TabBar,
            TabBarItem
        }
    }
</script>

<style scoped>
    .title-style {
        font-weight: bold;
        position: relative;
        top: 20px;
        cursor: pointer;
    }
    .title-style:hover {
        color: #48d1cc;
    }

    .logo-image {
        width: 197px;
        position: relative;
        top: 9px
    }

    .hengxian {
        position: relative;
        top: 23px;
        width: 50px;
        height: 15px;
    }
</style>