import Vue from 'vue'
import Vuex from 'vuex'
import {DDApp, Service, UniqueQuery} from "@/classes/ddapp/core.js";
import {fold, fromNullable} from "fp-ts/lib/Option";
import {MockDDApp} from "@/classes/mock/mock.bc.net";
import moment from "moment";

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        ddapp: new MockDDApp() as DDApp,
        cached: new Array<{ value: Service, validAt: moment.Moment }>(),
        loadUnit: 100,
        location: "",
    },
    mutations: {},
    actions: {
        reload: (injectee, {request, page}) => {
            const app = injectee.getters.ddApp
            app.service(request)
            const cpage = fold(
                () => 1,
                a => {
                    if (typeof a === 'number') {
                        if (!Number.isInteger(a)) {
                            return 1
                        }
                        if (a < 1) {
                            return 1
                        }
                        return a
                    } else {
                        return 1
                    }
                }
            )(fromNullable(page))

        }

    },
    modules: {},
    getters: {
        ddApp: state => state.ddapp,
    }
})
