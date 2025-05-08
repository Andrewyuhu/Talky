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
    router.push("/messages");
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
  <div class="w-full">
    <form
      class="flex flex-col gap-4 items-center w-[85%] mt-24 mx-auto px-4 py-8 rounded-md shadow-md border-1 border-black.0.5"
    >
      <h2 class="text-xl self-start text-primary-text font-semibold">
        Sign in
      </h2>
      <div class="flex flex-col w-full">
        <label for="username" class="text-sm text-zinc-500 mb-0.5"
          >Username</label
        >
        <input
          type="text"
          name="username"
          v-model="usernameInput"
          class="border border-primary-text rounded-lg px-4 py-2 text-primary-text"
        />
        <span v-if="errors.has('username')" class="text-red-500 text-sm mt-0.5">
          {{ errors.get("username") }}
        </span>
      </div>
      <div class="flex flex-col w-full">
        <label for="password" class="text-sm text-zinc-500 mb-0.5"
          >Password</label
        >
        <input
          type="password"
          name="password"
          v-model="passwordInput"
          class="border border-primary-text rounded-lg px-4 py-2 text-primary-text"
        />
        <span v-if="errors.has('password')" class="text-red-500 text-sm mt-0.5">
          {{ errors.get("password") }}
        </span>
      </div>
      <span v-if="errors.has('login')" class="text-red-500 text-sm">
        {{ errors.get("login") }}
      </span>
      <button
        class="px-3 py-1 rounded-lg text-primary-text bg-primary-green border-solid border border-primary-text w-full"
        @click="submitForm"
      >
        Login
      </button>
    </form>
  </div>
</template>
