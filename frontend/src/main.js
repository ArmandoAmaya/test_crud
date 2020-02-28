import Vue from "vue";
import VueRouter from "vue-router";
import { store } from "./store/store";
import { routes } from "./routes";
import App from "./App.vue";
import axios from "axios";

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import VueLadda from 'vue-ladda'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import "./assets/sass/main.scss";
  

Vue.component('vue-ladda', VueLadda)
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)


// Usar el enrutador
Vue.use(VueRouter);
Vue.prototype.$http = axios

// Token de sesión
const token = localStorage.getItem("token");
if (token) {
  Vue.prototype.$http.defaults.headers.common["Authorization"] = token;
}

// Establecer las rutas
export const router = new VueRouter({
  routes
});

// Capturar las entradas
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      next();
      return;
    }
    next("/login");
  } else if(to.matched.some(record => record.meta.onlyVisitor)) {
    if(!store.getters.isLoggedIn) {
      next();
      return;
    }
    next("/");
  } else {
    next();
  }
});

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App),
  created() {
    this.$http.interceptors.response.use(undefined, function(err) {
      console.log("INTERCEPTADO")

      return new Promise(() => {
        if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
          this.$store.dispatch("logout");
        }
        throw err;
      });
    });
  }
}).$mount("#app");
