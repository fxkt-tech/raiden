<script setup>
import { ref } from 'vue'
import Header from "./Header.vue";

defineProps({
  msg: String
})

const count = ref(0)
</script>

<template>
  <div class="root-background">
    <el-container>
      <el-header class="frame-header-external">
        <Header></Header>
      </el-header>
      <el-main class="frame-main-external">
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>
.frame-header-external {
  padding-left: 0;
  padding-right: 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.frame-main-external {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}
</style>
