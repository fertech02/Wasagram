import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SearchView from '../views/SearchView.vue'
import ProfileView from '../views/ProfileView.vue'
import PersonalProfile from '../views/PersonalProfile.vue'
import SetMyUsername from '../views/SetMyUserName.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/home', component: HomeView},
		{path: '/search', component: SearchView},
		{path: '/personalProfile', component: PersonalProfile},
		{path: '/users/:uid/profile', component: ProfileView},
		{path: '/set-name', component: SetMyUsername},
	]
})

export default router