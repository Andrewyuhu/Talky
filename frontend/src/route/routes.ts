import { createWebHistory, createRouter } from "vue-router";

import HomePage from "../views/HomePage.vue";
import MessagesPage from "../views/MessagesPage.vue";
import SignUpPage from "../views/SignUpPage.vue";
import LoginPage from "../views/LoginPage.vue";
import { useAuthStore } from "../store/auth";

const routes = [
  { path: "/", component: HomePage },
  { path: "/messages", component: MessagesPage, meta: { requireAuth: true } },
  { path: "/signup", component: SignUpPage },
  { path: "/login", component: LoginPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, _, next) => {
  const auth = useAuthStore();

  if (!to.meta?.requireAuth) {
    next();
    return;
  }

  await auth.checkAuth();

  if (!auth.isAuthenticated) {
    next("/login");
    return;
  }
  next();
});

export default router;
