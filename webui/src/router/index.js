import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import LogoutView from '../views/LogoutView.vue'
import ProfileView from '../views/ProfileView.vue'
import PostPhotoView from '../views/PostPhotoView.vue'
import UserSearchView from '../views/SearchView.vue'
import StreamView from '../views/StreamView.vue'
import SetMyUserNameView from '../views/SetMyUserNameView.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/session/', component: LoginView},
		{path: '/photos/', component: PostPhotoView},
		{path: '/users/', component: UserSearchView},
		{path: '/set-name/', component: SetMyUserNameView},
		{path: '/users/:userId/profile', component: ProfileView},
		{path: '/users/:userId/stream', component: StreamView},

	]
})

export default router