import {login} from '@/api/user'
import {jsonInBlacklist} from '@/api/jwt'
import router from '@/router/index'

export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            nickName: "",
            headerImg: "",
            authority: "",
            phone: ""
        },
        token: "",
    },
    mutations: {
        setUserInfo(state, userInfo) {
            // 这里的 `state` 对象是模块的局部状态
            state.userInfo = userInfo
        },
        setToken(state, token) {
            // 这里的 `state` 对象是模块的局部状态
            state.token = token
        },
        NeedInit(state) {
            state.userInfo = {}
            state.token = ""
            sessionStorage.clear()
            router.push({name: 'init', replace: true})

        },
        LoginOut(state) {
            state.userInfo = {}
            state.token = ""
            sessionStorage.clear()
            router.push({name: 'login', replace: true})
            window.location.reload()
        },
        ResetUserInfo(state, userInfo = {}) {
            state.userInfo = {
                ...state.userInfo,
                ...userInfo
            }
        }
    },
    actions: {
        async LoginIn({commit, dispatch, rootGetters, getters}, loginInfo) {
            const res = await login(loginInfo)
            if (res.code === 0) {
                const path = process.env.VUE_APP_BASE_API;
                if (res.data.user.headerImg === "") {
                    if (res.data.user.authority.authorityId === "200" || res.data.user.authority.parentId === "200") {
                        res.data.user.headerImg = require("@/assets/students.png")
                    } else {
                        res.data.user.headerImg = require("@/assets/coach.png")
                    }
                } else if (res.data.user.headerImg.slice(0, 4) !== 'http') {
                    res.data.user.headerImg = path + res.data.user.headerImg
                }
                commit('setUserInfo', res.data.user)
                commit('setToken', res.data.token)
                await dispatch('router/SetAsyncRouter', {}, {root: true})
                const asyncRouters = rootGetters['router/asyncRouters']
                router.addRoutes(asyncRouters)
                // const redirect = router.history.current.query.redirect
                // console.log(redirect)
                // if (redirect) {
                //     router.push({ path: redirect })
                // } else {
                router.push({name: getters["userInfo"].authority.defaultRouter})
                // }
                return true
            }
        },
        async LoginOut({commit}) {
            const res = await jsonInBlacklist()
            if (res.code === 0) {
                commit("LoginOut")
            }
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        },
        token(state) {
            return state.token
        },

    }
}