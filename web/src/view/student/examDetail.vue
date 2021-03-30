<template>
  <div>
    <div class="block">
      <div class="radio">
        排序：
        <el-radio-group v-model="reverse">
          <el-radio :label="true">倒序</el-radio>
          <el-radio :label="false">正序</el-radio>
        </el-radio-group>
      </div>
      <el-divider>{{ title }}</el-divider>
      <el-timeline :reverse="reverse">
        <el-timeline-item
            v-for="(activity, index) in activities"
            :key="index"
            :timestamp="activity.time|formatDate"
            icon="el-icon-more"
            placement="top"
            size="large"
            type="primary">
          <el-card class="text item" shadow="hover">
            <div class="item"><h2>{{ activity.name }}</h2></div>
            <div class="item"><p>于{{ activity.time|formatDate }} - {{ activity.result }}</p></div>
            <div class="item"><p>{{ activity.message }}</p></div>
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </div>
  </div>
</template>

<script>
import infoList from "@/mixins/infoList";
import {getExamDetailList} from "@/api/exam";
import {formatTimeToStr} from "@/utils/date";
import {getStudentList} from "@/api/student";
import {mapGetters} from "vuex";

export default {
  name: "examDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getExamDetailList,
      title: "",
      reverse: true,
      stuOptions: [],
      activities: [],
    };
  }, computed: {
    ...mapGetters("user", ["userInfo"]),
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
    setOptions(stuData) {
      this.stuOptions = [];
      this.setStuOptions(stuData, this.stuOptions);

    },
    setStuOptions(StuData, optionsData) {
      StuData &&
      StuData.map(item => {
        if (item.ID === this.activities[0].userId) {
          this.title = item.ID + "号学员-" + item.nickName + "的考试情况"
        }
        const option = {
          stuId: item.ID,
          stuName: item.nickName
        };
        optionsData.push(option);
      });
    },
  },
  async beforeMount() {
    let userId = this.$route.query.id
    if (this.userInfo.authority.authorityId === '200' || this.userInfo.authority.parentId === '200') {
      userId = this.userInfo.ID
    }
    if (userId <= 0 || !userId) {
      this.$message.error({
        message: '数据不存在哦', onClose: function () {
          this.$router.back()
        }
      });
      // await this.$router.push({name: "404"})
    }
    const ExamRes = await getExamDetailList({userId});
    const stuRes = await getStudentList({page: 1, pageSize: 999});
    this.activities = ExamRes.data.list
    this.setOptions(stuRes.data.list);

  }
};

</script>

<style>
.block {
  margin: 10px auto;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
</style>