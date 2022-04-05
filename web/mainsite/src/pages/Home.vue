<script setup>
import { Search } from '@element-plus/icons-vue'
const handleSelect = (key, keyPath) => {
  console.log(key, keyPath)
}
</script>

<template>
  <div class="home-background">
    <el-row>
      <div class="space-theme space-user-info">
        <el-row style="margin-left: 15px; margin-bottom: 10px">
          <span style="width: 70px">
            <el-avatar size="large" class="nav-link-item" src="src/assets/avatar.jpg"></el-avatar>
          </span>
          <span>
            <el-row>
              <span style="margin-right: 5px">
                <strong>晴川丽丽二叉树</strong>
              </span>
              <el-tag
                size="small"
                color="orange"
                style="color: white;background-color: #ee672a;margin-right: 5px"
              >lv5</el-tag>
              <el-tag
                size="small"
                color="pink"
                style="color: white;background-color: rgb(251, 114, 153)"
              >年度大会员</el-tag>
            </el-row>
            <el-row>
              <span style="font-size: small">一半烟火，一半初心</span>
            </el-row>
          </span>
        </el-row>
      </div>
    </el-row>
    <el-row class="space-nav">
      <span>
        <el-menu default-active="1" mode="horizontal" :ellipsis="false" @select="handleSelect">
          <el-menu-item index="1">主页</el-menu-item>
          <el-menu-item index="2">动态</el-menu-item>
          <el-menu-item index="3">投稿</el-menu-item>
          <el-menu-item index="4">合集和列表</el-menu-item>
          <el-menu-item index="5">收藏</el-menu-item>
          <el-menu-item index="6">订阅</el-menu-item>
          <el-menu-item index="7">设置</el-menu-item>
        </el-menu>
      </span>
      <span>
        <el-input name="space-search" placeholder="搜索视频、动态">
          <template #append>
            <el-button :icon="Search" />
          </template>
        </el-input>
      </span>
      <span>
        <el-link class="nav-link-item" href="/" :underline="false">
          关注数
          <br />1
        </el-link>
        <el-link class="nav-link-item" href="/" :underline="false">
          粉丝数
          <br />2
        </el-link>
        <el-link class="nav-link-item" href="/" :underline="false">
          获赞数
          <br />10
        </el-link>
        <el-link class="nav-link-item" href="/" :underline="false">
          播放数
          <br />15
        </el-link>
        <el-link class="nav-link-item" href="/" :underline="false">
          阅读数
          <br />8
        </el-link>
      </span>
    </el-row>
  </div>
</template>

<style>
.home-background {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.space-theme {
  width: 1283px;
  height: 192px;
  background-color: #beaace;
  background-image: url(src/assets/space.jpg);
  background-size: cover;
}

.space-nav {
  width: 1283px;
  background-color: white;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.space-user-info {
  display: flex;
  flex-direction: column-reverse;
  align-items: left;
  color: white;
}

.el-menu--horizontal {
  border-bottom: solid 1px white;
}
</style>