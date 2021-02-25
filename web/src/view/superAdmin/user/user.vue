<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户名">
          <el-input placeholder="搜索条件" v-model="searchInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="姓名">
          <el-input placeholder="搜索条件" v-model="searchInfo.nickname"></el-input>
        </el-form-item>
        <el-form-item label="用户角色" prop="authorityId">
          <el-cascader
              v-model="searchInfo.authorityId"
              placeholder="请选择用户角色"
              :options="authOptions"
              :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              filterable></el-cascader>
        </el-form-item>

        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item >
          <el-button @click="addUser" type="primary">新增用户</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
        :data="tableData"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark">
      <el-table-column label="头像" width="100">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <CustomPic :picSrc="scope.row.headerImg"/>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="uuid" min-width="250" prop="uuid"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="userName"></el-table-column>
      <el-table-column label="姓名" min-width="150" prop="nickName"></el-table-column>

      <el-table-column label="用户角色" min-width="150">
        <template slot-scope="scope">
          <el-cascader
              @change="changeAuthority(scope.row)"
              v-model="scope.row.authority.authorityId"
              :options="authOptions"
              :show-all-levels="false"
              :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              filterable
          ></el-cascader>
        </template>
      </el-table-column>

      <el-table-column label="操作" min-width="150">
        <template slot-scope="scope">
          <el-button @click="toDetail(scope.row)" size="small" type="success">详情</el-button>
          <el-button @click="updateUser(scope.row)" size="small" type="primary">变更</el-button>
            <el-popover placement="top" width="160" v-model="scope.row.visible" style="margin-left: 10px">
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
              </div>
              <el-button type="danger" icon="el-icon-delete" size="small" slot="reference">删除</el-button>
            </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>

    <el-dialog :visible.sync="addUserDialog" custom-class="user-dialog" title="新增用户">
      <el-form :rules="rules" ref="userForm" :model="userInfo">
        <el-form-item label="用户名" label-width="80px" prop="username">
          <el-input v-model="userInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input v-model="userInfo.password"></el-input>
        </el-form-item>
        <el-form-item label="别名" label-width="80px" prop="nickName">
          <el-input v-model="userInfo.nickName"></el-input>
        </el-form-item>
        <el-form-item label="头像" label-width="80px">
          <div style="display:inline-block" @click="openHeaderChange">
            <!-- TODO:url错误-->
            <img class="header-img-box" v-if="userInfo.headerImg" :src="'api'+userInfo.headerImg" />
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </el-form-item>
        <el-form-item label="用户角色" label-width="80px" prop="authorityId">
          <el-cascader
              v-model="userInfo.authorityId"
              :options="authOptions"
              :show-all-levels="false"
              :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              filterable
          ></el-cascader>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddUserDialog">取 消</el-button>
        <el-button @click="enterAddUserDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :targetKey="`headerImg`"/>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import {getAuthorityList} from "@/api/authority";
import infoList from "@/mixins/infoList";
import {mapGetters} from "vuex";
import CustomPic from "@/components/customPic";
import ChooseImg from "@/components/chooseImg";
import {deleteUser, getUserList, register, setUserAuthority} from "@/api/user";

export default {
  name: "User",
  mixins: [infoList],
  components: {CustomPic, ChooseImg},
  data() {
    return {
      listApi: getUserList,
      dialogFormVisible: false,
      type: "",
      path: path,
      authOptions: [],
      addUserDialog: false,
      userInfo: {
        username: "",
        password: "",
        nickName: "",
        headerImg: "",
        authorityId: ""
      },
      formData: {
        name: null,
        type: null,
        status: true,
        desc: null
      },
      rules: {
        username: [
          {required: true, message: "请输入用户名", trigger: "blur"},
          {min: 6, message: "最低6位字符", trigger: "blur"}
        ],
        password: [
          {required: true, message: "请输入用户密码", trigger: "blur"},
          {min: 6, message: "最低6位字符", trigger: "blur"}
        ],
        nickName: [
          {required: true, message: "请输入用户昵称", trigger: "blur"}
        ],
        authorityId: [
          {required: true, message: "请选择用户角色", trigger: "blur"}
        ]
      }
    };
  },
  methods: {
    openHeaderChange() {
      this.$refs.chooseImg.open()
    },
    setOptions(authData) {
      this.authOptions = [];
      this.setAuthorityOptions(authData, this.authOptions);
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
      AuthorityData.map(item => {
        if (item.children && item.children.length) {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            children: []
          };
          this.setAuthorityOptions(item.children, option.children);
          optionsData.push(option);
        } else {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName
          };
          optionsData.push(option);
        }
      });
    },
    async deleteUser(row) {
      const res = await deleteUser({id: row.ID});
      if (res.code === 0) {
        this.getTableData();
        row.visible = false;
      }
    },
    async enterAddUserDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo);
          if (res.code === 0) {
            this.$message({type: "success", message: "创建成功"});
          }
          await this.getTableData();
          this.closeAddUserDialog();
        }
      });
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields();
      this.addUserDialog = false;
    },
    handleAvatarSuccess(res) {
      this.userInfo.headerImg = res.data.file.url;
    },
    addUser() {
      this.addUserDialog = true;
    },
    async changeAuthority(row) {
      const res = await setUserAuthority({
        uuid: row.uuid,
        authorityId: row.authority.authorityId
      });
      if (res.code === 0) {
        this.$message({type: "success", message: "角色设置成功"});
      }
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      if (this.searchInfo.status === "") {
        this.searchInfo.status = null;
      }
      this.getTableData();
    },

    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        name: null,
        type: null,
        status: true,
        desc: null
      };
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  computed: {
    ...mapGetters("user", ["token"])
  },
  async created() {
    this.getTableData();
    const res = await getAuthorityList({page: 1, pageSize: 999});
    this.setOptions(res.data.list);
  }
};
</script>

<style>
</style>