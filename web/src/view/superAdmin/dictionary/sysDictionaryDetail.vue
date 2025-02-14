<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="展示值">
          <el-input placeholder="搜索条件" v-model="searchInfo.label"></el-input>
        </el-form-item>
        <el-form-item label="字典值">
          <el-input placeholder="搜索条件" v-model="searchInfo.value"></el-input>
        </el-form-item>
        <el-form-item label="启用状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择">
            <el-option key="true" label="是" value="true"></el-option>
            <el-option key="false" label="否" value="false"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增字典项</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="展示值" prop="label" width="120"></el-table-column>

      <el-table-column label="字典值" prop="value" width="120"></el-table-column>

      <el-table-column label="启用状态" prop="status" width="120">
        <template slot-scope="scope">{{scope.row.status|formatBoolean}}</template>
      </el-table-column>

      <el-table-column label="排序标记" prop="sort" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateSysDictionaryDetail(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteSysDictionaryDetail(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
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
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="110px">
        <el-form-item label="展示值" prop="label">
          <el-input
            v-model="formData.label"
            placeholder="请输入展示值"
            clearable
            :style="{width: '100%'}"
          ></el-input>
        </el-form-item>
        <el-form-item label="字典值" prop="value">
          <el-input-number
            v-model.number="formData.value"
            step-strictly
            :step="1"
            placeholder="请输入字典值"
            clearable
            :style="{width: '100%'}"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="启用状态" prop="status" required>
          <el-switch v-model="formData.status" active-text="开启" inactive-text="停用"></el-switch>
        </el-form-item>
        <el-form-item label="排序标记" prop="sort">
          <el-input-number v-model.number="formData.sort" placeholder="排序标记"></el-input-number>
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
  createSysDictionaryDetail,
  deleteSysDictionaryDetail,
  updateSysDictionaryDetail,
  findSysDictionaryDetail,
  getSysDictionaryDetailList
} from "@/api/sysDictionaryDetail";
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "SysDictionaryDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getSysDictionaryDetailList,
      dialogFormVisible: false,
      type: "",
      formData: {
        label: null,
        value: null,
        status: true,
        sort: null
      },
      rules: {
        label: [
          {
            required: true,
            message: "请输入展示值",
            trigger: "blur"
          }
        ],
        value: [
          {
            required: true,
            message: "请输入字典值",
            trigger: "blur"
          }
        ],
        sort: [
          {
            required: true,
            message: "排序标记",
            trigger: "blur"
          }
        ]
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    }
  },
  methods: {

    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      if (this.searchInfo.status == "") {
        this.searchInfo.status = null;
      }
      this.getTableData();
    },
    async updateSysDictionaryDetail(row) {
      const res = await findSysDictionaryDetail({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.resysDictionaryDetail;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        label: null,
        value: null,
        status: true,
        sort: null,
        sysDictionaryID: ""
      };
    },
    async deleteSysDictionaryDetail(row) {
      row.visible = false;
      const res = await deleteSysDictionaryDetail({ ID: row.ID });
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
      this.formData.sysDictionaryID = Number(this.$route.params.id);
      this.$refs["elForm"].validate(async valid => {
        if (!valid) return;
        let res;
        switch (this.type) {
          case "create":
            res = await createSysDictionaryDetail(this.formData);
            break;
          case "update":
            res = await updateSysDictionaryDetail(this.formData);
            break;
          default:
            res = await createSysDictionaryDetail(this.formData);
            break;
        }
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "创建/更改成功"
          });
          this.closeDialog();
          this.getTableData();
        }
      });
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.searchInfo.sysDictionaryID = this.$route.params.id;
    this.getTableData();
  }
};
</script>

<style>
</style>