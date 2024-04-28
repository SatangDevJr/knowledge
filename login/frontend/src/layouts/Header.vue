<template>
  <div class="horizontal-menu">
    <nav class="navbar top-navbar col-lg-12 col-12 p-0">
      <div class="container">
        <div
          class="
            text-center
            navbar-brand-wrapper
            d-flex
            align-items-center
            justify-content-center
          "
        >
          <router-link class="nav-link" to="/">
            <a class="navbar-brand brand-logo">
              <h3 style="color: #ffffff">{{ $t("systemname") }}</h3>
              <span
                v-if="env !== 'production'"
                class="font-12 d-block font-weight-light"
                >version:{{ version }} ({{ env }})</span
              >
              <span
                v-if="env == 'production'"
                class="font-12 d-block font-weight-light"
                >version:{{ version }}</span
              >
            </a>
          </router-link>
        </div>
        <div
          class="
            navbar-menu-wrapper
            d-flex
            align-items-center
            justify-content-end
          "
        >
          <ul class="navbar-nav navbar-nav-right">
            <li class="nav-item" hidden="true">
              <div class="nav-link d-flex">
                <div class="nav-item dropdown">
                  <a
                    class="
                      nav-link
                      count-indicator
                      dropdown-toggle
                      text-black
                      font-weight-semibold
                    "
                    id="notificationDropdown"
                    href="#"
                    data-bs-toggle="dropdown"
                  >
                    <span class="font-13 online-color">
                      <i v-bind:class="`flag flag-${locale}`"></i>
                      {{ $t("language") }}
                    </span>
                  </a>
                  <div
                    class="dropdown-menu navbar-dropdown"
                    aria-labelledby="notificationDropdown"
                  >
                    <a
                      class="dropdown-item"
                      v-for="availableLocale in SUPPORT_LOCALES"
                      value="availableLocale.langcode"
                      :key="availableLocale"
                      @click.stop="switchLocale(availableLocale)"
                    >
                      <span class="mr-4">
                        <i
                          v-bind:class="'flag flag-' + availableLocale.langcode"
                        ></i>
                      </span>
                      {{ availableLocale.langname }}
                    </a>
                  </div>
                </div>
              </div>
            </li>
            <li hidden="true"
              class="
                text-center
                navbar-brand-wrapper
                d-flex
                align-items-center
                justify-content-center
              "
            >
              <a-divider
                type="vertical"
                style="height: 30px; background-color: #ffffff"
              />
            </li>
            <li class="nav-item nav-profile dropdown">
              <a
                class="nav-link"
                id="profileDropdown"
                href="#"
                data-bs-toggle="dropdown"
                aria-expanded="false"
              >
                <div v-if="roleAdmin" class="nav-profile-img">
                  <img src="../assets/images/faces/admin.jpg" alt="image" />
                </div>
                <div v-else-if="roleAccount" class="nav-profile-img">
                  <img src="../assets/images/faces/ac.jpg" alt="image" />
                </div>
                <div v-else-if="roleBD" class="nav-profile-img">
                  <img src="../assets/images/faces/th.jpg" alt="image" />
                </div>
                <div v-else-if="rolePartner" class="nav-profile-img">
                  <img src="../assets/images/faces/pat.jpg" alt="image" />
                </div>
                <div v-else-if="roleSubPartner" class="nav-profile-img">
                  <img src="../assets/images/faces/sub.jpg" alt="image" />
                </div>
                <div v-else class="nav-profile-img">
                  <img src="../assets/images/faces/face1.jpg" alt="image" />
                </div>
                <div class="nav-profile-text">
                  <p class="font-weight-semibold m-0"></p>
                  <p class="font-10 online-color">
                    <i class="mdi mdi-chevron-down">{{ userCode }}</i>
                  </p>
                </div>
              </a>
              <div
                class="dropdown-menu navbar-dropdown"
                aria-labelledby="profileDropdown"
              >
                <a class="dropdown-item" @click="onLogout">
                  <i class="mdi mdi-logout me-2 text-primary">{{ $t("btnsignout") }}</i>
                </a>
              </div>
            </li>
          </ul>
          <button
            class="
              navbar-toggler navbar-toggler-right
              d-lg-none
              align-self-center
            "
            type="button"
            data-toggle="horizontal-menu-toggle"
          >
            <span class="mdi mdi-menu"></span>
          </button>
        </div>
      </div>
    </nav>

    <nav class="bottom-navbar">
      <div class="container">
        <ul class="nav page-navigation">
          <li class="nav-item" v-for="menu in menulist" :key="menu">
            <router-link class="nav-link" :to="`${menu.url}`">
              <i :class="`${menu.iconClassName}`"></i>
              <span class="menu-title">{{ menu.name }}</span>
              <i class="menu-arrow" v-if="menu.children.length > 0"></i>
            </router-link>
            <div class="submenu" v-if="menu.children.length > 0">
              <ul class="submenu-item">
                <li
                  class="nav-item"
                  v-for="children in menu.children"
                  :key="children"
                >
                  <router-link
                    class="nav-link"
                    :to="`${children.url}`"
                  >
                    <i :class="`${children.iconClassName}`"></i>
                    <span class="menu-title">{{ children.name }}</span>
                  </router-link>
                </li>
              </ul>
            </div>
          </li>
        </ul>
      </div>
    </nav>
  </div>
</template>

<script setup>
import axios from 'axios'
import { ref,onMounted } from "vue";
import { SUPPORT_LOCALES } from "./../i18n";
import { useI18n } from "vue-i18n";
import { useAccountStore } from "../stores/account.store";
import { router } from "../router/index";
import { logoutUser } from '../utils/auth'
import { isLoggedIn } from '../utils/auth'
import { loginUser } from '../utils/auth'
import { Modal } from 'ant-design-vue';
import { storageService } from '../services/storage.service'

const accountStore = useAccountStore();
const { locale, t } = useI18n();
const langs = ref([]);
const env = import.meta.env.VITE_APP_MODE;
const version = import.meta.env.VITE_APP_Version;
const userName = ref("")
const userCode = ref("")

const onLogout = () => {
  setTimeout(() => {
    logoutUser();
    router.push("/login");
  }, 500);
};

const menulist = ref([]);
const childrenMain = ref([]);
const isMain = ref(false);
const isSetting = ref(false);
const childSetting = ref([]);
const rolePartner = ref(false);
const roleSubPartner = ref(false);
const roleAccount = ref(false);
const roleBD = ref(false);
const roleAdmin = ref(false);

function onSettingMenu() {
  menulist.value = [];
  childSetting.value = [];
  childrenMain.value = [];
  
  let loginData = accountStore.getLoginData;
  if (loginData != undefined){
    let nameDisplay = loginData.User.firstnameth + " "+ loginData.User.lastnameth;
    if(loginData.User.type == "External"){
      if(loginData.User.companyname != ""){
        userName.value = loginData.User.companyname;
      }else{
        userName.value = nameDisplay;
      }
    }else{
      userName.value = nameDisplay;
    }
    userCode.value = loginData.User.username;

    rolePartner.value = false;
    roleSubPartner.value = false;
    roleAccount.value = false;
    roleAdmin.value = false;
    roleBD.value = false;

    loginData.Role.forEach((element) => {
      if (element.rolecode == "PAT"){
          rolePartner.value = true;
      }
      if (element.rolecode == "SUB"){
          roleSubPartner.value = true;
      }
      if (element.rolecode == "ACC"){
          roleAccount.value = true;
      }
      if (element.rolecode == "BDA"){
          roleBD.value = true;
      }
      if (element.rolecode == "SUA"){
          roleAdmin.value = true;
      }
    });

    if(loginData.MenuInfo != undefined){
      if(loginData.MenuInfo.length !== 0){
        let menuName =  "";
        let iconName = "mdi mdi-compass-outline menu-icon"
        loginData.MenuInfo.forEach((element) => {
          menuName =  "";
          iconName = "mdi mdi-compass-outline menu-icon";

          if(element.coremenu == "ออกใบส่งสินค้า"){
            menuName = `${t("menu.menu2")}`
            iconName = "mdi mdi-file-document-edit-outline menu-icon";
          }
          if(element.coremenu == "เอกสารวางบิล"){
            menuName = `${t("menu.menu3")}`
            iconName = "mdi mdi-clipboard-check-outline menu-icon";
          }
          if(element.coremenu == "รายการวางบิล"){
            menuName = `${t("menu.menu4")}`
            iconName = "mdi mdi-clipboard-arrow-left-outline menu-icon";
          }
          if(element.submenu == "ลงทะเบียน Whose Sales"){
            menuName = `${t("menu.menu5")}`
            iconName = "mdi mdi-account-card-details menu-icon submenu-icon";
          }
          if(element.submenu == "Manage Quota Whose Sales"){
            menuName = `${t("menu.menu8")}`
            iconName = "mdi mdi-library-books menu-icon submenu-icon";
          }
          if(element.submenu == "Reset Print Ticket"){
            menuName = `${t("menu.menu9")}`
            iconName = "mdi mdi-printer-settings menu-icon submenu-icon";
          }
          if(element.submenu == "จัดการข้อมูลผู้ใช้งาน"){
            menuName = `${t("menu.menu10")}`
            iconName = "mdi mdi-account-settings menu-icon submenu-icon";
          }

          if(element.coremenu == "หน้าหลัก"){
            childrenMain.value.push({
              name:  menuName,
              url:   "/" + element.menucode,
              iconClassName: iconName
            });
            isMain.value = true;
          }else{
            if(isMain.value == true){
              menulist.value.push({
                name: `${t("menu.menu6")}`,
                url: "/",
                iconClassName: "mdi mdi-apps menu-icon",
                children: childrenMain.value,
              });
            }
            if(element.coremenu == "จัดการข้อมูล"){
              childSetting.value.push({
                name:  menuName,
                url:   "/" + element.menucode,
                iconClassName: iconName
              });
              isSetting.value = true;
            }else{
              menulist.value.push({
                name: menuName,
                url: "/"  + element.menucode,
                iconClassName: iconName,
                children: [],
              });
            }
            isMain.value = false;
          }
        });
        if(isSetting.value == true){
          menulist.value.push({
            name: `${t("menu.menu7")}`,
            url: "/",
            style:"float:'right'",
            iconClassName: "mdi mdi-settings menu-icon",
              children: childSetting.value,
          });
          isSetting.value = false;
        }
      }else{
        var warningword = `${t("welcomewordstart")}` + " " + loginData.User.username + " " + `${t("welcomewordend")}`
         Modal.warning({
          title: `${t("titlesystem")}`,
          content: warningword,
        });
      }

      if(roleAccount.value == true){
        router.push("/deliverynotelist");
      }else{
        router.push("/");
      }
    }
  }
}

async function login(username,password,stafflogin) {
  try {
    await loginUser(username, password.replace("\"","").replace("\"",""),stafflogin);
    onSettingMenu();
  } catch (err) {
    router.push("/login");                 
  }
}

function onloadMenu() {
  let loginData = accountStore.getLoginData;
  if (loginData.Token == ""){
      if(localStorage.getItem("UserName") != undefined && localStorage.getItem("Password") != undefined && localStorage.getItem("StaffLogin") != undefined){
        let staffLogin = true;
        if(localStorage.getItem("StaffLogin").toLocaleLowerCase() == "false"){
           staffLogin = false;
        }
         login(localStorage.getItem("UserName"),storageService.getStorage("Password"),staffLogin);
      }
  }else{
    if (isLoggedIn){
      onSettingMenu();
    }else{
      logoutUser();
      router.push("/login");      
    }
  }
}



function switchLocale(lang) {
  locale.value = lang.langcode;
  onSettingMenu();
}

onMounted(() => {
  onloadMenu();
});

</script>

<style scoped>
@import "../assets/css/flags.css";
@media (min-width: 768px) {
  .horizontal-menu
    .bottom-navbar
    .page-navigation
    > .nav-item:not(.mega-menu)
    .submenu {
    top: 2rem;
  }
}
/*
.horizontal-menu
  .bottom-navbar
  .page-navigation
  > .nav-item:not(.mega-menu)
  .submenu
  ul {
  width: auto;
  padding: 10px;
}*/

.horizontal-menu .bottom-navbar .page-navigation>.nav-item>.nav-link {
    padding: 10px 40px 10px 7px !important;
    line-height: 1 !important;
    font-weight: 400 !important;
}
.nav-link i {
  font-size: 14px !important; 
}
/* .horizontal-menu
  .top-navbar
  .navbar-menu-wrapper
  .navbar-nav
  .nav-item.dropdown
  .navbar-dropdown {
  position: absolute;
  font-size: 0.9rem;
  margin-top: 0;
  right: 0;
  left: auto;
  top: 0px;
} */
</style>