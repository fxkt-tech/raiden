import {
    createRouter,
    createWebHistory,
} from 'vue-router'
import Framework from '../components/Framework.vue'
import Home from '../pages/Home.vue'

const routes = [
    {
        path: '/',
        name: 'Framework',
        component: Framework,
        children: [
            {
                path: "home",
                name: "Home",
                component: Home,
            }
        ]
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
