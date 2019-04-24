import Vue from "vue";
import app from "./app.vue";
import Environment from "./environment";
import * as store from "./store";

Vue.config.productionTip = false;

Environment.setup()
  .then(() => {
    store.setup({ url: Environment.URL });
    new Vue({
      render: h => h(app)
    }).$mount("#app");
  })
  .catch(err => {
    // eslint-disable-next-line no-console
    console.error("catch error err:", err);
  });
