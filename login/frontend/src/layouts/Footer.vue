<template>
  <div id="footermain" class="footer" >
    <span>{{ footerName }}</span>
  </div>
</template>
<script setup>
import { useI18n } from "vue-i18n";
import {  ref } from "vue";

const { locale, t } = useI18n();
const footerName = ref('Copyright © 2022 Thai Beverage Plc. All Rights Reserved. (1.0.0.0)');

</script>
<style scoped>
.footer {
  background: #759801;
  color: white;
  font-size: 12px;
  width: 100%;
  position: fixed;
  bottom: 0;
}
</style>