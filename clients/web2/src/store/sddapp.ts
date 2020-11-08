import Vue from 'vue'
import Vuex from 'vuex'
import {Service} from "@/classes/ddapp/core.js";
import {fold, fromNullable} from "fp-ts/lib/Option";
import {MockDDApp} from "@/classes/mock/mock.bc.net";
import moment from "moment";
import {AllQuery} from "@/classes/ddapp/query";
import {v4} from "uuid";
import {isLeft} from "fp-ts/lib/Either";

Vue.use(Vuex)

const USERS = [
    {id: "주식회사 분식", pw: "password"},
    {id: "놀부사업", pw: "password"},
    {id: "가족경영 방만", pw: "password"},
    {id: "공기업 나태", pw: "password"},
    {id: "(주) 주", pw: "password"},
]
const RESTAURANT = [
    {name: "동국짜장면", hint: "중국집", owner: USERS[0].id, addr: v4()},
    {name: "샤브샤브", hint: "맛있다", owner: USERS[0].id, addr: v4()},

    {name: "컴공파스타", hint: "패밀리 레스토랑", owner: USERS[1].id, addr: v4()},

    {name: "미국버거", hint: "칼로리 폭탄", owner: USERS[2].id, addr: v4()},
    {name: "치킨천지", hint: "치느님", owner: USERS[2].id, addr: v4()},

    {name: "엄마손맛", hint: "한정식", owner: USERS[3].id, addr: v4()},
    {name: "스시전문 이시국", hint: "일식", owner: USERS[3].id, addr: v4()},

    {name: "스시전문 이시국", hint: "일식", owner: USERS[3].id, addr: v4()},


    {name: "가게이름 짓기도 힘들어 죽겠네 0", hint: "대충", owner: USERS[4].id, addr: v4()},
    {name: "가게이름 짓기도 힘들어 죽겠네 1", hint: "대충", owner: USERS[4].id, addr: v4()},
    {name: "가게이름 짓기도 힘들어 죽겠네 2", hint: "대충", owner: USERS[4].id, addr: v4()},
    {name: "가게이름 짓기도 힘들어 죽겠네 3", hint: "대충", owner: USERS[4].id, addr: v4()},
    {name: "가게이름 짓기도 힘들어 죽겠네 4", hint: "대충", owner: USERS[4].id, addr: v4()},
]

export default new Vuex.Store({
    state: {
        inited: false,
        ddapp: new MockDDApp(),
        cached: new Array<{ value: Service; validAt: moment.Moment; key: string }>(),
        loadUnit: 30,
        location: "",
        validDuration: moment.duration(3, 'minute'),
        error: undefined,
    },
    mutations: {},
    actions: {
        reload: async (injectee, {request, limit}) => {
            if (!injectee.state.inited) {
                console.log("Initializing")
                for (const user of USERS) {
                    const res = injectee.state.ddapp.addUser(user.id, user.pw)
                    if (isLeft(res)) {
                        console.log(`[USER] : add fail cause '${res.left}' for ${user}`)
                    }
                }
                for (const r of RESTAURANT) {
                    const res = injectee.state.ddapp.addRestaurant(r.name, r.hint, r.owner, r.addr)
                    if (isLeft(res)) {
                        console.log(`[RESTAURANT] : add fail cause '${res.left}' for ${r}`)
                    }
                }
                injectee.state.inited = true
            }
            const cpage = fold(
                () => injectee.state.loadUnit,
                a => {
                    if (typeof a === 'number') {
                        if (!Number.isInteger(a)) {
                            return injectee.state.loadUnit
                        }
                        if (a < 1) {
                            return injectee.state.loadUnit
                        }
                        return a
                    } else {
                        return injectee.state.loadUnit
                    }
                }
            )(fromNullable(limit))
            // TODO : page
            //
            const svrs = await injectee.getters.ddApp.services(request === undefined ? new AllQuery() : request)
            injectee.state.cached.length = 0
            for (const svr of svrs) {
                injectee.state.cached.push({
                    value: svr,
                    validAt: moment().add(injectee.state.validDuration),
                    key: svr.query.key(),
                })
            }
        }

    },
    modules: {},
    getters: {
        ddApp: state => state.ddapp,
    }
})
