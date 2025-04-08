import { createApp } from "vue";
import PrimeVue from "primevue/config";
import "./style.css";
import App from "./App.vue";
import Aura from "@primeuix/themes/aura";
import routes from "./route/routes.ts";
import Ripple from "primevue/ripple";

const app = createApp(App);
app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
});
app.use(routes);
app.directive("ripple", Ripple);
app.mount("#app");
