<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img alt="" class="logo_login" src="@/assets/logo_login.png"/>
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">驾校管理系统</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form
            ref="loginForm"
            :model="loginForm"
            :rules="rules"
            @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
              ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
                v-model="loginForm.password"
                :type="lock === 'lock' ? 'password' : 'text'"
                placeholder="请输入密码"
            >
              <i
                  slot="suffix"
                  :class="'el-input__icon el-icon-' + lock"
                  @click="changeLock"
              ></i>
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative">
            <el-input
                v-model="loginForm.captcha"
                name="logVerify"
                placeholder="请输入验证码"
                style="width: 60%"
            />
            <div class="vPic">
              <img
                  v-if="picPath"
                  :src="picPath"
                  width="100%"
                  height="100%"
                  alt="请输入验证码"
                  @click="loginVerify()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%"
            >登 录
            </el-button
            >
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">
        <div class="title">个人空间</div>
        <div class="links">
          <a href="https://gitee.com/FlameMida" target="_blank">
            <img class="link-icon" src="@/assets/gitee.png"/>
          </a>
          <a href="https://www.yuque.com/FlameMida/" target="_blank">
            <img class="link-icon" src="@/assets/yuque.png"/>
          </a>
          <a href="https://github.com/Flamemida" target="_blank">
            <img class="link-icon" src="@/assets/github.png"/>
          </a>
          <a href="https://space.bilibili.com/106948903" target="_blank">
            <img class="link-icon" src="@/assets/bilibili.png"/>
          </a>
        </div>
        <div class="copyright">
          Copyright &copy; {{ curYear }} 💖 FlameMida
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {mapActions} from "vuex";
import {captcha} from "@/api/user";

export default {
  name: "Login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的用户名"));
      } else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };
    return {
      curYear: 0,
      lock: "lock",
      loginForm: {
        username: "admin",
        password: "123456",
        captcha: "",
        captchaId: "",
      },
      rules: {
        username: [{validator: checkUsername, trigger: "blur"}],
        password: [{validator: checkPassword, trigger: "blur"}],
      },
      logVerify: "",
      picPath: "",
    };
  },
  created() {
    this.loginVerify();
    this.curYear = new Date().getFullYear();
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    async login() {
      return await this.LoginIn(this.loginForm);
    },
    async submitForm() {
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          const flag = await this.login();
          if (!flag) {
            this.loginVerify();
          }
        } else {
          this.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true,
          });
          this.loginVerify();
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    loginVerify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/style/login.scss";
</style>
