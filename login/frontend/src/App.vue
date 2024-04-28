<template class="scrollable-container">
  <Header class="header" v-if="!$route.meta.guest"></Header>
  <Content></Content>
  <Footer></Footer>

</template>
<script setup>
import Header from "./layouts/Header.vue";
import Content from "./layouts/Content.vue";
import Footer from "./layouts/Footer.vue";
import { onMounted } from "vue";

// lifecycle hooks
onMounted(() => {

});
</script>
<style>

html {
  height: 100%;
  width: 100%;
  position: relative;
}

body {
  padding-bottom: 256px;
  width: 100%;
  min-width: 1024px;
  min-height: 100%;
  position: relative;
  margin: 0px auto;
}

.header {
  top: 0;
}

.ant-form-item {
  margin-bottom: 6px !important;
}

input[type="number"] {
  text-align: right !important;
}

.spantext {
  width: 90%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  display: inline-block;
  padding-top: 6px !important;
}

.ant-table-tbody > tr > td,
.ant-table.ant-table-small tfoot > tr > th,
.ant-table.ant-table-small tfoot > tr > td {
  padding: 0px 0px !important;
}

.ant-empty {
  height: calc(100vh - 420px) !important;
  overflow: hidden !important;
}

.ant-table-container {
  height: calc(100vh - 280px) !important;
}
.ant-table-footer {
  background-color: lightgray !important;
}

</style>
