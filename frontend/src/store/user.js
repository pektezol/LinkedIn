import { defineStore } from "pinia" 
import userService from "../services/user" 

export const useUserStore = defineStore({
    id: "auth",
    state: () => ({ 
        user_data: null
    }),

    //login logout
    actions: {
        getUser(data){ 
            userService.getUser(data).then(res => { 
                this.user_data = res
           })
        }, 
    }, 
    getters: { 
        getUserData(state) {
            return state.user_data
        }
    }
})