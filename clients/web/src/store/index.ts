import Vue from 'vue'
import Vuex from 'vuex'
import {login} from './login.module'

Vue.use(Vuex)

export default new Vuex.Store({
  mutations : {
  },
  modules : {
    login : login,
  }
})
