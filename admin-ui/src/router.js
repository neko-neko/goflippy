import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      redirect: "/projects"
    },
    {
      path: "/projects",
      component: () => import("./pages/projects/Index.vue")
    },
    {
      path: "/projects/:id",
      component: () => import("./pages/projects/Show.vue")
    },
    {
      path: "/projects/:id/users/:uuid",
      component: () => import("./pages/users/Show.vue")
    },
    {
      path: "/projects/:id/features/:key",
      component: () => import("./pages/features/Show.vue")
    }
  ]
});
