import { defineStore } from "pinia" 
import userhService from "../services/user"

export const useAuthStore = defineStore({
    id: "auth",
    state: () => ({
        user: null,
        loading: false
    }),

    //login logout
    actions: {
        getUser(data){ 
           authService.user().then(res => {
            console.log(res);
           })
        },
        getProfile(data){ 
            authService.user().then(res => {
             console.log(res);
            })
         },
    }
})