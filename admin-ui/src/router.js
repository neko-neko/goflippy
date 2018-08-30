import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/projects'
    }, {
      path: '/projects',
      component: () => import('./views/projects/Index.vue')
    }, {
      path: '/projects/:id',
      component: () => import('./views/projects/Show.vue')
    }, {
      path: '/projects/:id/users/:uuid',
      component: () => import('./views/users/Show.vue')
    }, {
      path: '/projects/:id/features/:key',
      component: () => import('./views/features/Show.vue')
    },
  ]
})
