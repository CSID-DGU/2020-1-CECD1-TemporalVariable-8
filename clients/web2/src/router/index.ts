import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
    {
        path: process.env.VUE_APP_SHOPS_URL as string,
        name: 'FoodGrid',
        component: () => import("@/views/Shops.vue"),
    },
    {
        path: process.env.VUE_APP_FAVOR_URL as string,
        name: 'Favorite',
        component: () => import('@/views/Favorite.vue')
    },
    {
        path: process.env.VUE_APP_ORDER_URL as string,
        name: 'Order',
        component: () => import('@/views/Order.vue')
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
