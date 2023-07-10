<!--
 * @Description: 项目根组件
 -->
 <template>
  <div id="app" name="app">
    <el-container>
      

    <!-- 顶部导航栏 -->
    <div class="topbar" v-show="$route.meta.showMenu!==false">
        <div class="nav">
          <ul>
            <li v-if="!this.$store.getters.getUser">
              <div style="font-size:14px;display:inline-block;align-items: center;">
                <el-button type="text" @click="login">登录</el-button>
                <!-- <router-link to="/login">登录</router-link> -->
              </div>
            </li>
            <li v-else class="header-user-con">
              <!-- 用户头像 -->
              <div class="user-avator">
                <img :src="this.$store.getters.getUser.avatar" />

              </div>
              <!-- 用户名下拉菜单 -->
              <div class="user-name">
                <el-dropdown>
                  <span class="el-dropdown-link">
                    {{this.$store.getters.getUser.nickname}}
                    <i class="el-icon-caret-bottom"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    
                    <router-link to="/center">
                      <el-dropdown-item>个人中心</el-dropdown-item>
                    </router-link>
                    <a @click="logout">
                      <el-dropdown-item>退出登录</el-dropdown-item>
                    </a>
                  </el-dropdown-menu>
                </el-dropdown>
              </div>
            </li>
            <li class="font">
              <router-link to="/collect">我的收藏</router-link>
            </li>
            <li class="font">
              <router-link to="/order">我的订单</router-link>
            </li>

            <li :class="getNum > 0 ? 'shopCart-full' : 'shopCart'">
              <div style="">
                <router-link to="/shoppingCart">
                  <i class="el-icon-shopping-cart-full"></i>购物车
                  <span>({{getNum}})</span>
                </router-link>
              </div>
            </li>
          </ul>
        </div>
      </div>
      <!-- 顶部导航栏END -->

      <main-tab-bar></main-tab-bar>

      <!-- 登录模块 -->
      <MyLogin></MyLogin>
      <!-- 注册模块 -->
      <MyRegister :register="register" @fromChild="isRegister"></MyRegister>

      <!-- 主要区域容器 -->
      <el-main>
        <keep-alive>
          <router-view></router-view>
        </keep-alive>
      </el-main>
      <!-- 主要区域容器END -->

      <!-- 底栏容器 -->
      <el-footer>
        <div class="footer">
          <div class="ng-promise-box">
            <div class="ng-promise">
              <p class="text">
                <a class="icon1" href="javascript:;">免费送货</a>
                <a class="icon2" href="javascript:;">准时配送</a>
                <a class="icon3" style="margin-right: 0" href="javascript:;">质量第一，用心服务</a>
              </p>
            </div>
          </div>
          <div class="mod_help">
            <p>
              <router-link to="/">首页</router-link>
              <span>|</span>
              <router-link to="/goods">全部商品</router-link>
              <span>|</span>
              <router-link to="/community">好评榜</router-link>
            </p>
            <p class="coty">商城版权所有 &copy; 2012-2021</p>
          </div>
        </div>
      </el-footer>
      <!-- 底栏容器END -->
    </el-container>
    <div class="animation-container">
  <img class="image" src="../src/assets/imgs/longmao.gif" alt="动图" />
</div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { mapGetters } from "vuex";
import MainTabBar from "./components/tabbar/MainTabBar";
export default {
  name:"App",
  components:{
    MainTabBar
  },
  beforeUpdate() {
    this.activeIndex = this.$route.path;
  },
  data() {
    return {
      activeIndex: "", // 头部导航栏选中的标签
      search: "", // 搜索条件
      register: false, // 是否显示注册组件
      visible: false // 是否退出登录
    };
  },
  created() {
    // 获取浏览器localStorage，判断用户是否已经登录
    if (localStorage.getItem("user")) {
      // 如果已经登录，设置vuex登录状态
      this.setUser(JSON.parse(localStorage.getItem("user")));
    } 
  },
  computed: {
    ...mapGetters(["getUser", "getNum"])
  },
  watch: {
    // 获取vuex的登录状态
    getUser: function(val) {
      if (val === "") {
        // 用户没有登录
        this.setShoppingCart([]);
      } else {
        // 用户已经登录,获取该用户的购物车信息
        this.$axios
          .post("/api/users/getCart", {
            user_id: val.user_id
          })
          .then(res => {
            if (res.data.status === 200) {
              // 001 为成功, 更新vuex购物车状态
              this.setShoppingCart(res.data.data.carts);
            } else {
              // 提示失败信息
              this.notifyError(res.data.msg);
            }
          })
          .catch(err => {
            return Promise.reject(err);
          });
      }
    }
  },
  methods: {
    ...mapActions(["setUser", "setShowLogin", "setShoppingCart"]),
    login() {
      // 点击登录按钮, 通过更改vuex的showLogin值显示登录组件
      this.setShowLogin(true);
      
    },
    // 退出登录
    logout() {
      this.visible = false;
      // 清空本地登录信息
      localStorage.setItem("user", "");
      // 清空vuex登录信息
      this.setUser("");
      this.notifySucceed("成功退出登录");
    },
    // 接收注册子组件传过来的数据
    isRegister(val) {
      this.register = val;
    },
    // 点击搜索按钮
    searchClick() {
      if (this.search != "") {
        // 跳转到全部商品页面,并传递搜索条件
        this.$router.push({ path: "/goods", query: { search: this.search } });
        this.search = "";
      }
    }
  }
};
</script>

<style>
/* 全局CSS */
* {
  padding: 0;
  margin: 0;
  border: 0;
  list-style: none;
}

.topbar .nav li {
    float: left;
    height: 50px;
    color: #b0b0b0;
    font-size: 16px;
    text-align: center;
    line-height: 40px;
    margin-left: 20px;
}

.shopCart{
  background-color:#014824;

}


.header-user-con {
  display: flex;
  height: 50px;
  align-items: center;
}
.user-name {
  margin-left: 10px;
}
.user-avator img {
  display: block;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.el-dropdown-link {
  color: #fff;
  cursor: pointer;
}

.el-dropdown-menu__item {
  text-align: center;
}
.el-dropdown-item:hover {
  color: #ff6700;
}
#app .el-header {
  padding: 0;
}
#app .el-main {
  min-height: 300px;
  /* padding: 20px 0; */
  background-image: url(assets/imgs/sysuback.jpg);
}
#app .el-footer {
  padding: 0;
}
a,
a:hover {
  text-decoration: none;
}
/* 全局CSS END */

/* 顶部导航栏CSS */
.topbar {
  height: 40px;
  background-color: #014824;
  /* margin-bottom: 20px; */
}
.topbar .nav {
  width: 1225px;
  margin: 0 auto;
}
.topbar .nav ul {
  float: right;
}
.topbar .nav li {
  float: left;
  height: 40px;
  color: #b0b0b0;
  font-size: 14px;
  text-align: center;
  line-height: 40px;
  margin-left: 20px;
}
.topbar .nav .sep {
  color: #b0b0b0;
  font-size: 12px;
  margin: 0 5px;
}
.topbar .nav li .el-button {
  color: #b0b0b0;
}
.topbar .nav .el-button:hover {
  color: #fff;
}
.topbar .nav li a {
  color: #b0b0b0;
}
.topbar .nav a:hover {
  color: #fff;
}
.topbar .nav .shopCart {
  width: 120px;
  background: #014824;
}
.topbar .nav .shopCart:hover {
  background: #fff;
}
.topbar .nav .shopCart:hover a {
  color: #ff6700;
}
.topbar .nav .shopCart-full {
  width: 120px;
  background: #ff6700;
}
.topbar .nav .shopCart-full a {
  color: white;
}
/* 顶部导航栏CSS END */

/* 顶栏容器CSS */
.el-header .el-menu {
  max-width: 1225px;
  margin: 0 auto;
}
.el-header .logo {
  height: 60px;
  width: 189px;
  float: left;
  margin-right: 100px;
}
.el-header .so {
  margin-top: 10px;
  width: 300px;
  float: right;
}
/* 顶栏容器CSS END */

/* 底栏容器CSS */
.footer {
  width: 100%;
  text-align: center;
  background: #014824;
  padding-bottom: 20px;
}
.footer .ng-promise-box {
  border-bottom: 1px solid #3d3d3d;
  line-height: 145px;
}
.footer .ng-promise-box {
  margin: 0 auto;
  border-bottom: 1px solid #3d3d3d;
  line-height: 145px;
}
.footer .ng-promise-box .ng-promise p a {
  color: #fff;
  font-size: 20px;
  margin-right: 210px;
  padding-left: 44px;
  height: 40px;
  display: inline-block;
  line-height: 40px;
  text-decoration: none;
  /* background: url("./assets/imgs/us-icon.png") no-repeat left 0; */
}
.footer .github {
  height: 50px;
  line-height: 50px;
  margin-top: 20px;
}
.footer .github .github-but {
  width: 50px;
  height: 50px;
  margin: 0 auto;
  /* background: url("./assets/imgs/github.png") no-repeat; */
}
.footer .mod_help {
  text-align: center;
  color: #888888;
}
.footer .mod_help p {
  margin: 20px 0 16px 0;
}

.footer .mod_help p a {
  color: #888888;
  text-decoration: none;
}
.footer .mod_help p a:hover {
  color: #fff;
}
.footer .mod_help p span {
  padding: 0 22px;
}
/* 底栏容器CSS END */


.animation-container {
    position: fixed;
    bottom: 82px;
    right: 251px;
    width: 102px;
    height: 125px;
}

.image {
    position: fixed;
    bottom: 14px;
    right: 38px;
    width: 136px;
    height: 136px;
}
</style>