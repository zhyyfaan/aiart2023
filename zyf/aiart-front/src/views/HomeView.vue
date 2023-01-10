<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <div class="header">
      <img
        src="/imgs/header.png"
        alt=""
        style="object-fit: contain; height: 250px"
      />
    </div>
    <div class="divider">
      <el-divider><i class="el-icon-caret-bottom"></i></el-divider>
    </div>

    <div class="main">
      <div class="uploader">
        <el-upload
          class="upload"
          ref="upload"
          drag
          action="string"
          :file-list="fileList"
          :auto-upload="false"
          :http-request="uploadFile"
          :on-change="handleChange"
          :on-preview="handlePreview"
          :on-remove="handleRemove"
          multiple="multiple"
        >
          <i class="el-icon-upload"></i>
          <div class="el-upload__text">
            将文件拖到此处，或 <em>点击上传</em>
          </div>
          <div class="el-upload__tip" slot="tip">
            只能上传jpg/png文件，且不超过500kb
          </div>
          <!-- <el-button slot="trigger"
                   size="small"
                   type="primary"
                   @click="delFile">选取文件
          </el-button> -->
        </el-upload>
        <div class="uploadButton">
          <el-button type="primary" @click="onSubmit">保存</el-button>
        </div>
      </div>
      <div class="generator">
        <div class="photo">
          <div class="demo-image">
            <div class="block">
              <el-image
                style="width: 250px; height: 200px"
                :src="photoUrl"
                :fit="fit"
              ></el-image>
            </div>
            <div class="tag" slot="tip">
              照片
            </div>
          </div>
        </div>
        <div class="sketch">         
            <div class="demo-image">
              <div class="block">
                <el-image
                  style="width: 250px; height: 200px"
                  :src="sketchUrl"
                  :fit="fit"
                ></el-image>
              </div>
            </div>
            <div class="tag" slot="tip">
              素描
            </div>
        </div>
      </div>
    </div>
    <div class="footer"></div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: "HomeView",
  // components: {
  //   HelloWorld
  // }
  data() {
    return {
      // el-upload组件绑定的属性—— :file-list=“fileList”,渲染后fileList[index]是Object类型,而不是后台所需的File类型,
      // 而这个组件已经把对应的File类型存储到了fileList[index].raw这个属性里,直接拿来用就好.
      fileList: [],
      // 不支持多选
      multiple: false,
      formData: "",
      fit: "contain",
      photoUrl: "",
      sketchUrl: "",
    };
  },
  methods: {
    //点击上传文件触发的额外事件,清空fileList
    delFile() {
      this.fileList = [];
    },
    handleChange(file, fileList) {
      this.fileList = fileList;
      // console.log(this.fileList, "sb");
    },
    //自定义上传文件
    uploadFile(file) {
      this.formData.append("file", file.file);
      // console.log(file.file, "sb2");
    },
    //删除文件
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    // 点击文件
    handlePreview(file) {
      console.log(file);
    },
    //保存按钮
    onSubmit() {
      let formData = new FormData();
      formData.append("file", this.fileList[0].raw); //拿到存在fileList的文件存放到formData中
      //下面数据是我自己设置的数据,可自行添加数据到formData(使用键值对方式存储)
      formData.append("title", this.title);
      this.axios
        .post(post请求的具体路径, formData, {
          "Content-Type": "multipart/form-data;charset=utf-8",
        })
        .then((res) => {
          if (res.data === "SUCCESS") {
            this.$notify({
              title: "成功",
              message: "提交成功",
              type: "success",
              duration: 1000,
            });
            let blob = new Blob([res.data]); // 返回的文件流数据
            this.sketchUrl = window.URL.createObjectURL(blob); // 将他转化为路径
          }
        });
    },
  },
};
</script>
<style lang="scss">
.header {
  // background-color: grey;
  height: 250px;
  width: 100%;
}
.divider {
  margin-top: 50px;
}
.main {
  // background-color: pink;
  height: 500px;
  width: 100%;
  position: relative;
  .uploader {
    padding-top: 70px;
    padding-left: 200px;
    float: left;
  }
  .uploadButton {
    margin-top: 30px;
    padding-left: 140px;
    float: left;
  }
  .generator {
    padding-top: 70px;
    .photo {
      display: inline-block;
    }
    .sketch {
      display: inline-block;
      margin-left: 20px;
    }
    .tag{
      margin-top:20px;
    }
  }
}
.footer {
  // background-color: grey;
  height: 100px;
  width: 100%;
}
</style>