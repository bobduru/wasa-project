import {createRouter, createWebHashHistory} from 'vue-router'
import UserProfileView from '../views/UserProfileView.vue'
import LoginView from '../views/LoginView.vue'
import FeedView from '../views/FeedView.vue'
import CreatePostView from '../views/CreatePostView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: FeedView},
		{path: '/user/:userId',name:"profile", component: UserProfileView},
		{path: '/login', component: LoginView},
		{path: '/create', component: CreatePostView},
		{path: '/link2', component: UserProfileView},
		{path: '/some/:id/link', component: UserProfileView},
	]
})

export default router
