import { defineStore } from "pinia" 
import messageService from "../services/message" 

export const useMessageStore = defineStore({
    id: "message",
    state: () => ({ 
        message_data: null
    }),

    //login logout
    actions: {
        getUser(data){ 
            messageService.getUser(data).then(res => { 
                this.message_data = res
           })
        }, 
    }, 
    getters: { 
        getMessageData(state) {
            return state.message_data
        }
    }
})