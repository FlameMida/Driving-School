<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="学员姓名">
          <template>
            <el-cascader
                v-model="searchInfo.userId"
                :options="stuOptions"
                :props="{ checkStrictly: true,label:'stuName',value:'stuId',disabled:'disabled',emitPath:false}"
                :show-all-levels="false"
                filterable
                clearable
            ></el-cascader>
          </template>
        </el-form-item>
        <el-form-item label="考试名">
          <el-input v-model="searchInfo.name" placeholder="搜索条件"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="openDialog">新增考试信息</el-button>
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
        ref="multipleTable"
        :data="tableData"
        border
        stripe
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55"></el-table-column>

      <el-table-column label="学员ID" prop="userId" width="80px"></el-table-column>

      <el-table-column label="学员姓名" width="220px">
        <template slot-scope="scope">
          <el-cascader
              v-model="scope.row.userId"
              :options="stuOptions"
              :props="{ checkStrictly: true,label:'stuName',value:'stuId',disabled:'disabled',emitPath:false}"
              :show-all-levels="false"
              disabled
              filterable
          ></el-cascader>
        </template>
      </el-table-column>

      <el-table-column label="考试名称" prop="name" width="220px"></el-table-column>

      <el-table-column label="考试结果" prop="result" width="220px"></el-table-column>

      <el-table-column label="考试时间" width="220px">
        <template slot-scope="scope">{{ scope.row.time|formatDate }}</template>
      </el-table-column>

      <el-table-column label="备注信息" prop="message" width="220px"></el-table-column>


      <el-table-column label="按钮组" min-width="400px">
        <template slot-scope="scope">
          <el-button class="table-button" icon="el-icon-edit" size="small" type="primary"
                     @click="checkExam(scope.row)">查看学员详情
          </el-button>
          <el-button class="table-button" icon="el-icon-edit" size="small" type="primary"
                     @click="updateExam(scope.row)">修改考试信息
          </el-button>
          <el-button class="table-button" icon="el-icon-delete" size="small" type="danger"
                     @click="deleteExam(scope.row)">删除
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
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :title="title" :visible.sync="dialogFormVisible">
      <el-form ref="userForm" :model="formData" :rules="rules" label-position="right" label-width="130px" size="medium">
        <el-form-item label="学员姓名:" prop="userId">
          <template>
            <el-cascader
                v-model="formData.userId"
                :options="stuOptions"
                :props="{ checkStrictly: true,label:'stuName',value:'stuId',disabled:'disabled',emitPath:false}"
                :show-all-levels="false"
                filterable
            ></el-cascader>
          </template>
        </el-form-item>

        <el-form-item label="考试名称:" prop="name" required>
          <el-input v-model="formData.name" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="考试结果:" prop="result" required>
          <template>
            <el-cascader
                v-model="formData.result"
                :options="resultOptions"
                :props="{ checkStrictly: true,label:'result',value:'result',disabled:'disabled',emitPath:false}"
                :show-all-levels="false"
                filterable
            ></el-cascader>
          </template>
        </el-form-item>

        <el-form-item label="备注:" prop="message">
          <el-input v-model="formData.message" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="考试时间:" prop="time" required>
          <el-date-picker v-model="formData.time" placeholder="选择日期" style="width: 100%;"
                          type="datetime"></el-date-picker>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import infoList from "@/mixins/infoList";
import {createExam, deleteExam, deleteExamByIds, findExam, getExamList, updateExam,} from "@/api/exam";
import {formatTimeToStr} from "@/utils/date";
import {getStudentList} from "@/api/student";

export default {
  name: "exam",
  mixins: [infoList],
  data() {
    return {
      listApi: getExamList,
      dialogFormVisible: false,
      type: "",
      title: "",
      stuOptions: [],
      resultOptions: [],
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        id: "",
        userId: "",
        name: "",
        result: "",
        message: "",
        time: "",
      },
      rules: {
        userId: [
          {required: true, message: '请选择学员', trigger: 'blur'}
        ],
        name: [
          {required: true, message: "请输入考试名", trigger: "blur"},
          {min: 3, message: "最低3位字符", trigger: "blur"}
        ],
        result: [
          {required: true, message: "请选择考试结果", trigger: "blur"},
        ],
        time: [
          {required: true, message: "请选择考试时间", trigger: "blur"}
        ],
      }
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time !== "") {
        const date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
  },
  methods: {
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    setOptions(stuData) {
      this.stuOptions = [];
      this.setStuOptions(stuData, this.stuOptions);

    },
    checkExam(row) {
      this.$router.push({name: 'examDetail', query: {id: row.userId}})
    },
    setStuOptions(StuData, optionsData) {
      StuData &&
      StuData.map(item => {
        const option = {
          stuId: item.ID,
          stuName: item.nickName
        };
        optionsData.push(option);
      });
      const option1 = {
        result: "准备考试"
      };
      const option2 = {
        result: "通过考试"
      };
      const option3 = {
        result: "考试失败"
      };
      const option4 = {
        result: "重考",
        children: [
          {result: "重考1"},
          {result: "重考2"},
        ]
      };
      this.resultOptions.push(option1, option2, option3, option4)
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteExam(row);
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
      const res = await deleteExamByIds({ids})
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        await this.getTableData()
      }
    },
    async updateExam(row) {
      //todo
      const res = await findExam({ID: row.ID});
      this.type = "update";
      this.title = "修改考试情况"
      if (res.code === 0) {
        this.formData = res.data.reExam;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        id: "",
        userId: "",
        name: "",
        result: "",
        message: "",
        time: "",
      };
    },
    async deleteExam(row) {
      const res = await deleteExam({ID: row.ID});
      if (res.code === 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length === 1) {
          this.page--;
        }
        this.deleteVisible2 = false
        await this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          this.$refs.userForm.validate(async valid => {
            if (valid) {
              res = await createExam(this.formData);
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
              res = await updateExam(this.formData);
              if (res.code === 0) {
                this.$message({type: "success", message: "修改成功"});
              }
              this.closeDialog();
              await this.getTableData();
            }
          });
          break;
        default:
          res = await createExam(this.formData);
          break;
      }
    },
    openDialog() {
      this.type = "create";
      this.title = "新增考试情况"
      this.dialogFormVisible = true;
    }
  },

  async beforeMount() {
    if (this.$route.query.open === '1') {
      this.formData.userId = this.$route.query.id
      this.openDialog()
    }
    await this.getTableData();
    const stuRes = await getStudentList({page: 1, pageSize: 999});
    this.setOptions(stuRes.data.list);
  }
};

</script>

<style>
</style>