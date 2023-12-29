import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import AuthWrapper from './components/AuthWrapper.vue'
import Post from './components/Post.vue'
import Modal from './components/Modal.vue'

import './assets/dashboard.css'
import './assets/main.css'
import './assets/feed.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("AuthWrapper", AuthWrapper);
app.component("Post",Post)
app.component("Modal",Modal)
app.use(router)
app.mount('#app')
