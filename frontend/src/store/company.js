import { defineStore } from "pinia" 
import companyService from "../services/user" 

export const useUserStore = defineStore({
    id: "company",
    state: () => ({ 
        company_data: null
    }),

    //login logout
    actions: {
        getCompany(data){ 
            companyService.getCompany(data).then(res => { 
                this.company_data = res
           })
        }, 
    }, 
    getters: { 
        getUserData(state) {
            return state.company_data
        }
    }
})