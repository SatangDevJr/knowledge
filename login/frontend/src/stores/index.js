import { createStore, createLogger } from "vuex";
import createPersistedState from "vuex-persistedstate";
import { accountStore } from "./account.store";

import SecureLS from "secure-ls";
var ls = new SecureLS({ isCompression: false });

const debug = process.env.NODE_ENV !== "productiontest";

export default createStore({
    modules: {
        account: accountStore,
    },
    strict: debug,
    plugins: [
        createLogger(),
        createPersistedState({
            paths: ["account"],
            storage: {
                getItem: (key) => ls.get(key),
                setItem: (key, value) => ls.set(key, value),
                removeItem: (key) => ls.remove(key),
            },
        }),
    ],
});