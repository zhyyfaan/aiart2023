<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <div class="header">
      <div class="title">
        <p class="big">Facial Sketch Synthesis</p>
        <p class="small">
          Generate a quality and vivid sketch portrait from a given photo
        </p>
      </div>
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
      </div>
      <div class="uploadButton">
        <el-button type="primary" :loading="loading" @click="onSubmit">生成素描</el-button>
      </div>
      <div class="generator">
        <div class="photo">
          <div class="demo-image">
            <div class="block">
              <el-image
                style="width: 200px; height: 250px"
                :src="photoUrl"
                :fit="fit"
              >
              </el-image>
            </div>
            <div class="tag" slot="tip">照片</div>
          </div>
        </div>
        <div class="sketch">
          <div class="demo-image">
            <div class="block">
              <el-image
                style="width: 200px; height: 250px"
                :src="sketchUrl"
                :fit="fit"
              ></el-image>
            </div>
          </div>
          <div class="tag" slot="tip">素描</div>
        </div>
      </div>
    </div>

    <div class="footer">
      <p>The Website Design and FSS algorithm are supported by the group 'AiArt' from HDU </p>
      <p>If you want to contact us, please send Email to 2961695289@qq.com</p>
    </div>

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
      loading: false,
    };
  },
  methods: {
    //点击上传文件触发的额外事件,清空fileList
    delFile() {
      this.fileList = [];
    },
    handleChange(file, fileList) {
      // console.log(file)
      // this.photoUrl = file
      this.fileList = fileList;
      // console.log(this.fileList, "sb");
    },
    //自定义上传文件
    uploadFile(file) {
      this.formData.append("file", file.file);
      // console.log(file)
      // this.photoUrl = file
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
      this.loading = true;
      console.log(this.loading);
      let formData = new FormData();
      // formData.append("sender", "ZYF"); //发起者
      formData.append("algorithm_name", "PGP"); //算法
      // formData.append("store_loc", ""); //临时位置
      formData.append("file_name", this.fileList[0].raw.name); //文件名
      formData.append("image", this.fileList[0].raw); //照片
      // console.log(formData.get("algorithm"))
      this.axios
        .post("", formData, {
          "Content-Type": "multipart/form-data;charset=utf-8",
          // "responseType": "blob",
          "responseType": "json",
        })
        .then((res) => {
          if (res.status == 200) {
            this.$notify({
              title: "上传成功",
              message: "正在生成",
              type: "success",
              duration: 1200,
            });
            console.log("res:", res);
            // let blob = new Blob([res.data]); // 返回的文件流数据
            // let photoBlob = blob.slice(0, blob.size/2);//切割，由于两张图片大小相同，因此对半切割
            // let sketchBlob = blob.slice(blob.size/2,blob.size);
            // console.log(photoBlob.size);
            // console.log(sketchBlob.size);
            // this.photoUrl = window.URL.createObjectURL(photoBlob); // 将他转化为路径
            // this.sketchUrl = window.URL.createObjectURL(sketchBlob); // 将他转化为路径

            this.photoUrl = "data:image/jpeg;base64,"+res.data.source_img_base64;
            this.sketchUrl = "data:image/jpeg;base64,"+res.data.result_img_base64;

            console.log("this.photoUrl:", this.photoUrl);
            console.log("this.sketchUrl:", this.sketchUrl);
            // this.sketchUrl = "/imgs/header.png";
          }
        }).finally(()=>{
          this.loading = false;
        })
    },
  },
};
</script>
<style lang="scss">
.header {
  // background-color: grey;
  height: 250px;
  width: 100%;
  .title {
    position: absolute;
    text-align: left;
    left: 220px;
    top: 90px;
    p {
      font-size: 40px;
      font-family: "Helvetica Neue";
    }
    .big {
      margin-bottom: 20px;
      font-weight: 800;
    }
    .small {
      font-size: 20px;
      margin: 0;
      color: gray;
    }
  }
  img {
    display: block;
    margin-left: 60%;
  }
}
.divider {
  margin: 50px 160px 10px 180px;
}
.main {
  // background-color: pink;
  height: 360px;
  width: 100%;
  position: relative;
  .uploader {
    padding-top: 70px;
    padding-left: 200px;
    float: left;
  }
  .uploadButton {
    margin-left: 100px;
    margin-top: 149px;
    // padding-left: 135px;
    float: left;
  }
  .generator {
    display: inline-block;
    padding-top: 40px;
    margin-right: 70px;
    .photo {
      display: inline-block;
    }
    .sketch {
      display: inline-block;
      margin-left: 20px;
    }
    .tag {
      margin-top: 20px;
    }
  }
}
.footer {
  background-color: #f5f7fa;
  height: 100px;
  width: 100%;
  padding-top:18px;
  p{
    font-size: 13px;
    color:#a4b2c7;
  }
}
</style>