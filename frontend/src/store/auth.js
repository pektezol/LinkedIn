import { defineStore } from "pinia" 
import authService from "../services/auth"

export const useAuthStore = defineStore({
    id: "auth",
    state: () => ({
        user: null,
        loading: false
    }),

    //login logout
    actions: {
        handleLogin(data){
           console.log("ss",data) 
           authService.user().then(res => {
            console.log(res);
           })
        },
        handleRegister(data) {
            console.log(data);
            authService.userRegister(data).then(res => {
                console.log(res)
            })
        }
    }
})