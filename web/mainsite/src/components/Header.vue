<script setup>
import { Search, Iphone } from '@element-plus/icons-vue'
function upload() {

}
</script>

<template>
  <el-row class="nav-header-right">
    <span>
      <strong
        class="nav-link-item"
        href="/"
        :underline="false"
        style="color: #91bef0; font-size: larger;"
      >Raiden Blade!</strong>
      <el-link class="nav-link-item" href="/" :underline="false">番剧</el-link>
      <el-link class="nav-link-item" href="/" :underline="false">游戏中心</el-link>
      <el-link class="nav-link-item" href="/" :underline="false">直播</el-link>
      <el-link class="nav-link-item" href="/" :underline="false">会员购</el-link>
      <el-link class="nav-link-item" href="/" :underline="false">漫画</el-link>
      <el-link class="nav-link-item" href="/" :underline="false">赛事</el-link>
      <el-link class="nav-link-item" href="/" :underline="false">下载APP</el-link>
    </span>
    <span>
      <el-input name="search" size="large" placeholder="此刻！寂灭之时" style="width: 100%;">
        <template #append>
          <el-button :icon="Search" />
        </template>
      </el-input>
    </span>
    <span>
      <el-row class="nav-header-right-row">
        <span>
          <el-avatar class="nav-link-item" src="src/assets/avatar.jpg"></el-avatar>
        </span>
        <span>
          <el-link class="nav-link-item" href="/" :underline="false">大会员</el-link>
          <el-link class="nav-link-item" href="/" :underline="false">消息</el-link>
          <el-link class="nav-link-item" href="/" :underline="false">动态</el-link>
          <el-link class="nav-link-item" href="/" :underline="false">收藏</el-link>
          <el-link class="nav-link-item" href="/" :underline="false">历史</el-link>
          <el-link class="nav-link-item" href="/" :underline="false">创作中心</el-link>
          <el-button size="large" color="#eb96ab" style="color: white" @click="upload()">
            <strong>&nbsp;&nbsp;&nbsp;&nbsp;投稿&nbsp;&nbsp;&nbsp;&nbsp;</strong>
          </el-button>
        </span>
      </el-row>
    </span>
  </el-row>
</template>

<style>
.nav-header-right {
  width: 100%;
  height: 56px;
  padding-left: 20px;
  padding-right: 20px;
  box-shadow: 0 2px 4px 0 rgb(0 0 0 / 8%);
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: white;
}

.nav-header-right-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.nav-link-item {
  margin-right: 12px;
}
</style>