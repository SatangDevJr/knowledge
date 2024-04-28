import { defineStore, acceptHMRUpdate } from "pinia";

export const useAccountStore = defineStore({
  id: "account",
  state: () => ({
    isLogin: false,
    isPreview: false,
    loginData: {
      User: {},
      Role: [],
      Token: "",
      MenuInfo: [],
    }
  }),
  mutations: {
    setLoginData(state, value){
      state.loginData = {};
      state.loginData = value;
    },
  },
  getters: {
    getAuthen() {
      return this.isLogin;
    },
    getIsPreviewPDF() {
      return this.isPreview;
    },

    getLoginData(state){
      return state.loginData;
    },
  },
  actions: {
    setAuthen() {
      this.$patch({
        isLogin: true,
      });
    },
    setPreviewPDF() {
      this.$patch({
        isPreview: true,
      });
    },
    setloginData(response){
      this.$patch({
        loginData: response,
      });
    },
  },
});
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAccountStore, import.meta.hot));
}
