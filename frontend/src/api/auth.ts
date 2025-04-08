import axios from "axios";

export type SignUpInput = {
  email: string;
  username: string;
  password: string;
};

function signup(payload: SignUpInput) {
  return axios.post("http://localhost:8080/v1/user/signup", payload, {
    withCredentials: true,
  });
}

export type LoginInput = {
  username: string;
  password: string;
};

function login(payload: LoginInput) {
  return axios.post("http://localhost:8080/v1/user/login", payload, {
    withCredentials: true,
  });
}

export { signup, login };
