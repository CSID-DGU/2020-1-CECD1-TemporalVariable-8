import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store/sddapp'
import vuetify from './plugins/vuetify';
import vueSmoothScroll from 'vue2-smooth-scroll'

Vue.config.productionTip = false

Vue.use(vueSmoothScroll)

const vm = new Vue({
    router,
    store,
    vuetify,
    render: h => h(App),
})
//
vm.$mount('#app')

