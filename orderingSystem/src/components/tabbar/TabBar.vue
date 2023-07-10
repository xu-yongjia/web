<template>
    <div id="tab-bar">
        <slot></slot>
    </div>
</template>

<script>
    export default {
        name: "TabBar",
    }
</script>

<style scoped>
    #tab-bar {
        position: relative;
        z-index: 1;
        display: flex;
        /*filter: blur(1px);*/
        background: #fff;
        box-shadow: -3px 0 1px rgba(100,100,100,.2);
    }
</style>