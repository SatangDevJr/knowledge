import { createApp } from "vue";
import { router } from "./router";
import App from "./App.vue";
import { createPinia } from "pinia";
import { i18n } from "./i18n";
import Antd from "ant-design-vue";
import antdConfig from "./antd.config";
import { axios } from "axios";
import transformingData from "./filter/transformingData.filter";

import "./assets/vendors/mdi/css/materialdesignicons.min.css";
import "./assets/css/demo_2/style.css";
import "ant-design-vue/dist/antd.css";
import "./assets/css/custom.css";
import "./assets/js/THSarabunNew Italic-italic";
import "./assets/js/THSarabunNew-normal";
import "./assets/js/THSarabunNew Bold-normal";
import "./assets/js/THSarabunNew BoldItalic-normal";
// import firebase from 'firebase/app'

// const config = {
//     apiKey: 'AIzaSyDK3HZ2EJTI_3N2ya26MFQ8xLfj0v6ab78',
//     authDomain: 'test2023-45324.firebaseapp.com',
//     databaseURL: 'https://test2023-45324.firebaseio.com',
//     projectId: 'test2023-45324',
//     storageBucket: 'test2023-45324.appspot.com',
//     messagingSenderId: '153149032163'
// };
// firebase.initializeApp(config)

const app = createApp(App);
app.config.globalProperties.$filters = transformingData;
app.config.errorHandler = (err, instance, info) => {
    console.log("errorHandler", err, instance, info);
};


app
    .use(createPinia())
    .use(router)
    .use(i18n)
    .use(axios)
    .use(Antd)
    .use(antdConfig)
    .mount("#app");