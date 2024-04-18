import { createRouter, createWebHistory, RouterView } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView,
    },
    {
      path: '/cidr',
      name: 'Cidr',
      component: () => import('../views/CidrView'),
    },
    {
      path: '/json',
      name: 'json',
      component: RouterView,
      children: [
        {
          path: 'jwt',
          component: () => import('../views/JwtDecodeView'),
        },
        {
          path: 'json-beautify',
          component: () => import('../views/JsonBeautifyView'),
        },
        {
          path: 'json-to-yaml',
          component: () => import('../views/JsonToYamlView'),
        },
      ],
    },
    {
      path: '/hash',
      name: 'hash',
      component: () => import('../views/HashView'),
    },
    {
      path: '/base64',
      name: 'Base64',
      component: RouterView,
      children: [
        {
          path: 'encode',
          component: () => import(/* webpackChunkName: "Base64" */ '../views/Base64View'),
          props: {
            type: 'encode',
          },
        },
        {
          path: 'decode',
          component: () => import(/* webpackChunkName: "Base64" */ '../views/Base64View'),
          props: {
            type: 'decode',
          },
        },
      ],
    },
  ],
})

export default router
