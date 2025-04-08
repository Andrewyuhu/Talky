<script lang="ts" setup>
import { useRouter } from "vue-router";
import { signup } from "../api/auth.ts";
import type { SignUpInput } from "../api/auth.ts";
import {
  isNotBlank,
  isValidEmail,
  addError,
  isMinLength,
  isMaxLength,
  isAlphanumeric,
} from "../utils/validator.ts";
import { ref } from "vue";
import { isAxiosError } from "axios";

const router = useRouter();
const emailInput = ref("");
const passwordInput = ref("");
const usernameInput = ref("");
const errors = ref(new Map<string, string>());

async function submitForm(e: Event) {
  e.preventDefault();
  if (!isValidForm()) {
    return;
  }
  try {
    const payload: SignUpInput = {
      email: emailInput.value,
      username: usernameInput.value,
      password: passwordInput.value,
    };
    await signup(payload);
    router.push("/");
  } catch (error) {
    if (isAxiosError(error)) {
      switch (error.response?.status) {
        case 422:
          if (error.response?.data.error == "EMAIL_EXISTS") {
            errors.value = new Map([["email", "email already exists"]]);
          }
          if (error.response?.data.error == "USERNAME_EXISTS") {
            errors.value = new Map([["username", "username already exists"]]);
          }
      }
      return;
    }
  }
}

function isValidForm(): boolean {
  const newErrors: Map<string, string> = new Map();
  addError(
    isNotBlank(emailInput.value),
    "email",
    "email cannot be blank",
    newErrors
  );
  addError(
    isValidEmail(emailInput.value),
    "email",
    "email is not in a valid format",
    newErrors
  );
  addError(
    isNotBlank(usernameInput.value),
    "username",
    "username cannot be blank",
    newErrors
  );
  addError(
    isAlphanumeric(usernameInput.value),
    "username",
    "username cannot contain non-alphanumeric characters",
    newErrors
  );
  addError(
    isMinLength(usernameInput.value, 3),
    "username",
    "username must be at least 3 characters",
    newErrors
  );
  addError(
    isMaxLength(usernameInput.value, 12),
    "username",
    "username must be less than 12 characters",
    newErrors
  );
  addError(
    isNotBlank(emailInput.value),
    "email",
    "email cannot be blank",
    newErrors
  );
  addError(
    isNotBlank(passwordInput.value),
    "password",
    "password cannot be blank",
    newErrors
  );
  addError(
    isMinLength(passwordInput.value, 6),
    "password",
    "password must be at least 6 characters",
    newErrors
  );
  console.log(newErrors);
  errors.value = newErrors;
  return errors.value.size === 0;
}
</script>

<template>
  <div
    class="w-9/10 sm:w-2/3 lg:w-2/5 2xl:w-1/5 border-2 self-start p-8 rounded-md"
  >
    <form class="flex flex-col gap-4 items-center">
      <h2 class="text-2xl font-bold">Sign up right now!</h2>
      <div class="flex flex-col w-full">
        <label for="email">Email</label>
        <input
          v-model="emailInput"
          type="email"
          name="email"
          class="border border-gray-400 rounded-lg px-4 py-2"
        />
        <span v-if="errors.has('email')" class="text-red-600">{{
          errors.get("email")
        }}</span>
      </div>

      <div class="flex flex-col w-full">
        <label for="username">Username</label>
        <input
          v-model="usernameInput"
          type="text"
          name="username"
          class="border border-gray-400 rounded-lg px-4 py-2"
        />
        <span v-if="errors.has('username')" class="text-red-600">{{
          errors.get("username")
        }}</span>
      </div>
      <div class="flex flex-col w-full">
        <label for="passowrd">Password</label>
        <input
          v-model="passwordInput"
          type="password"
          name="password"
          class="border border-gray-400 rounded-lg px-4 py-2"
        />
        <span v-if="errors.has('password')" class="text-red-600">{{
          errors.get("password")
        }}</span>
      </div>
      <button
        class="relative border-2 px-4 py-2 rounded-md border-black z-10 bg-white w-full"
        @click="submitForm"
      >
        Sign Up
      </button>
    </form>
  </div>
</template>
