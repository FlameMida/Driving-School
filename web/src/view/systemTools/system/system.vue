<template>
  <div class="system">
    <el-form ref="form" :model="config" class="system" label-width="200px">
      <!--  System start  -->
      <!--      TODO：music & weather config-->
      <h2>系统配置</h2>
      <el-form-item label="服务端口">
        <el-input v-model.number="config.system.addr"></el-input>
      </el-form-item>
      <el-form-item label="数据库类型">
        <el-select v-model="config.system.dbType">
          <el-option value="mysql"></el-option>
          <el-option value="sqlite"></el-option>
          <el-option value="sqlserver"></el-option>
          <el-option value="postgresql"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Oss类型">
        <el-select v-model="config.system.ossType">
          <el-option value="local"></el-option>

        </el-select>
      </el-form-item>
      <el-form-item label="系统环境">
        <el-input v-model.number="config.system.env"></el-input>
      </el-form-item>
      <el-form-item label="数据初始化">
        <el-checkbox v-model="config.system.needInitData">开启</el-checkbox>
      </el-form-item>
      <el-form-item label="多点登录拦截">
        <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
      </el-form-item>
      <!--  System end  -->

      <!--  JWT start  -->
      <h2>jwt签名</h2>
      <el-form-item label="jwt签名">
        <el-input v-model="config.jwt.signingKey"></el-input>
      </el-form-item>
      <!--  JWT end  -->

      <!--  Zap start  -->
      <h2>Zap日志配置</h2>
      <el-form-item label="级别">
        <el-input v-model.number="config.zap.level"></el-input>
      </el-form-item>
      <el-form-item label="输出">
        <el-input v-model="config.zap.format"></el-input>
      </el-form-item>
      <el-form-item label="日志前缀">
        <el-input v-model="config.zap.prefix"></el-input>
      </el-form-item>
      <el-form-item label="日志文件夹">
        <el-input v-model="config.zap.director"></el-input>
      </el-form-item>
      <el-form-item label="软链接名称">
        <el-input v-model="config.zap.linkName"></el-input>
      </el-form-item>
      <el-form-item label="编码类型">
        <el-input v-model="config.zap.encodeLevel"></el-input>
      </el-form-item>
      <el-form-item label="追踪栈名">
        <el-input v-model="config.zap.stacktraceKey"></el-input>
      </el-form-item>
      <el-form-item label="显示行">
        <el-checkbox v-model="config.zap.showLine"></el-checkbox>
      </el-form-item>
      <el-form-item label="输出控制台">
        <el-checkbox v-model="config.zap.logInConsole"></el-checkbox>
      </el-form-item>
      <!--  Zap end  -->

      <!--  Redis start  -->
      <h2>Redis admin数据库配置</h2>
      <el-form-item label="db">
        <el-input v-model="config.redis.db"></el-input>
      </el-form-item>
      <el-form-item label="地址">
        <el-input v-model="config.redis.addr"></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="config.redis.password"></el-input>
      </el-form-item>
      <!--  Redis end  -->

      <!--  Email start  -->
      <h2>邮箱配置</h2>
      <el-form-item label="接收者邮箱">
        <el-input v-model="config.email.to" placeholder="可多个，以逗号分隔"></el-input>
      </el-form-item>
      <el-form-item label="端口">
        <el-input v-model.number="config.email.port"></el-input>
      </el-form-item>
      <el-form-item label="发送者邮箱">
        <el-input v-model="config.email.from"></el-input>
      </el-form-item>
      <el-form-item label="host">
        <el-input v-model="config.email.host"></el-input>
      </el-form-item>
      <el-form-item label="是否为ssl">
        <el-checkbox v-model="config.email.isSSL"></el-checkbox>
      </el-form-item>
      <el-form-item label="secret">
        <el-input v-model="config.email.secret"></el-input>
      </el-form-item>
      <el-form-item label="测试邮件">
        <el-button @click="email">测试邮件</el-button>
      </el-form-item>
      <!--  Email end  -->

      <!--  Casbin start  -->
      <h2>casbin配置</h2>
      <el-form-item label="模型地址">
        <el-input v-model="config.casbin.modelPath"></el-input>
      </el-form-item>
      <!--  Casbin end  -->

      <!--  Captcha start  -->
      <h2>验证码配置</h2>
      <el-form-item label="验证码长度">
        <el-input v-model.number="config.captcha.keyLong"></el-input>
      </el-form-item>
      <el-form-item label="验证码图片宽度">
        <el-input v-model.number="config.captcha.imgWidth"></el-input>
      </el-form-item>
      <el-form-item label="验证码图片高度">
        <el-input v-model.number="config.captcha.imgHeight"></el-input>
      </el-form-item>
      <!--  Captcha end  -->

      <!--  dbType start  -->
      <template v-if="config.system.dbType == 'mysql'">
        <h2>mysql admin数据库配置</h2>
        <el-form-item label="username">
          <el-input v-model="config.mysql.username"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.mysql.password"></el-input>
        </el-form-item>
        <el-form-item label="DSN">
          <el-input v-model="config.mysql.path"></el-input>
        </el-form-item>
        <el-form-item label="数据库名">
          <el-input v-model="config.mysql.dbname"></el-input>
        </el-form-item>
        <el-form-item label="最大闲置连接数">
          <el-input v-model.number="config.mysql.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="最大连接数">
          <el-input v-model.number="config.mysql.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="日志模式">
          <el-checkbox v-model="config.mysql.logMode"></el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.dbType == 'sqlite'">
        <h2>sqlite admin数据库配置</h2>
        <el-form-item label="path">
          <el-input v-model="config.mysql.path"></el-input>
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.mysql.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.mysql.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.mysql.logger"></el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.dbType == 'sqlserver'">
        <h2>sqlserver admin数据库配置</h2>
        <el-form-item label="username">
          <el-input v-model="config.sqlserver.username"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.sqlserver.password"></el-input>
        </el-form-item>
        <el-form-item label="path">
          <el-input v-model="config.sqlserver.path"></el-input>
        </el-form-item>
        <el-form-item label="dbname">
          <el-input v-model="config.sqlserver.dbname"></el-input>
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.sqlserver.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.sqlserver.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.sqlserver.logger"></el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.dbType == 'postgresql'">
        <h2>postgresql admin数据库配置</h2>
        <el-form-item label="username">
          <el-input v-model="config.mysql.username"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.mysql.password"></el-input>
        </el-form-item>
        <el-form-item label="dbName">
          <el-input v-model="config.mysql.dbName"></el-input>
        </el-form-item>
        <el-form-item label="port">
          <el-input v-model="config.mysql.port"></el-input>
        </el-form-item>
        <el-form-item label="config">
          <el-input v-model="config.mysql.config"></el-input>
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.mysql.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.mysql.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.mysql.logger"></el-checkbox>
        </el-form-item>
        <el-form-item label="prefer-simple-protocol">
          <el-checkbox v-model="config.mysql.preferSimpleProtocol"></el-checkbox>
        </el-form-item>
      </template>
      <!--  dbType end  -->

      <!--  ossType start  -->
      <template v-if="config.system.ossType == 'local'">
        <h2>本地上传配置</h2>
        <el-form-item label="本地文件路径">
          <el-input v-model="config.local.path"></el-input>
        </el-form-item>
      </template>

      <!--  ossType end  -->

      <el-form-item>
        <el-button @click="update" type="primary">立即更新</el-button>
        <el-button @click="reload" type="primary">重启服务（开发中）</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from "@/api/system";
import { emailTest } from "@/api/email";
export default {
  name: "Config",
  data() {
    return {
      config: {
        system: {},
        jwt: {},
        casbin: {},
        mysql: {},
        sqlite: {},
        redis: {},
        captcha: {},
        zap: {},
        local: {},
        email: {}
      }
    };
  },
  async created() {
    await this.initForm();
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig();
      if (res.code == 0) {
        this.config = res.data.config;
      }
    },
    reload() {},
    async update() {
      const res = await setSystemConfig({ config: this.config });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "配置文件设置成功"
        });
        await this.initForm();
      }
    },
    async email() {
      const res = await emailTest();
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "邮件发送成功"
        });
        await this.initForm();
      } else {
        this.$message({
          type: "error",
          message: "邮件发送失败"
        });
      }
    }
  }
};
</script>
<style lang="scss">
.system {
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 1px 3px 0px #e7e8e8;
  }
}
</style>
