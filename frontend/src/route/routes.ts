import { createWebHistory, createRouter } from "vue-router";

import HomePage from "../views/HomePage.vue";
import MessagesPage from "../views/MessagesPage.vue";
import SignUpPage from "../views/SignUpPage.vue";

const routes = [
  { path: "/", component: HomePage },
  { path: "/messages", component: MessagesPage },
  { path: "/signup", component: SignUpPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
