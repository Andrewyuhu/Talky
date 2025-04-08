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

export { signup };
