import Vue from 'vue'
import VueRouter from 'vue-router'
import ItemList from '../views/ItemList.vue'
import NewItem from "@/views/NewItem";
import store from "../store/Store";

Vue.use(VueRouter)

const routes = [
  
    {
        path: '/',
        name: 'ItemList',
        component: ItemList,
        meta: {
            requiresAuth: true
        }
    },
    {
        path: '/new',
        name: 'NewItem',
        component: NewItem,
        meta: {
            requiresAuth: true
        }
    },

]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        const userInfo = store.userInfo
        if (userInfo == null) {
            next({path: '/login', params: {nextUrl: to.fullPath}})
            return
        }
        if (to.meta.role && !userInfo.roles.includes(to.meta.role)) {
          next({path: '/login', params: {nextUrl: to.fullPath}})
          return
        }
        next()
        return
    } 
    next()
    
})
export default router
