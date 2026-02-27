// javascript
import { createRouter, createWebHistory } from 'vue-router'
import GameView from '@/views/game.vue'
import HomeView from "@/views/home.vue";
import LoginView from "@/views/auth/login.vue";
import RegisterView from "@/views/auth/register.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        { path: '/', name: 'Home', component: HomeView},
        { path: '/auth/login', name: 'Login', component: LoginView},
        { path: '/auth/register', name: 'Register', component: RegisterView},
        { path: '/game', name: 'Game', component: GameView,},
    ],
})

const publicPaths = ['/', '/auth/login', '/auth/register']

function parseJwt(token) {
    try {
        const base64Url = token.split('.')[1]
        const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
        const jsonPayload = decodeURIComponent(
            atob(base64)
                .split('')
                .map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
                .join('')
        )
        return JSON.parse(jsonPayload)
    } catch (e) {
        return null
    }
}

function isTokenValid(token) {
    if (!token)
        return false

    const payload = parseJwt(token)
    if (!payload || !payload.exp)
        return false

    return payload.exp * 1000 > Date.now()
}

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')

    if (publicPaths.some(p => to.path === p || to.path.startsWith(p + '/'))) {
        // si on va vers login/register et qu'on a un token valide -> redirige
        if ((to.path === '/auth/login' || to.path === '/auth/register') && isTokenValid(token)) {
            return next('/')
        }
        // si token présent mais invalide, le supprimer pour éviter un état stale
        if (token && !isTokenValid(token)) {
            localStorage.removeItem('token')
        }
        return next()
    }

    // routes protégées : si token valide -> ok, sinon supprimer et rediriger vers login
    if (isTokenValid(token)) return next()

    if (token) {
        localStorage.removeItem('token')
    }
    next({ path: '/auth/login', query: { redirect: to.fullPath } })
})

export default router
