import { defineStore } from "pinia" 
import homeService from "../services/home" 

export const useHomeStore = defineStore({
    id: "home",
    state: () => ({ 
        post_data: null
    }),  

    actions: {
        getPosts(data){ 
            homeService.getPosts(data).then(res => { 
                this.user_data = res
           })
        }, 
    }, 

    getters: { 
        getPostsData(state) {
            return state.user_data
        }
    }
})