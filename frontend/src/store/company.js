import { defineStore } from "pinia" 
import companyService from "../services/user" 

export const useUserStore = defineStore({
    id: "company",
    state: () => ({ 
        company_data: null,
        update_company_mesage: null,
        create_company_message: null
    }),

    //login logout
    actions: {
        getCompany(data){ 
            companyService.getCompany(data).then(res => { 
                this.company_data = res
           })
        }, 
        updateCompany(data){ 
            companyService.updateCompany(data).then(res => { 
                this.update_company_mesage = res
           })
        }, 
        createCompany(data){ 
            companyService.createCompany(data).then(res => { 
                this.create_company_message = res
           })
        },  
    }, 
    getters: { 
        getUserData(state) {
            return state.company_data
        }
    }
})