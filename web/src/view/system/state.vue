<template>
  <div>
    <el-row :gutter="15" class="system_state">
      <el-col :span="24">
        <el-card v-if="state.os" class="card_item">
          <div slot="header">服务器环境：</div>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">操作系统:</el-col>
              <el-col :span="12" v-text="state.os.goos"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">CPU 数量:</el-col>
              <el-col :span="12" v-text="state.os.numCpu"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">编译器:</el-col>
              <el-col :span="12" v-text="state.os.compiler"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">Go 版本:</el-col>
              <el-col :span="12" v-text="state.os.goVersion"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">Go协程数量 :</el-col>
              <el-col :span="12" v-text="state.os.numGoroutine"></el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>

    </el-row>
    <el-row :gutter="15" class="system_state">
      <el-col :span="24">
        <el-card
            class="CPU_card_item"
            v-if="state.cpu"
            :body-style="{ height: '400px', 'overflow-y': 'scroll' }">
          <div slot="header">CPU：</div>
          <div>
            <el-row :gutter="10" style="margin-bottom: 24px">
              <el-col :span="20">
                物理核心数量 ：
              </el-col>
              <el-col :span="20">
                <el-slider
                    style="text-align: center;margin: 0 auto"
                    :value="state.cpu.cores"
                    show-stops
                    :max="state.cpu.cores?state.cpu.cores:16"
                    :step="2"
                    :marks="marks"
                    disabled>
                </el-slider>
              </el-col>
            </el-row>
            <el-row  :gutter="10">
              <el-col :span="24" style="padding-bottom: 16px">
                核心使用情况 ：
              </el-col>
              <el-col :span="4" :body-style="{ padding: '0px' }" v-for="(item, index) in state.cpu.cpus" :key="index">
                <el-card shadow="hover">
                  <div style="text-align: center">
                    <el-progress type="dashboard" :percentage="+item.toFixed(0)" :color="colors">
                    </el-progress>
                  </div>
                  <div style="padding: 14px;text-align: center">
                    <span>Core {{ index }}</span>
                  </div>
                </el-card>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="15" class="system_state">
      <el-col :span="12">
        <el-card class="card_item" v-if="state.ram">
          <div slot="header">RAM：</div>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">
                <el-row :gutter="10">
                  <el-col :span="12">总共容量 (MB)</el-col>
                  <el-col :span="12" v-text="state.ram.totalMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用容量 (MB)</el-col>
                  <el-col :span="12" v-text="state.ram.usedMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">总共容量 (GB)</el-col>
                  <el-col :span="12" v-text="state.ram.totalMb / 1024"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用容量 (GB)</el-col>
                  <el-col
                      :span="12"
                      v-text="(state.ram.usedMb / 1024).toFixed(2)"
                  ></el-col>
                </el-row>
              </el-col>
              <el-col :span="12">
                <el-progress
                    type="dashboard"
                    :percentage="state.ram.usedPercent"
                    :color="colors"
                ></el-progress>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card v-if="state.disk" class="card_item">
          <div slot="header">磁盘情况：</div>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">
                <el-row :gutter="10">
                  <el-col :span="12">总共容量 (MB)</el-col>
                  <el-col :span="12" v-text="state.disk.totalMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用容量 (MB)</el-col>
                  <el-col :span="12" v-text="state.disk.usedMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">总共容量 (GB)</el-col>
                  <el-col :span="12" v-text="state.disk.totalGb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用容量 (GB)</el-col>
                  <el-col :span="12" v-text="state.disk.usedGb"></el-col>
                </el-row>
              </el-col>
              <el-col :span="12">
                <el-progress
                    type="dashboard"
                    :percentage="state.disk.usedPercent"
                    :color="colors"
                ></el-progress>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {getSystemState} from "@/api/system.js";

export default {
  name: "State",
  data() {
    return {
      marks: {
        1: '1',
        2: '2',
        4: '4',
        6: '6',
        8: '8',
        10: '10',
        12: '12',
        14: '14',
        16: '16'
      },
      timer: null,
      state: {},
      colors: [
        {color: "#5cb87a", percentage: 20},
        {color: "#e6a23c", percentage: 40},
        {color: "#f56c6c", percentage: 80},
      ],
    };
  },
  created() {
    this.reload();
    this.timer = setInterval(() => {
      this.reload();
    }, 1000 * 10);
  },
  beforeDestroy() {
    clearInterval(this.timer)
    this.timer = null
  },
  methods: {
    async reload() {
      const {data} = await getSystemState();
      this.state = data.server;
    },
  },
};
</script>

<style>
.system_state {
  padding: 10px;
}

.card_item {
  height: 280px;
}

.CPU_card_item {
  height: 480px;
}

</style>
