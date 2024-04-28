import { createRouter, createWebHistory } from "vue-router";
import Login from "../components/Login.vue";
import Register from "../components/Register.vue";
import { logoutUser, isLoggedIn } from '../utils/auth'

const routes = [{
        path: "/",
        name: "/",
        component: () =>
            import ("../views/Home.vue"),
        beforeEnter: authGuard,
    },
    {
        path: "/login",
        name: "/login",
        component: Login,
        meta: {
            guest: true,
        },
    },
    {
        path: "/register",
        name: "/register",
        component: Register,
        meta: {
            guest: true,
        },
    },

];

export const router = createRouter({
    history: createWebHistory(
        import.meta.env.BASE_URL),
    mode: "history",
    base: "/",
    routes,
    scrollBehavior(to, from, savedPosition) {
        if (to.hash) {
            // BEFORE:
            // return { selector: to.hash }

            return { el: to.hash }
        }
    },
});

export function authGuard(to) {
    const islogin = isLoggedIn()
    if (islogin) {
        return true;
    } else {
        setTimeout(() => {
            logoutUser();
            router.push("/login");
            return false;
        }, 500);

    }
}