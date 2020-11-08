import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'

Vue.use(VueRouter)
import Search from '@/views/Search.vue'
import Favorite from '@/views/Favorite.vue'
import History from '@/views/History.vue'
import Config from '@/views/Config.vue'
import Error from '@/views/Error.vue'
const routes: Array<RouteConfig> = [
    {
        alias : '',
        path : '/search',
        component : Search
    },
    {
        alias : '',
        path : '/favorite',
        component : Favorite
    },
    {
        alias : '',
        path : '/history',
        component : History
    },
    {
        alias : '',
        path : '/config',
        component : Config
    },
    {
        alias : '',
        path : '*',
        component : Error
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
