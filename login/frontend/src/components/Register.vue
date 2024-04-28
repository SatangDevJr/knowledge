<script setup>
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import { reactive, ref } from "vue";
import { router } from "../router/index";
import { useI18n } from "vue-i18n";
import useUser from '../hooks/useUser';


const { locale, t } = useI18n();

function switchLocale(lang) {
  locale.value = lang;
}
const formRef = ref()

const formState = reactive({
  username: "",
  password: "",
  confirmPassword:"",
});


const handlerLogin = () => {
  formRef.value.validate().then(async () => {
    if (formState.username && formState.password ) {
      const param = {
        username: formState.username,
        password: formState.password,
      }

      setTimeout(() => {
        useUser.onSaveUser(param).then((res) => {
          clearFormState();
          if (res.status == 200) {
            router.push("/login");
          }
        });
      },500)

    }
  })
  .catch((error) => {
  });
};

const clearFormState = () =>{
  formState.username = "";
  formState.password = "";
  formState.confirmPassword = ""
}

let validatePassConfirmPassword = async (_rule, value) => {
      console.log("value ", value)
      if (value === '') {
        return Promise.reject('Please input the password again');
      } else if (value !== formState.password) {
        return Promise.reject("Two inputs don't match!");
      } else {
        return Promise.resolve();
      }
    };

const rules = {
      username: [{ required: true, message: t('placeholder') + t('login.internal.username') }],
      password: [{ required: true, message: t('placeholder') + t('login.internal.password') }],
      confirmPassword: [{ required: true, validator: validatePassConfirmPassword }],
};


</script>
<template>
  <!-- class="login-page" v-bind:class="{ error: emptyFields }"-->
  <div>
    <div class="container">
      <div class="row" style="margin-top:250px">
        <div class="col-lg-4 col-md-6 col-sm-8 mx-auto">
          <div class="card login">
            <div align="right" :hidden=true>
              <a @click.stop="switchLocale('th')">TH</a> |
              <a @click.stop="switchLocale('en')">EN</a>
            </div>
            <div align="left">
              <h1>สมัครสมาชิก</h1>
              <a-typography-text type="secondary">Welcome to Test System</a-typography-text>
            </div>
            <a-form layout="vertical" ref="formRef" :model="formState" :rules="rules" >
              <br />
              <a-form-item name="username">
                <a-input v-model:value="formState.username" :placeholder="`${$t('login.internal.username')}`"
                  allow-clear autocomplete="off">
                  <template #prefix>
                    <UserOutlined style="color: rgba(0, 0, 0, 0.25)" />
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item name="password" >
                <a-input-password v-model:value="formState.password" :placeholder="`${$t('login.internal.password')}`"
                  allow-clear autocomplete="current-password">
                  <template #prefix>
                    <LockOutlined class="site-form-item-icon" />
                  </template>
                </a-input-password>
              </a-form-item>

              <a-form-item name="confirmPassword">
                <a-input-password v-model:value="formState.confirmPassword" :placeholder="`${$t('login.internal.confirmpassword')}`"
                  allow-clear autocomplete="current-password">
                  <template #prefix>
                    <LockOutlined class="site-form-item-icon" />
                  </template>
                </a-input-password>
              </a-form-item>

              <a-form-item>
                <a-button type="primary" html-type="submit" block @click="handlerLogin">{{ $t("login.register")
                }}</a-button>
              </a-form-item>
            </a-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.center{
  text-align: center;
}
.ant-input-affix-wrapper:not(.ant-input-affix-wrapper-disabled):hover {
  border-color: #759801;
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
</style>