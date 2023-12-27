import {createRouter, createWebHashHistory} from 'vue-router'
import UserProfileView from '../views/UserProfileView.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: UserProfileView},
		{path: '/login', component: LoginView},
		{path: '/link1', component: UserProfileView},
		{path: '/link2', component: UserProfileView},
		{path: '/some/:id/link', component: UserProfileView},
	]
})

export default router
