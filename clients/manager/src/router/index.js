import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Menu from '../views/Menu.vue'
import Add from '../views/Add.vue'
import OrderVie from '../views/OrderVie.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/menu',
    name: 'Menu',
    component: Menu
  },
  {
    path: '/add',
    name: 'Add',
    component: Add
  },
  {
    path: '/order',
    name: 'Order',
    component: OrderVie
  }
]

const router = new VueRouter({
  routes
})

export default router
