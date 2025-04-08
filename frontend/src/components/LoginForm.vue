<script lang="ts" setup>
import { ref } from "vue";
import { addError, isNotBlank } from "../utils/validator";
import { login } from "../api/auth";
import type { LoginInput } from "../api/auth";
import { isAxiosError } from "axios";
import { useRouter } from "vue-router";
const usernameInput = ref("");
const passwordInput = ref("");
const errors = ref<Map<string, string>>(new Map());
const router = useRouter();

async function submitForm(e: Event) {
  e.preventDefault();
  console.log(errors.value);
  if (!isValidForm()) {
    console.log("invalid");
    return;
  }

  try {
    const payload: LoginInput = {
      username: usernameInput.value,
      password: passwordInput.value,
    };
    await login(payload);
    router.push("/");
  } catch (error) {
    if (isAxiosError(error)) {
      switch (error.response?.status) {
        case 401:
          errors.value = new Map([["login", "invalid login credentials"]]);
          break;
        default:
          console.log(error);
      }
    }
  }
}

function isValidForm(): boolean {
  const newErrors: Map<string, string> = new Map();
  addError(
    isNotBlank(usernameInput.value),
    "username",
    "username cannot be blank",
    newErrors
  );
  addError(
    isNotBlank(passwordInput.value),
    "password",
    "password cannot be blank",
    newErrors
  );

  errors.value = newErrors;
  return errors.value.size === 0;
}
</script>

<template>
  <div
    class="w-9/10 sm:w-2/3 lg:w-2/5 2xl:w-1/5 border-2 self-start p-8 rounded-md"
  >
    <form class="flex flex-col gap-4 items-center">
      <h2 class="text-2xl font-bold">Sign in</h2>
      <div class="flex flex-col w-full">
        <label for="username">Username</label>
        <input
          type="text"
          name="username"
          v-model="usernameInput"
          class="border border-gray-400 rounded-lg px-4 py-2"
        />
        <span v-if="errors.has('username')" class="text-red-600">
          {{ errors.get("username") }}
        </span>
      </div>
      <div class="flex flex-col w-full">
        <label for="password">Password</label>
        <input
          type="password"
          name="password"
          v-model="passwordInput"
          class="border border-gray-400 rounded-lg px-4 py-2"
        />
        <span v-if="errors.has('password')" class="text-red-600">
          {{ errors.get("password") }}
        </span>
      </div>
      <span v-if="errors.has('login')" class="text-red-600">
        {{ errors.get("login") }}
      </span>
      <button
        class="relative border-2 px-4 py-2 rounded-md border-black z-10 bg-white w-full"
        @click="submitForm"
      >
        Login
      </button>
    </form>
  </div>
</template>
