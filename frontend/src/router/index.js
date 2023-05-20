import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Jobs from "../views/Job.vue";
import Network from "../views/Network.vue";
import Profile from "../views/Profile.vue";

Vue.use(VueRouter);

const router = new VueRouter({
  mode: "history",
  base: import.meta.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      component: Home ,
    },
    {
      path: "/job",
      name: "job",
      component: Jobs ,
    },
    {
      path: "/network",
      name: "network",
      component: Network ,
    },
    {
      path: "/profile",
      name: "profile",
      component: Profile ,
    }
  ],
});

export default router;
