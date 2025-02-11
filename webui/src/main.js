import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import PhotoCard from './components/PhotoCard.vue'
import CommentModal from './components/CommentModal.vue'
import CommentListModal from './components/CommentListModal.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("PhotoCard", PhotoCard);
app.component("CommentModal", CommentModal);
app.component("CommentListModal", CommentListModal);
app.use(router)
app.mount('#app')
