<template>
  <div>
    <el-row>
      <el-col :span="6">
        <div class="fl-left avatar-box">
          <div class="user-card">
            <div
                :style="{ 'background-image': `url(${(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg})`,'background-repeat':'no-repeat','background-size':'cover' }"
                class="user-headpic-update">
              <span class="update" @click="openChooseImg">
                <i class="el-icon-edit"></i>
                重新上传</span>
            </div>
            <div class="user-personality">
              <p class="nickname">{{ userInfo.nickName }}</p>
              <p class="person-info"></p>
            </div>
            <div class="user-information">
              <ul>
                <li>
                  <i class="el-icon-user"></i>{{ userInfo.nickName }}
                </li>
                <li>
                  <i class="el-icon-phone-outline"></i>{{ userInfo.phone.length > 0 ? userInfo.phone : "暂无联系方式" }}
                </li>
                <li>
                  <i class="el-icon-data-analysis"></i>广东培正学院驾校
                </li>
                <li>
                  <i class="el-icon-video-camera-solid"></i>中国·广东·广州市
                </li>
                <!--                <li>-->
                <!--                  <i class="el-icon-medal-1"></i>-->
                <!--                </li>-->
              </ul>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="18">
        <div class="user-addcount">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="账号绑定" name="second">
              <ul>
                <li>
                  <p class="title">修改手机号</p>
                  <p class="desc">
                    修改手机号
                    <a href="#" @click="showPhone=true">修改手机号</a>
                  </p>
                </li>
                <li>
                  <p class="title">修改密码</p>
                  <p class="desc">
                    修改个人密码
                    <a href="#" @click="showPassword=true">修改密码</a>
                  </p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-col>

      <el-col :span="18">

      </el-col>
    </el-row>


    <ChooseImg ref="chooseImg" @enter-img="enterImg"/>

    <el-dialog :visible.sync="showPassword" @close="clearPassword" title="修改密码" width="360px">
      <el-form :model="pwdModify" :rules="rules" label-width="80px" ref="modifyPwdForm">
        <el-form-item :minlength="6" label="原密码" prop="password">
          <el-input show-password v-model="pwdModify.password"></el-input>
        </el-form-item>
        <el-form-item :minlength="6" label="新密码" prop="newPassword">
          <el-input show-password v-model="pwdModify.newPassword"></el-input>
        </el-form-item>
        <el-form-item :minlength="6" label="确认密码" prop="confirmPassword">
          <el-input show-password v-model="pwdModify.confirmPassword"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="showPassword=false">取 消</el-button>
        <el-button @click="savePassword" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="showPhone" title="修改手机号" width="360px" @close="clearPhone">
      <el-form ref="modifyPhoneForm" :model="phoneModify" :rules="rules_phone" label-width="80px">
        <el-form-item :minlength="11" label="原手机" prop="phone">
          <el-input v-model="phoneModify.phone" show-phone></el-input>
        </el-form-item>
        <el-form-item :minlength="11" label="新手机" prop="newPhone">
          <el-input v-model="phoneModify.newPhone" show-phone></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="showPhone=false">取 消</el-button>
        <el-button type="primary" @click="savePhone">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import ChooseImg from "@/components/chooseImg";
import {changePassword, changePhone, setUserInfo} from "@/api/user";

import {mapGetters, mapMutations} from "vuex";

const path = process.env.VUE_APP_BASE_API;
export default {
  name: "Person",
  data() {
    return {
      path: path,
      activeName: "second",
      showPassword: false,
      showPhone: false,
      pwdModify: {},
      phoneModify: {},
      rules: {
        password: [
          {required: true, message: "请输入密码", trigger: "blur"},
          {min: 6, message: "最少6个字符", trigger: "blur"}
        ],
        newPassword: [
          {required: true, message: "请输入新密码", trigger: "blur"},
          {min: 6, message: "最少6个字符", trigger: "blur"}
        ],
        confirmPassword: [
          {required: true, message: "请输入确认密码", trigger: "blur"},
          {min: 6, message: "最少6个字符", trigger: "blur"},
          {
            validator: (rule, value, callback) => {
              if (value !== this.pwdModify.newPassword) {
                callback(new Error("两次密码不一致"));
              } else {
                callback();
              }
            },
            trigger: "blur"
          }
        ]
      },
      rules_phone: {
        phone: [
          {required: false, message: "请输入原手机号", trigger: "blur"},
          {min: 0, max: 11, message: "最少11个字符", trigger: "blur"},
        ],
        newPhone: [
          {required: true, message: "请输入新手机号", trigger: "blur"},
          {min: 11, max: 11, message: "最少11个字符", trigger: "blur"},
        ],

      },
    };
  },
  components: {
    ChooseImg
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"])
  },
  methods: {
    ...mapMutations("user", ["ResetUserInfo"]),
    savePassword() {
      this.$refs.modifyPwdForm.validate(valid => {
        if (valid) {
          changePassword({
            username: this.userInfo.username,
            password: this.pwdModify.password,
            newPassword: this.pwdModify.newPassword
          }).then((res) => {
            if (res.code === 0) {
              this.$message.success("修改密码成功！");
            }
            this.showPassword = false;
          });
        } else {
          return false;
        }
      });
    },
    savePhone() {
      this.$refs.modifyPhoneForm.validate(valid => {
        if (valid) {
          changePhone({
            username: this.userInfo.username,
            phone: this.phoneModify.phone,
            newPhone: this.phoneModify.newPhone
          }).then((res) => {
            if (res.code === 0) {
              this.$message.success("修改手机号成功！");
            }
            this.showPhone = false;
          });
        } else {
          return false;
        }
      });
    },
    clearPassword() {
      this.pwdModify = {
        password: "",
        newPassword: "",
        confirmPassword: ""
      };
      this.$refs.modifyPwdForm.clearValidate();
    },
    clearPhone() {
      this.phoneModify = {
        phone: "",
        newPhone: "",
      };
      this.$refs.modifyPhoneForm.clearValidate();
    },
    openChooseImg() {
      this.$refs.chooseImg.open();
    },
    async enterImg(url) {
      const res = await setUserInfo({headerImg: url, ID: this.userInfo.ID});
      if (res.code === 0) {
        this.ResetUserInfo({headerImg: url});
        this.$message({
          type: "success",
          message: "设置成功"
        });
      }
    },
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>
<style lang="scss">
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}

.avatar {
  width: 178px;
  height: 178px;
  display: block;
}

.avatar-box {
  box-shadow: -2px 0 20px -16px;
  width: 80%;
  height: 100%;

  .user-card {
    min-height: calc(90vh - 200px);
    padding: 30px 20px;
    text-align: center;

    .el-avatar {
      border-radius: 50%;
    }

    .user-personality {
      padding: 24px 0;
      text-align: center;

      p {
        font-size: 16px;
      }

      .nickname {
        font-size: 26px;
      }

      .person-info {
        margin-top: 6px;
        font-size: 14px;
        color: #999
      }
    }

    .user-information {
      width: 100%;
      height: 100%;
      text-align: left;

      ul {
        display: inline-block;
        height: 100%;

        li {
          i {
            margin-right: 8px;
          }

          padding: 20px 0;
          font-size: 16px;
          font-weight: 400;
          color: #606266;
        }
      }
    }
  }
}

.user-addcount {
  ul {
    li {
      .title {
        padding: 10px;
        font-size: 18px;
        color: #696969;
      }

      .desc {
        font-size: 16px;
        padding: 0 10px 20px 10px;
        color: #a9a9a9;

        a {
          color: rgb(64, 158, 255);
          float: right;
        }
      }

      border-bottom: 2px solid #f0f2f5;
    }
  }
}

.user-headpic-update {
  width: 120px;
  height: 120px;
  line-height: 120px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  border-radius: 20px;

  &:hover {
    color: #fff;
    background: linear-gradient(to bottom, rgba(255, 255, 255, 0.15) 0%, rgba(0, 0, 0, 0.15) 100%), radial-gradient(at top center, rgba(255, 255, 255, 0.40) 0%, rgba(0, 0, 0, 0.40) 120%) #989898;
    background-blend-mode: multiply, multiply;

    .update {
      color: #fff;
    }
  }

  .update {
    height: 120px;
    width: 120px;
    text-align: center;
    color: transparent;
  }
}
</style>