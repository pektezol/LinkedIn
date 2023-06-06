import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Jobs from "../views/Job.vue";
import Network from "../views/Network.vue";
import Profile from "../views/Profile.vue";
import Login from "../views/Login.vue"
import Register from "../views/Register.vue"
import User from "../views/User.vue"
import Company from "../views/CompanyProfile.vue"
Vue.use(VueRouter);

const router = new VueRouter({
    mode: "history",
    base: import.meta.env.BASE_URL,
    routes: [{
            path: "/login",
            name: "login",
            component: Login,
        },
        {
            path: "/register",
            name: "register",
            component: Register,
        },
        {
            path: "/",
            name: "home",
            component: Home,
        },
        {
            path: "/job",
            name: "job",
            component: Jobs,
        },
        {
            path: "/network",
            name: "network",
            component: Network,
        },
        {
            path: "/profile",
            name: "profile",
            component: Profile,
        },
        {
            path: "/user/:user_name",
            name: "user",
            component: User,
        },
        {
            path: "/companyProfile",
            name: "company",
            component: Company,
        }
    ],
});

export default router;