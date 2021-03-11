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
          <el-button type="primary" @click="openDialog">新增学员</el-button>
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


      <el-table-column label="用户角色ID" prop="authorityId" width="220px"></el-table-column>

      <el-table-column label="用户头像" prop="headerImg" width="220px"></el-table-column>

      <el-table-column label="用户昵称" prop="nickName" width="220px"></el-table-column>

      <el-table-column label="用户登录名" prop="username" width="220px"></el-table-column>

      <el-table-column label="用户UUID" prop="uuid" width="320px"></el-table-column>

      <el-table-column label="入学日期" width="280px">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" icon="el-icon-edit" size="small" type="primary"
                     @click="updateStudent(scope.row)">变更
          </el-button>
          <el-button icon="el-icon-delete" size="mini" type="danger" @click="deleteRow(scope.row)">删除</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="130px" size="medium">
        <el-form-item label="用户角色ID:">
          <el-input v-model="formData.authorityId" clearable placeholder="请输入"></el-input>
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
  createStudent,
  deleteStudent,
  deleteStudentByIds,
  findStudent,
  getStudentList,
  updateStudent
} from "@/api/student"; //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "Student",
  mixins: [infoList],
  data() {
    return {
      listApi: getStudentList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [], formData: {
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
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
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
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.hasServerInfo == "") {
        this.searchInfo.hasServerInfo = null
      }
      this.getTableData()
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
        this.deleteStudent(row);
      });
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length == 0) {
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
      const res = await deleteStudentByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateStudent(row) {
      const res = await findStudent({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.restudent;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        authorityId: "",
        hasServerInfo: false,
        headerImg: "",
        nickName: "",
        password: "",
        username: "",
        uuid: "",

      };
    },
    async deleteStudent(row) {
      const res = await deleteStudent({ID: row.ID});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createStudent(this.formData);
          break;
        case "update":
          res = await updateStudent(this.formData);
          break;
        default:
          res = await createStudent(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();

  }
};
</script>

<style>
</style>
