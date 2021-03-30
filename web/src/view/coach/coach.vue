<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户昵称">
          <el-input v-model="searchInfo.nickName" placeholder="搜索条件"></el-input>
        </el-form-item>
        <el-form-item label="用户登录名">
          <el-input v-model="searchInfo.username" placeholder="搜索条件"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增教练</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
        :data="tableData"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>

      <el-table-column label="用户UUID" prop="uuid" width="320px"></el-table-column>
      <!--      <el-table-column label="用户角色"  width="220px"></el-table-column>-->

      <!--      <el-table-column label="用户头像" prop="headerImg" width="220px"></el-table-column>-->

      <el-table-column label="用户姓名" prop="nickName" width="220px"></el-table-column>

      <el-table-column label="用户登录名" prop="username" width="220px"></el-table-column>

      <el-table-column label="用户手机" prop="phone" width="220px"></el-table-column>

      <el-table-column label="入职日期" width="280px">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>

      <el-table-column label="按钮组" min-width="220px">
        <template slot-scope="scope">
          <el-button class="table-button" icon="el-icon-edit" size="small" type="primary"
                     @click="updateCoach(scope.row)">修改
          </el-button>
          <el-button class="table-button" icon="el-icon-delete" size="small" type="danger"
                     @click="deleteRow(scope.row)">删除
          </el-button>
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
        layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
    <el-dialog :before-close="closeDialog" :title="title" :visible.sync="dialogFormVisible">
      <el-form ref="userForm" :model="formData" :rules="rules" label-position="right" label-width="130px" size="medium">
        <el-form-item label="用户角色" prop="authorityId" required>
          <template v-if="type==='create'">
            <el-cascader
                v-model="formData.authorityId"
                :options="authOptions"
                :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
                :show-all-levels="false"
                disabled
                filterable
            ></el-cascader>
          </template>
          <template v-if="type==='update'">
            <el-cascader
                v-model="formData.authorityId"
                :options="authOptions"
                :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
                :show-all-levels="false"
                filterable
            ></el-cascader>
          </template>
        </el-form-item>

        <el-form-item label="用户账号:" prop="username" required>
          <el-input v-model="formData.username" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="用户姓名:" prop="nickName" required>
          <el-input v-model="formData.nickName" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="用户手机:" prop="phone">
          <el-input v-model="formData.phone" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item v-if="type==='create'" label="用户登录密码:" prop="password" required>
          <el-input v-model="formData.password" clearable placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item v-if="type==='update'" label="用户登录密码:" prop="password">
          <el-input v-model="formData.password" clearable placeholder="请输入"></el-input>
        </el-form-item>

      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createCoach,
  deleteCoach,
  deleteCoachByIds,
  findCoach,
  getCoachList,
  updateCoach,
  updateCoachPWD
} from "@/api/coach";
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {getAuthorityList} from "@/api/authority";


export default {
  name: "Coach",
  mixins: [infoList],
  data() {
    return {
      listApi: getCoachList,
      title: "",
      dialogFormVisible: false,
      type: "",
      authOptions: [],
      deleteVisible: false,

      multipleSelection: [],
      formData: {
        authorityId: "",
        hasServerInfo: false,
        headerImg: "",
        nickName: "",
        password: "",
        username: "",
        uuid: "",
        phone: "",
      },
      rules: {
        phone: [
          {min: 11, max: 11, message: '长度为11位', trigger: 'blur'}
        ],
        username: [
          {required: true, message: "请输入用户名", trigger: "blur"},
          {min: 6, message: "最低6位字符", trigger: "blur"}
        ],
        password: [
          {message: "请输入用户密码", trigger: "blur"},
          {min: 6, message: "最低6位字符", trigger: "blur"}
        ],
        nickName: [
          {required: true, message: "请输入用户姓名", trigger: "blur"}
        ],
        authorityId: [
          {required: true, message: "请选择用户角色", trigger: "blur"}
        ]
      }
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time !== "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    }
  },
  methods: {

    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.hasServerInfo === "") {
        this.searchInfo.hasServerInfo = null
      }
      this.getTableData()
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
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteCoach(row);
      });
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteCoachByIds({ids})
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateCoach(row) {
      const res = await findCoach({ID: row.ID});
      this.type = "update";
      this.title = "修改教练"
      if (res.code === 0) {
        this.formData = res.data.recoach;
        this.dialogFormVisible = true;
      }
    },

    closeDialog() {
      this.dialogFormVisible = false;
      this.type = ""
      this.formData = {
        authorityId: "",
        hasServerInfo: false,
        headerImg: "",
        nickName: "",
        username: "",
        uuid: "",

      };
    },
    async deleteCoach(row) {
      const res = await deleteCoach({ID: row.ID});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        await this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          this.$refs.userForm.validate(async valid => {
            if (valid) {
              res = await createCoach(this.formData);
              if (res.code === 0) {
                this.$message({type: "success", message: "创建成功"});

              }
              this.closeDialog();
              await this.getTableData();
            }
          });

          break;
        case "update":
          this.$refs.userForm.validate(async valid => {
            if (valid) {
              if (this.formData.password) {
                res = await updateCoachPWD(this.formData)
              }
              res = await updateCoach(this.formData);
              if (res.code === 0) {
                this.$message({type: "success", message: "修改成功"});
              }
              this.closeDialog();
              await this.getTableData();
            }
          });
          break;
        default:
          res = await createCoach(this.formData);
          break;
      }

    },
    openDialog() {
      this.type = "create";
      this.formData.authorityId = "100"
      this.title = "新增教练"
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
    const res = await getAuthorityList({page: 1, pageSize: 999});
    this.setOptions(res.data.list);
  }
};
</script>

<style>
</style>
