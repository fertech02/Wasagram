import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import LogoutView from '../views/LogoutView.vue'
import ProfileView from '../views/ProfileView.vue'
import UploadPhotoView from '../views/UploadPhoto.vue'
import SearchUserView from '../views/SearchUserView.vue'
import StreamView from '../views/StreamView.vue'
import SetUserNameView from '../views/SetUserNameView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/session/', component: LoginView},
		{path: '/logout/', component: LogoutView},
		{path: '/users/:uid/profile/', component: ProfileView},
		{path: '/photos/', component: UploadPhotoView},
		{path: '/users/:uid/stream/', component: StreamView},
		{path: '/set-name/', component: SetUserNameView},
		{path: '/users/', component: SearchUserView}

	]
})

export default router
