<template>
    <div class="star-rating">
      <img v-for="i in 5" :key="i" :src="`../assets/${starUrl(i)}.png`" @click="onStarClick(i)">
      <span>{{ ratingText }}</span>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      rating: {
        type: Number,
        required: true,
        validator: (value) => value >= 0 && value <= 5,
      },
    },
    computed: {
            starUrl() {
    return index => {
      return index <= this.rating ? 'full' : 'empty';
    };
  },
    
      ratingText() {
        switch (this.rating) {
          case 1:
            return '很差';
          case 2:
            return '较差';
          case 3:
            return '一般';
          case 4:
            return '很好';
          case 5:
            return '非常好';
          default:
            return '';
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .star-rating {
    position: relative;
  }
  
  .star-rating img {
    width: 30px;
    height: 30px;
    margin-right: 10px;
    cursor: pointer;
  }
  
  .star-rating span {
    position: absolute;
    left: 46px;
    top: 8px;
    color: #bbce12;
    font-size: 14px;
  }
  </style>
  