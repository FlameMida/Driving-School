<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="用户角色ID:">
        <el-input v-model="formData.authorityId" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="hasServerInfo字段:">
        <el-switch v-model="formData.hasServerInfo" active-color="#13ce66" active-text="是" clearable
                   inactive-color="#ff4949" inactive-text="否"></el-switch>
      </el-form-item>

      <el-form-item label="用户头像:">
        <el-input v-model="formData.headerImg" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="用户昵称:">
        <el-input v-model="formData.nickName" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="用户登录密码:">
        <el-input v-model="formData.password" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="用户登录名:">
        <el-input v-model="formData.username" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="用户UUID:">
        <el-input v-model="formData.uuid" clearable placeholder="请输入"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="save">保存</el-button>
        <el-button type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {
  createCoach,
  updateCoach,
  findCoach
} from "@/api/coach";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";

export default {
  name: "Coach",
  mixins: [infoList],
  data() {
    return {
      type: "", formData: {
        authorityId: "",
        hasServerInfo: false,
        headerImg: "",
        nickName: "",
        password: "",
        username: "",
        uuid: "",

      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createCoach(this.formData);
          break;
        case "update":
          res = await updateCoach(this.formData);
          break;
        default:
          res = await createCoach(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findCoach({ID: this.$route.query.id})
      if (res.code == 0) {
        this.formData = res.data.recoach
        this.type == "update"
      }
    } else {
      this.type == "create"
    }

  }
};
</script>

<style>
</style>