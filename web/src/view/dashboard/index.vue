<template>
  <div class="big">
    <el-row>
      <div class="card">
        <el-col :xs="24" :lg="16" :md="16">
          <div class="car-left">
            <el-row>
              <div>
                <el-col :xs="4" :md="3" :lg="3">
                  <span class="card-img">
                    <img :src="userInfo.headerImg" alt=""/>
                  </span>
                </el-col>
                <el-col :xs="20" :lg="12" :md="12">
                  <div class="text">
                    <h4><span>{{ TimeInfo }}</span>，{{ userInfo.authority.authorityName }}，{{ TimeMsg }}</h4>
                    <p class="tips-text">
                      <i class="el-icon-sunny el-icon"></i>
                      <span>当前天气为 {{ weather }} ，气温 {{ temp }} ℃。</span>
                    </p>
                  </div>
                </el-col>
              </div>
            </el-row>
          </div>
        </el-col>
        <el-col :lg='8' :md="8" :xs="32">
          <div class="car-right">
            <el-row>
              <el-col :span="6">
                <el-card shadow="hover">
                  <div class="car-item">
                    <span class="flow"><i class="el-icon-s-grid"></i></span>
                    <span>今日流量：</span>
                    <b>{{ ActiveUsers }}</b>
                  </div>
                </el-card>
              </el-col>

              <el-col :span="6">
                <el-card shadow="hover">
                  <div class="car-item">
                  <span class="user-number">
                    <i class="el-icon-s-custom"></i>
                  </span>
                    <span>总用户：</span>
                    <b>{{ TotalUsers }}</b>
                  </div>
                </el-card>
              </el-col>
              <el-col :span="6">
                <el-card shadow="hover">
                  <div class="car-item">
                  <span class="feedback">
                    <i class="el-icon-s-custom"></i>
                  </span>
                    <span>教练数：</span>
                    <b>{{ TotalCoaches }}</b>
                  </div>
                </el-card>
              </el-col>
              <el-col :span="6">
                <el-card shadow="hover">
                  <div class="car-item">
                  <span class="feedback">
                    <i class="el-icon-s-custom"></i>
                  </span>
                    <span>学员数：</span>
                    <b>{{ TotalStudents }}</b>
                  </div>
                </el-card>
              </el-col>
            </el-row>
          </div>
        </el-col>
      </div>
    </el-row>
    <el-row>
      <el-card shadow="hover">

        <div></div>
      </el-card>
    </el-row>
    <div class="shadow">
      <el-row :gutter="20">
        <el-col
            :span="4"
            v-for="(card, key) in toolCards"
            :key="key"
            @click.native="toTarget(card.name)"
            :xs="8"
        >
          <el-card shadow="hover" class="grid-content">
            <i :class="card.icon" :style="{ color: card.color }"></i>
            <p>{{ card.label }}</p>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div class="bottom">
      <!--      <el-row :gutter="24">-->
      <!--        <el-col :xs="24" :sm="24" :lg="24">-->
      <!--          <div class="chart-player">-->
      <!--            <musicPlayer/>-->
      <!--          </div>-->
      <!--        </el-col>-->
      <!--        <el-col :xs="24" :sm="24" :lg="12">-->
      <!--          <div class="chart-player">-->
      <!--            <to-do-list />-->
      <!--          </div>-->
      <!--        </el-col>-->
      <!--      </el-row>-->
      <state></state>
    </div>
    <div>

    </div>
  </div>
</template>

<script>
import state from "../system/state.vue"
// import musicPlayer from "./component/musicPlayer";
// import TodoList from "./component/todoList";
import {mapGetters} from "vuex";
import {getDashboardState, getWeather} from "@/api/system";

export default {
  name: "Dashboard",
  created() {
    this.reloadDashBoard();
    this.timer = setInterval(() => {
      this.reloadDashBoard();
    }, 1000 * 1800);
  },
  mounted() {
    this.getLocation();
  },
  beforeDestroy() {
    clearInterval(this.timer)
    this.timer = null
  },
  data() {
    return {
      weather: "晴",
      temp: "24",
      latitude: 0,
      longitude: 0,
      timer: null,
      TimeInfo: "欢迎 ",
      TimeMsg: "进入管理面板~",
      ActiveUsers: "0",
      TotalUsers: "0",
      TotalStudents: "0",
      TotalCoaches: "0",
      toolCards: [
        {
          label: "用户管理",
          icon: "el-icon el-icon-monitor",
          name: "user",
          color: "#ff9c6e",
        },
        {
          label: "角色管理",
          icon: "el-icon el-icon-setting",
          name: "authority",
          color: "#69c0ff",
        },
        {
          label: "菜单管理",
          icon: "el-icon el-icon-menu",
          name: "menu",
          color: "#b37feb",
        },
        {
          label: "学员管理",
          icon: "el-icon el-icon-cpu",
          name: "studentM",
          color: "#ffd666",
        },
        {
          label: "教练管理",
          icon: "el-icon el-icon-document-checked",
          name: "coachM",
          color: "#ff85c0",
        },
        {
          label: "关于驾校",
          icon: "el-icon el-icon-user",
          name: "about",
          color: "#5cdbd3",
        },
      ],
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
  },
  components: {
    state,
    // musicPlayer, //音乐播放器
    // TodoList, //TodoList
    // RaddarChart, //雷达图
    // stackMap, //堆叠图
    // Sunburst, //旭日图
  },
  methods: {
    getLocation() {
      let log, lat
      if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(
            function (position) {
              log = position.coords.longitude
              lat = position.coords.latitude
            },
            function (e) {
              throw(e.message);
            }
        )
        this.loadWeather(log, lat)
      }

    },
    async loadWeather(lat, log) {
      const {data} = await getWeather(lat, log)
      let weather = JSON.parse(data.weather)
      this.weather = weather.current.weather[0].description
      this.temp = weather.current.temp

    },
    toTarget(name) {
      this.$router.push({name});
    },
    async reloadDashBoard() {
      const {data} = await getDashboardState();
      this.TimeInfo = data.dashboard.time_info;
      this.TimeMsg = data.dashboard.time_msg;
      this.ActiveUsers = data.dashboard.active_users;
      this.TotalUsers = data.dashboard.total_users;
      this.TotalStudents = data.dashboard.total_students;
      this.TotalCoaches = data.dashboard.total_coaches;

    },
  },
};
</script>

<style lang="scss" scoped>
.big {
  margin: 100px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;

  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;

    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }

  .mid {
    width: 100%;
    height: 380px;

    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }

  .bottom {
    width: 100%;
    height: 300px;
    // margin: 20px 0;
    .el-row {
      margin-right: 4px !important;
    }

    .chart-player {
      width: 100%;
      height: 270px;
      padding: 10px;
      background-color: #fff;
    }
  }
}
</style>
