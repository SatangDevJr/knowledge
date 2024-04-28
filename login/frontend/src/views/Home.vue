<template>
  <div class="card-nav">
    <div class="card-contenthalf">
      <div class="card-oversize">
        <div class="container-fluid page-body-wrapper">
          <div class="container-scroller">
            <div class="main-panel">
              <div class="content-wrapper pb-0">
                <div class="row">
                  <div class="col-xl-12 stretch-card grid-margin justify-content-center text-welcome">
                    <h2>
                     Welcome User: {{ userCode }}
                    </h2>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref,onMounted } from "vue";
import { useAccountStore } from "../stores/account.store";
const accountStore = useAccountStore();
const userCode = ref("")

function onloadMenu() {
  let loginData = accountStore.getLoginData;
  console.log("loginData :",loginData)
  if (loginData != undefined){
    let nameDisplay = loginData.User.firstnameth + " "+ loginData.User.lastnameth;
    userCode.value = loginData.User.username;
  }
}


onMounted(() => {
  onloadMenu();
});

</script>

<style>
.text-welcome{
  margin-top: 36vh;
}
</style>