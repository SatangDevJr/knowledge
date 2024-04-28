<script setup>
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import { reactive } from "vue";
import { router } from "../router/index";
import { useI18n } from "vue-i18n";
import { number } from '@intlify/core-base';
import { loginUser } from '../utils/auth'
// import firebase from "firebase";

const { locale } = useI18n();

function switchLocale(lang) {
  locale.value = lang;
}

const formState = reactive({
  username: "",
  password: "",
});

async function login(username, password) {
  try {
    await loginUser(username, password);
    router.push("/");
  } catch (err) {
    alert(`Error: ${err}`)
  }
}

const handlerLogin = () => {
  if (formState.username && formState.password) {
    login(formState.username, formState.password, formState.stafflogin)
  }
};

const register = () => {
  router.push("/register");
};

// async function loginFirebase(username, password) {
//   firebase.auth().signInWithEmailAndPassword(username, password).then(
//     user => {
//       router.push("/");
//     },
//     err => {
//       alert(err.message)
//     });
//   e.preventDefault();
// }

// const loginWithThirdParty = () => {
//   var provider = new firebase.auth.GoogleAuthProvider();
//   firebase.auth().signInWithPopup(provider).then(
//     user => {
//       router.push("/");
//     },
//     err => {
//       alert(err.message)
//   });
// }

</script>
<template>
  <!-- class="login-page" v-bind:class="{ error: emptyFields }"-->
  <div>
    <div class="container">
      <div class="row margin-top-12">
        <div class="col-lg-4 col-md-6 col-sm-8 mx-auto">
          <div class="card login">
            <div align="right" :hidden=true>
              <a @click.stop="switchLocale('th')">TH</a> |
              <a @click.stop="switchLocale('en')">EN</a>
            </div>
            <div align="left">
              <h1>เข้าสู่ระบบ</h1>
              <a-typography-text type="secondary">Welcome to Test System</a-typography-text>
            </div>
            <a-form layout="vertical" :model="formState">
              <br />
              <a-form-item name="username" :rules="[
                {
                  required: true,
                  message: `${$t('placeholder')}${$t(
                    'login.internal.username'
                  )}`,
                },
              ]">
                <a-input v-model:value="formState.username" :placeholder="`${$t('login.internal.username')}`"
                  allow-clear autocomplete="off">
                  <template #prefix>
                    <UserOutlined style="color: rgba(0, 0, 0, 0.25)" />
                  </template>
                </a-input>
                <!-- <label> {{ $t('login.internal.username') }} </label> -->
              </a-form-item>
              <a-form-item name="password" :rules="[
                {
                  required: true,
                  message: `${$t('placeholder')}${$t(
                    'login.internal.password'
                  )}`,
                },
              ]">
                <a-input-password v-model:value="formState.password" :placeholder="`${$t('login.internal.password')}`"
                  allow-clear autocomplete="current-password">
                  <template #prefix>
                    <LockOutlined class="site-form-item-icon" />
                  </template>
                </a-input-password>
                <!-- <label class="labeltag">{{ $t('login.internal.password') }} </label> -->
              </a-form-item>

              <a-form-item>
                <a-button type="primary" html-type="submit" block @click="handlerLogin">{{ $t("login.internal.btnlogin")
                }}</a-button>
              </a-form-item>
              <a-form-item class="center">
                <a href="#" @click="register">{{ $t("login.register") }}</a>
              </a-form-item>
            </a-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.ant-input-affix-wrapper:not(.ant-input-affix-wrapper-disabled):hover {
  border-color: #759801;
}

.center{
  text-align: center;
}

.ant-input-affix-wrapper:focus {
  border-color: #759801;
  border-right-width: 1px !important;
  outline: 0;
  box-shadow: 0 0 0 0px rgb(118 152 1 / 16%);
}

.ant-input-affix-wrapper:focus,
.ant-input-affix-wrapper-focused {
  border-color: #759801;
  border-right-width: 1px !important;
  outline: 0;
  box-shadow: 0 0 0 0px rgb(118 152 1 / 16%);
}

.ant-input-affix-wrapper {
  border-radius: 6px;
}

h1[data-v-ef68022e] {
  margin-bottom: 0.25rem;
}

h4 {
  margin-bottom: 1rem;
}

h1,
.h1,
h2,
.h2,
h3,
.h3,
h4,
.h4,
h5,
.h5,
h6,
.h6,
.h1,
.h2,
.h3,
.h4,
.h5,
.h6 {
  font-family: "Kanit-Regular", sans-serif;
  color: rgb(0 0 0 / 85%);
}

:lang(en)>q {
  quotes: "\201C""\201D""\2018""\2019";
}

:lang(fr)>q {
  quotes: "« "" »";
}

p {
  line-height: 1rem;
}

.card {
  border: 0px;
  padding: 20px;
  background: none;
  /* background: #bfbfbf; */
  /* border: 1px solid #a0a0a0; */
  border-radius: 6px;
}

.form-group,
input {
  margin-bottom: 20px;
}

.login-page {
  align-items: center;
  display: flex;
  height: 80vh;
}

.wallpaper-login {
  background: url(https://images.pexels.com/photos/32237/pexels-photo.jpg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260) no-repeat center center;
  background-size: cover;
  height: 100%;
  position: absolute;
  width: 100%;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}

.wallpaper-register {
  background: url(https://images.pexels.com/photos/533671/pexels-photo-533671.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260) no-repeat center center;
  background-size: cover;
  height: 100%;
  position: absolute;
  width: 100%;
  z-index: -1;
}

h1 {
  margin-bottom: 1.5rem;
}

.error {
  animation-name: errorShake;
  animation-duration: 0.3s;
}

@keyframes errorShake {
  0% {
    transform: translateX(-25px);
  }

  25% {
    transform: translateX(25px);
  }

  50% {
    transform: translateX(-25px);
  }

  75% {
    transform: translateX(25px);
  }

  100% {
    transform: translateX(0);
  }
}

.login-form-wrap {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
}

.card.login {
  border-radius: 10px;
  border: 2px solid #ddd;
  box-shadow: 0 4px 21px -12px rgba(0, 0, 0, .66);
  min-height: 100px;
  transition: box-shadow 0.2s ease, transform 0.2s ease;
}

.margin-top-12 {
  margin-top: 12rem;
}
</style>