import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import installI18n from './lang/index'

const app = createApp(App)
installI18n(app)
app.use(store).use(router).mount('#app')
