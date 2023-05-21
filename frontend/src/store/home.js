import { defineStore } from "pinia" 
import homeService from "../services/home" 

export const useHomeStore = defineStore({
    id: "home",
    state: () => ({ 
        post_data: null,
        update_message: null,
        delete_message: null
    }),  

    actions: {
        getPosts(token){ 
            homeService.getPosts(token).then(res => { 
                this.post_data = res.data
           })
        }, 
        /* updatePosts(data, token){ 
            homeService.updatePosts(data, token).then(res => { 
                this.update_message = res.data
           })
        },  */
        deletePosts(data, token){ 
            homeService.deletePosts(data, token).then(res => { 
                this.delete_message = res.data
           })
        }, 
    }, 

    getters: { 
        getPostsData(state) {
            return state.post_data
        },
       /*  getUpdateMessage(state) {
            return state.update_message
        }, */
        getDeleteMessage(state) {
            return state.delete_message
        }

    }
})