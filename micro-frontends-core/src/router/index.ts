import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/About.vue')
  },
  {
    path:"/traefikUI",
    name:"TraefikUI",
    component:()=>import("@/views/Traefik.vue")
  },{
    path:"/traefikDoc",
    name:"TraefikDoc",
    component:()=>import("@/views/TraefikDoc.vue")
  },
  {
    path:"/remcoDoc",
    name:"RemcoDoc",
    component:()=>import("@/views/RemcoDoc.vue")
  },
  {
    path:"/consulManager",
    name:"ConsulManager",
    component:()=>import("@/views/ConsulManager.vue")
  },
  {
    path:"/portainer",
    name:"Portainer",
    component:()=>import("@/views/Portainer.vue")
  },
  {
    path: "/:vue(vue||vue/+.*)",
    name: 'Vue',
    component: () => import('@/views/Vue.vue')
  },
  {
    path: "/:ncda(ncda||ncda/+.*)",
    name: 'Ncda',
    component: () => import('@/views/Ncda.vue')
  },
]

const router = createRouter({
  // history: createWebHistory(process.env.BASE_URL),
  history: createWebHistory("/"),
  routes
})

export default router
