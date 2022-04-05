import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// import store from './store/index'
import App from './App.vue'
import router from './router/index'

const app = createApp(App)

app.use(router)
    .use(ElementPlus)
    // .use(store)
    .mount('#app')