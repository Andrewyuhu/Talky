import { defineStore } from "pinia";
import { authenticateUser } from "../api/auth";
import { logout as logoutAPI } from "../api/auth";
import { type User } from "../types/users";

type AuthState = {
  user: User | null;
  isAuthenticated: boolean;
  error: boolean;
};

export const useAuthStore = defineStore("auth", {
  state: (): AuthState => ({
    user: null,
    isAuthenticated: false,
    error: false,
  }),
  actions: {
    async checkAuth() {
      console.log("checking auth");
      try {
        const response = await authenticateUser();
        this.user = response.data.user;
        this.isAuthenticated = true;
      } catch (error) {
        this.user = null;
        this.isAuthenticated = false;
        console.log(error);
        return;
      }
    },
    async logout() {
      this.error = false;
      this.user = null;
      this.isAuthenticated = false;
      try {
        await logoutAPI();
      } catch (err) {
        this.error = true;
      }
    },
  },
});
