import { defineStore } from "pinia" 
import userService from "../services/profile"
import router from "../router/index"

export const useProfileStore = defineStore({
    id: "auth",
    state: () => ({
        profile_data: null, 
    }),

    //login logout
    actions: { 
        getProfile(token){ 
            userService.getProfile(token).then(res => {
                console.log("profile123",res.data.user);
                if (res.data.status == "error") {
                    router.push('/login')
                }
                if (true){
                    this.profile_data = res.data.user
                }   
            })
         },
    }, 
    getters: {
        getProfileData(state) {
            return state.profile_data
        }, 
    }
})