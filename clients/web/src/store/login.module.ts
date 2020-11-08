import {Module} from "vuex"

export interface ILogin {
    supportAuth : boolean
    loginStates: boolean | string,
    name: string,
    avatar?: URL,
}

export const login: Module<ILogin, any> = {
    namespaced: true,
    // state: undefined,
    state: {
        supportAuth : true,
        loginStates: false,
        name: "",
        avatar: undefined,
    },
    mutations: {
        loginSuccess(state, loginInfo: ILogin) {
            Object.assign(state, loginInfo)
            state.loginStates = true
        },
        loginError(state, error : string) {
            state.loginStates = error
        },
        logout(state) {
            state.loginStates = false
        },
    },
    actions: {
        dLogin({commit}, name: String) {
            commit('loginSuccess', {
                name: name
            })
        }
    },
    getters : {
        isLogin(state) : boolean{
            if (typeof state.loginStates === "string") {
                return false
            }else{
                return state.loginStates
            }
        },
    }
}